package middleware

import (
	"strings"

	"github.com/ffauzann/mini-e-wallet/internal/constant"
	"github.com/ffauzann/mini-e-wallet/internal/model"
	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
)

func Token(key string) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			mErr := model.Response{
				Status: constant.ResponseStatusFail,
				Data:   map[string]string{"error": constant.ErrUnauthorized.Error()},
			}

			auth := c.Request().Header.Get("Authorization")
			s := strings.Split(auth, " ")
			if len(s) != 2 || s[0] != "Token" {
				return c.JSON(401, mErr)
			}

			accountID, err := util.Decrypt(s[1], key)
			if err != nil {
				return c.JSON(401, mErr)
			}

			c.Set("accountID", accountID)

			return next(c)
		}
	}
}
