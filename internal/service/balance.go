package service

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) Balance(ctx context.Context, req *model.BalanceRequest) (res *model.BalanceResult, err error) {
	// Validate whether account is enabled
	if !r.isAccountEnabled(ctx, req.AccountID) {
		err = constant.ErrAccountDisabled
		return
	}

	// Get account detail
	account, err := r.getAccountDetail(ctx, req.AccountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Validate whether it's enabled
	if !account.IsEnabled {
		err = constant.ErrAccountDisabled
		return
	}

	return &model.BalanceResult{
		Wallet: model.BalanceResultWallet{
			ID:        account.ID,
			OwnedBy:   account.ID,
			Status:    "enabled",
			EnabledAt: account.UpdatedAt,
			Balance:   account.Balance,
		},
	}, nil
}
