package service

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) isAccountEnabled(ctx context.Context, id string) (enabled bool) { // nolint
	enabled, err := r.repository.redis.IsAccountEnabled(ctx, id)
	if err != nil {
		util.Log().Error(err.Error())
		if _, err = r.getAccountDetail(ctx, id); err != nil {
			util.Log().Error(err.Error())
			return
		}
		return r.isAccountEnabled(ctx, id) // Recursion
	}
	return
}

func (r *service) getAccountDetail(ctx context.Context, id string) (account *model.AccountDetail, err error) {
	account, err = r.repository.db.GetAccountDetail(ctx, id)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}
	go r.repository.redis.SetAccountStatus(context.Background(), id, account.IsEnabled) // nolint

	return
}
