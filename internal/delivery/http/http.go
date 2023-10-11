package http

import (
	"github.com/ffauzann/mini-e-wallet/internal/middleware"
	"github.com/ffauzann/mini-e-wallet/internal/service"
	"github.com/labstack/echo/v4"
)

type srv struct {
	service service.Service
}

func New(e *echo.Echo, serviceIns service.Service, key string) {
	s := srv{
		service: serviceIns,
	}

	g := e.Group("/api/v1")
	healthRoutes(g)
	accountRoutes(g, &s, key)
	transactionRoutes(g, &s, key)
}

func healthRoutes(g *echo.Group) {
	g.GET("/health", health)
	g.GET("/readiness", readiness)
}

func accountRoutes(g *echo.Group, s *srv, key string) {
	g.POST("/init", s.Init)

	g.Use(middleware.Token(key))
	g.GET("/wallet", s.GetBalance)
	g.POST("/wallet", s.EnableWallet)
	g.PATCH("/wallet", s.DisableWallet)
}

func transactionRoutes(g *echo.Group, s *srv, key string) {
	g.Use(middleware.Token(key))
	g.GET("/wallet/transactions", s.History)
	g.POST("/wallet/withdrawals", s.Withdrawal)
	g.POST("/wallet/deposits", s.Deposit)
}
