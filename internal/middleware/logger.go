package middleware

import (
	"encoding/json"
	"regexp"

	"github.com/ffauzann/mini-e-wallet/internal/util"
	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func Logger() echo.MiddlewareFunc {
	return em.RequestLoggerWithConfig(em.RequestLoggerConfig{
		LogMethod:  true,
		LogStatus:  true,
		LogLatency: true,
		LogURIPath: true,
		LogError:   true,
		Skipper: func(c echo.Context) bool {
			return regexp.MustCompile("swagger").MatchString(c.Path())
		},
		LogValuesFunc: func(c echo.Context, v em.RequestLoggerValues) error {
			util.Log().Info(
				"request",
				zap.String("method", v.Method),
				zap.String("path", v.URIPath),
				zap.Duration("latency", v.Latency),
				zap.Int("code", v.Status),
				zap.Any("error", v.Error),
			)
			return nil
		},
	})
}

func BodyDump() echo.MiddlewareFunc {
	return em.BodyDumpWithConfig(em.BodyDumpConfig{
		Skipper: func(c echo.Context) bool {
			switch path := c.Path(); path {
			case "/api/v1/wallet/transactions":
				return true
			default:
				regex := regexp.MustCompile("swagger")
				if regex.MatchString(path) {
					return true
				}
				return false
			}
		},
		Handler: func(ctx echo.Context, reqBody, resBody []byte) {
			res := make(map[string]interface{})
			json.Unmarshal(resBody, &res) // nolint

			util.Log().Info("body", zap.String("req", string(reqBody)), zap.Any("res", res))
		},
	})
}
