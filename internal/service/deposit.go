package service

import (
	"context"
	"time"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) Deposit(ctx context.Context, req *model.CreateTransactionRequest) (res *model.DepositResult, err error) {
	// Validate whether account is enabled
	if !r.isAccountEnabled(ctx, req.AccountID) {
		err = constant.ErrAccountDisabled
		return
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

	// Force set status to success since there's no certain requirements for deposit to be succeed
	req.Status = constant.TransactionStatusSuccess

	// Create transaction and update account balance
	if err = r.repository.db.CreateTransaction(ctx, req); err != nil {
		util.Log().Error(err.Error())
		return
	}

	return &model.DepositResult{
		Deposit: model.DepositResultDeposit{
			ID:          req.ID,
			DepositedBy: req.AccountID,
			Status:      req.Status,
			DepositedAt: time.Now(), // Should've taken from database or set this value instead of default value of column
			Amount:      req.Amount,
			ReferenceID: req.ReferenceID,
		},
	}, nil
}
