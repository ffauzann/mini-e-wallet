package repository

import (
	"context"
	"database/sql"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/google/uuid"
)

func (r *dbRepository) GetAllTransactionHistory(ctx context.Context, accountID string) (res []*model.HistoryResultTransaction, err error) {
	query := `SELECT id, reference_id, type, amount, status, created_at FROM transactions WHERE account_id = ? AND deleted_at IS NULL ORDER BY created_at DESC LIMIT 100`
	rows, err := r.db.QueryContext(ctx, query, accountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	for rows.Next() {
		var trx model.HistoryResultTransaction
		if err = rows.Scan(
			&trx.ID,
			&trx.ReferenceID,
			&trx.Type,
			&trx.Amount,
			&trx.Status,
			&trx.TransactedAt,
		); err != nil {
			util.Log().Error(err.Error())
			return
		}

		res = append(res, &trx)
	}

	return
}

func (r *dbRepository) IsReferenceIDExists(ctx context.Context, accountID, referenceID string) (exists bool, err error) {
	query := `SELECT exists (SELECT id FROM transactions WHERE account_id = ? AND reference_id = ? AND deleted_at IS NULL)`
	if err = r.db.QueryRowContext(ctx, query, accountID, referenceID).Scan(&exists); err != nil {
		util.Log().Error(err.Error())
		return
	}
	return
}

func (r *dbRepository) CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (err error) {
	// Begin transaction
	tx, err := r.BeginTX(ctx)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Generate UUID
	req.ID = uuid.NewString()

	// Insert transaction data
	query := `INSERT INTO transactions(id, account_id, reference_id, type, amount, status) VALUES(?, ?, ?, ?, ?, ?)`
	if _, err = tx.ExecContext(
		ctx,
		query,
		req.ID,
		req.AccountID,
		req.ReferenceID,
		req.Type,
		req.Amount,
		req.Status,
	); err != nil {
		util.Log().Error(err.Error())
		if err = tx.Rollback(); err != nil {
			util.Log().Error(err.Error())
			return
		}
		return
	}

	// Update balance if trx status is success
	if req.Status == constant.TransactionStatusSuccess {
		return r.updateBalance(ctx, tx, req)
	}

	if err = tx.Commit(); err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) updateBalance(ctx context.Context, tx *sql.Tx, req *model.CreateTransactionRequest) (err error) {
	amount := req.Amount
	if req.Type == constant.TransactionTypeWithdrawal {
		amount *= -1
	}

	query := `UPDATE accounts SET balance = balance + ? WHERE id = ? AND deleted_at IS NULL`
	if _, err = tx.ExecContext(ctx, query, amount, req.AccountID); err != nil {
		util.Log().Error(err.Error())
		if err = tx.Rollback(); err != nil {
			util.Log().Error(err.Error())
			return
		}
		return
	}

	if err = tx.Commit(); err != nil {
		util.Log().Error(err.Error())
		return
	}
	return
}
