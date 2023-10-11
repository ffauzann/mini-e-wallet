package repository

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *dbRepository) CreateAccount(ctx context.Context, id string) (exists bool, err error) {
	// qExists := `SELECT exists (SELECT id FROM accounts WHERE id = ? AND deleted_at IS NULL)`
	// if err = r.db.QueryRowContext(ctx, qExists, id).Scan(&exists); err != nil {
	// 	util.Log().Error(err.Error())
	// 	return
	// }
	// if exists {
	// 	return
	// }

	qInsert := `INSERT IGNORE INTO accounts(id) VALUES(?)`
	result, err := r.db.ExecContext(ctx, qInsert, id)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	affected, err := result.RowsAffected()
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return affected == 0, nil
}

func (r *dbRepository) GetAccountDetail(ctx context.Context, id string) (res *model.AccountDetail, err error) {
	query := `SELECT id, balance, is_enabled, updated_at FROM accounts WHERE id = ? AND deleted_at IS NULL`

	res = new(model.AccountDetail)
	if err = r.db.QueryRowContext(ctx, query, id).Scan(
		&res.ID,
		&res.Balance,
		&res.IsEnabled,
		&res.UpdatedAt,
	); err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}

func (r *dbRepository) UpdateAccountStatus(ctx context.Context, req *model.ActivationRequest) (err error) {
	query := `UPDATE accounts SET is_enabled = ?, updated_at = ? WHERE id = ? AND deleted_at IS NULL`
	_, err = r.db.ExecContext(ctx, query, req.IsEnabled, req.UpdatedAt, req.AccountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}
