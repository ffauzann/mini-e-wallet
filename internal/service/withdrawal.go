package service

import (
	"context"
	"time"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) Withdrawal(ctx context.Context, req *model.CreateTransactionRequest) (res *model.WithdrawalResult, err error) {
	// Validate whether account is enabled
	if !r.isAccountEnabled(ctx, req.AccountID) {
		err = constant.ErrAccountDisabled
		return
	}

	// Get account detail to compare between the available balance and requested withdrawal amount
	account, err := r.getAccountDetail(ctx, req.AccountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Set status, record anyway even if the balance is insufficient but return error later (if needed)
	req.Status = constant.TransactionStatusSuccess
	if account.Balance < req.Amount {
		req.Status = constant.TransactionStatusFailed
	}

	// Validate reference ID
	exists, err := r.repository.db.IsReferenceIDExists(ctx, req.AccountID, req.ReferenceID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}
	if exists {
		err = constant.ErrDuplicateReferenceID
		return
	}

	// Create transaction and update account balance
	if err = r.repository.db.CreateTransaction(ctx, req); err != nil {
		util.Log().Error(err.Error())
		return
	}

	// // Return error after transaction recorded
	// if account.Balance < req.Amount {
	// 	err = constant.ErrInsufficientBalance
	// 	return
	// }

	return &model.WithdrawalResult{
		Withdrawal: model.WithdrawalResultWithdrawal{
			ID:          req.ID,
			WithdrawnBy: req.AccountID,
			Status:      req.Status,
			WithdrawnAt: time.Now(), // Should've taken from database or set this value instead of default value of column
			Amount:      req.Amount,
			ReferenceID: req.ReferenceID,
		},
	}, nil
}
