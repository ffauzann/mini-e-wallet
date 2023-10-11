package http

import (
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
)

// @Summary	Enable wallet
// @Tags	Account Management
// @Accept	json
// @Produce	json
// @Param	Authorization	header		string	true	"Token"		default(Token xxx)
// @Router	/wallet	[post]
func (s *srv) EnableWallet(c echo.Context) (err error) {
	ctx := c.Request().Context()
	accountID := c.Get("accountID").(string)
	req := &model.TokenRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.Activation(ctx, &model.ActivationRequest{
		AccountID: accountID,
		IsEnabled: true,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	return s.success(c, 200, res.ToResponseEnable())
}

// @Summary	Disable wallet
// @Tags	Account Management
// @Accept	json
// @Produce	json
// @Param	Authorization	header		string	true	"Token"		default(Token xxx)
// @Router	/wallet	[patch]
func (s *srv) DisableWallet(c echo.Context) (err error) {
	ctx := c.Request().Context()
	accountID := c.Get("accountID").(string)
	req := &model.TokenRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.Activation(ctx, &model.ActivationRequest{
		AccountID: accountID,
		IsEnabled: false,
	})
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	return s.success(c, 200, res.ToResponseDisable())
}
