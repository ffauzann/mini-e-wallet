package repository

import (
	"context"
	"database/sql"

	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *dbRepository) BeginTX(ctx context.Context) (tx *sql.Tx, err error) {
	tx, err = r.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}
