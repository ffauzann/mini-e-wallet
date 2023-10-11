package service

import (
	"context"
	"time"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) Activation(ctx context.Context, req *model.ActivationRequest) (res *model.ActivationResult, err error) {
	// Get account detail
	account, err := r.getAccountDetail(ctx, req.AccountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Validate whether the requested status is already applied
	if account.IsEnabled && req.IsEnabled {
		err = constant.ErrAccountAlreadyEnabled
		return
	}
	if !account.IsEnabled && !req.IsEnabled {
		err = constant.ErrAccountAlreadyDisabled
		return
	}

	// Update account status
	req.UpdatedAt = time.Now()
	if err = r.repository.db.UpdateAccountStatus(ctx, req); err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Update account status in redis as well
	go r.repository.redis.SetAccountStatus(context.Background(), req.AccountID, req.IsEnabled) // nolint

	return &model.ActivationResult{
		AccountID: account.ID,
		IsEnabled: account.IsEnabled,
		Balance:   account.Balance,
		UpdatedAt: account.UpdatedAt,
	}, nil
}
