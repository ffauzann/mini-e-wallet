package http

import (
	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
)

// @Summary	Get transaction history
// @Tags	Transaction Management
// @Accept	json
// @Produce	json
// @Param	Authorization	header		string	true	"Token"		default(Token xxx)
// @Router	/wallet/transactions	[get]
func (s *srv) History(c echo.Context) (err error) {
	ctx := c.Request().Context()
	accountID := c.Get("accountID").(string)
	req := &model.TokenRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.History(ctx, &model.HistoryRequest{
		AccountID: accountID,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	// Prevent sending null `data.transactions` and send empty slice instead
	if len(res.Transactions) == 0 {
		res.Transactions = make([]*model.HistoryResultTransaction, 0)
	}

	return s.success(c, 200, res)
}

// @Summary	Make a deposit
// @Tags	Transaction Management
// @Accept	application/x-www-form-urlencoded
// @Produce	json
// @Param	Authorization	header		string	true	"Token"				default(Token xxx)
// @Param	reference_id	formData	string	true	"UUID"				default(49437636-fa79-40fb-b5cf-5f066235fdda)
// @Param	amount			formData	float64	true	"Deposit Amount"	default(10000.00)
// @Router	/wallet/deposits	[post]
func (s *srv) Deposit(c echo.Context) (err error) {
	ctx := c.Request().Context()
	accountID := c.Get("accountID").(string)
	req := &model.CreateTransactionRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.Deposit(ctx, &model.CreateTransactionRequest{
		AccountID:   accountID,
		Type:        constant.TransactionTypeDeposit,
		Amount:      req.Amount,
		ReferenceID: req.ReferenceID,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	return s.success(c, 201, res)
}

// @Summary	Request a withdrawal
// @Tags	Transaction Management
// @Accept	application/x-www-form-urlencoded
// @Produce	json
// @Param	Authorization	header		string	true	"Token"				default(Token xxx)
// @Param	reference_id	formData	string	true	"UUID"				default(49437636-fa79-40fb-b5cf-5f066235fdda)
// @Param	amount			formData	float64	true	"Withdrawal Amount"	default(10000.00)
// @Router	/wallet/withdrawals	[post]
func (s *srv) Withdrawal(c echo.Context) (err error) {
	ctx := c.Request().Context()
	accountID := c.Get("accountID").(string)
	req := &model.CreateTransactionRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.Withdrawal(ctx, &model.CreateTransactionRequest{
		AccountID:   accountID,
		Type:        constant.TransactionTypeWithdrawal,
		Amount:      req.Amount,
		ReferenceID: req.ReferenceID,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	return s.success(c, 201, res)
}
