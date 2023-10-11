package service

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"go.uber.org/zap"
)

func (r *service) Init(ctx context.Context, req *model.InitRequest) (res *model.InitResult, err error) {
	// Validate whether account with given ID exists then create one if it's not
	exists, err := r.repository.db.CreateAccount(ctx, req.CustomerXID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}
	if exists {
		util.Log().Info("account already exists", zap.String("id", req.CustomerXID))
		err = constant.ErrAccountAlreadyExists
		return
	}

	// Generate token using AES-GCM and return as hex
	token, err := util.Encrypt(req.CustomerXID, r.config.Encryption.Key)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	// Store is_enabled to redis since it will be checked frequently
	go r.repository.redis.SetAccountStatus(context.Background(), req.CustomerXID, false) // nolint

	return &model.InitResult{
		Token: token,
	}, nil
}
