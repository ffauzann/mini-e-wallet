package service

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/repository"
)

type Service interface {
	AccountService
	TransactionService
}

type AccountService interface {
	Init(ctx context.Context, req *model.InitRequest) (res *model.InitResult, err error)
	Balance(ctx context.Context, req *model.BalanceRequest) (res *model.BalanceResult, err error)
	Activation(ctx context.Context, req *model.ActivationRequest) (res *model.ActivationResult, err error)
}

type TransactionService interface {
	History(ctx context.Context, req *model.HistoryRequest) (res *model.HistoryResult, err error)
	Deposit(ctx context.Context, req *model.CreateTransactionRequest) (res *model.DepositResult, err error)
	Withdrawal(ctx context.Context, req *model.CreateTransactionRequest) (res *model.WithdrawalResult, err error)
}

type service struct {
	config     *model.AppConfig
	repository repositoryWrapper
}

type repositoryWrapper struct {
	db    repository.DBRepository
	redis repository.RedisRepository
}

func New(db repository.DBRepository, redis repository.RedisRepository, config *model.AppConfig) Service {
	return &service{
		config: config,
		repository: repositoryWrapper{
			db:    db,
			redis: redis,
		},
	}
}
