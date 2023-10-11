package main

import (
	"github.com/ffauzann/mini-e-wallet/internal/app"
)

// ========================================
// SWAGGER ANNOTATION
// ========================================
//	@title		Mini E-Wallet API
//	@version	1.0
//	@host		localhost:2201
//	@BasePath	/api/v1
// ========================================

var cfg app.Config

func init() {
	cfg.Setup()
}

func main() {
	cfg.StartServer()
}
