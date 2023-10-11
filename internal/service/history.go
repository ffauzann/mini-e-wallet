package service

import (
	"context"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
)

func (r *service) History(ctx context.Context, req *model.HistoryRequest) (res *model.HistoryResult, err error) {
	// Validate whether account is enabled
	if !r.isAccountEnabled(ctx, req.AccountID) {
		err = constant.ErrAccountDisabled
		return
	}

	// Get all transaction history up to 100 records to reduce memory consumption and the risk being exploit
	res = new(model.HistoryResult)
	res.Transactions, err = r.repository.db.GetAllTransactionHistory(ctx, req.AccountID)
	if err != nil {
		util.Log().Error(err.Error())
		return
	}

	return
}
