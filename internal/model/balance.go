package model

import "time"

type (
	BalanceRequest struct {
		AccountID string
	}

	BalanceResult struct {
		Wallet BalanceResultWallet `json:"wallet"`
	}
	BalanceResultWallet struct {
		ID        string    `json:"id"`
		OwnedBy   string    `json:"owned_by"`
		Status    string    `json:"status"`
		EnabledAt time.Time `json:"enabled_at"`
		Balance   float64   `json:"balance"`
	}
)
