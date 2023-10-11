package repository

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *redisRepository) IsAccountEnabled(ctx context.Context, id string) (enabled bool, err error) {
	cmd := r.redis.Get(ctx, fmt.Sprintf(constant.RedisAccountStatus, id))
	if err = cmd.Err(); err != nil {
		util.Log().Error(err.Error())
		return
	}

	enabled, err = strconv.ParseBool(cmd.Val())
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}

func (r *redisRepository) SetAccountStatus(ctx context.Context, id string, enabled bool) error {
	return r.redis.Set(ctx, fmt.Sprintf(constant.RedisAccountStatus, id), enabled, constant.RedisDefaultExp).Err()
}
