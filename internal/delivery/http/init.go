package http

import (
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
)

// @Summary	Create new account
// @Tags	Account Management
// @Accept	application/x-www-form-urlencoded
// @Produce	json
// @Param	customer_xid	formData	string	true	"uuid" default(49437636-fa79-40fb-b5cf-5f066235fdda)
// @Router	/init	[post]
func (s *srv) Init(c echo.Context) (err error) {
	ctx := c.Request().Context()
	req := &model.InitRequest{}

	if code, err := bindAndValidate(c, req); err != nil {
		util.Log().Error(err.Error())
		return s.error(c, code, err)
	}

	res, err := s.service.Init(ctx, req)
	if err != nil {
		util.Log().Error(err.Error())
		return s.error(c, 0, err)
	}

	return s.success(c, 200, res)
}
