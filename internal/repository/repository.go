package repository

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/model"

	"github.com/go-redis/redis/v9"
	"github.com/jmoiron/sqlx"
)

func NewDB(db *sqlx.DB, config *model.AppConfig) DBRepository {
	return &dbRepository{
		db: db,
		common: common{
			config: config,
		},
	}
}

func NewRedis(client *redis.Client, config *model.AppConfig) RedisRepository {
	return &redisRepository{
		redis: client,
		common: common{
			config: config,
		},
	}
}

type DBRepository interface {
	// Account-related
	CreateAccount(ctx context.Context, id string) (exists bool, err error)
	GetAccountDetail(ctx context.Context, id string) (res *model.AccountDetail, err error)
	UpdateAccountStatus(ctx context.Context, req *model.ActivationRequest) (err error)

	// Transaction-related
	GetAllTransactionHistory(ctx context.Context, accountID string) (res []*model.HistoryResultTransaction, err error)
	IsReferenceIDExists(ctx context.Context, accountID, referenceID string) (exists bool, err error)
	CreateTransaction(ctx context.Context, req *model.CreateTransactionRequest) (err error)
}

type RedisRepository interface {
	IsAccountEnabled(ctx context.Context, id string) (enabled bool, err error)
	SetAccountStatus(ctx context.Context, id string, enabled bool) error
}

type common struct {
	config *model.AppConfig
}

type dbRepository struct {
	db *sqlx.DB
	common
}

type redisRepository struct {
	redis *redis.Client
	common
}
