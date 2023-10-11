package http

import (
	"net/http"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
)

// error handle status code mapping and filter error message into known or generic message
func (s *srv) error(c echo.Context, code int, err error) error {
	util.Log().Error(err.Error())

	if code == 0 { // Find proper status code if it's not given
		var ok bool
		code, ok = constant.MapErrorToStatusCode[err] // Check whether it's registered in map
		if !ok {
			// Default status code & error
			err = constant.ErrInternal
			code = http.StatusInternalServerError
		}
	}

	return c.JSON(code, model.Response{
		Status: constant.ResponseStatusFail,
		Data:   map[string]interface{}{"error": err.Error()},
	})
}

// error handle status code mapping and filter error message into known or generic message
func (s *srv) success(c echo.Context, code int, data interface{}) error {
	return c.JSON(code, model.Response{
		Status: constant.ResponseStatusSuccess,
		Data:   data,
	})
}
