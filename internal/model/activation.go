package model

import "time"

type (
	ActivationRequest struct {
		AccountID string
		IsEnabled bool
		UpdatedAt time.Time
	}
	ActivationResult struct {
		AccountID string
		IsEnabled bool
		Balance   float64
		UpdatedAt time.Time
	}

	EnableWalletResponse struct {
		Wallet EnableWalletResponseWallet `json:"wallet"`
	}
	EnableWalletResponseWallet struct {
		ID        string    `json:"id"`
		OwnedBy   string    `json:"owned_by"`
		Status    string    `json:"status"`
		EnabledAt time.Time `json:"enabled_at"`
		Balance   float64   `json:"balance"`
	}

	DisableWalletResponse struct {
		Wallet DisableWalletResponseWallet `json:"wallet"`
	}
	DisableWalletResponseWallet struct {
		ID         string    `json:"id"`
		OwnedBy    string    `json:"owned_by"`
		Status     string    `json:"status"`
		DisabledAt time.Time `json:"disabled_at"`
		Balance    float64   `json:"balance"`
	}
)

func (r *ActivationResult) ToResponseEnable() *EnableWalletResponse {
	return &EnableWalletResponse{
		Wallet: EnableWalletResponseWallet{
			ID:        r.AccountID,
			OwnedBy:   r.AccountID,
			Status:    "enabled",
			EnabledAt: r.UpdatedAt,
			Balance:   r.Balance,
		},
	}
}

func (r *ActivationResult) ToResponseDisable() *DisableWalletResponse {
	return &DisableWalletResponse{
		Wallet: DisableWalletResponseWallet{
			ID:         r.AccountID,
			OwnedBy:    r.AccountID,
			Status:     "disabled",
			DisabledAt: r.UpdatedAt,
			Balance:    r.Balance,
		},
	}
}
