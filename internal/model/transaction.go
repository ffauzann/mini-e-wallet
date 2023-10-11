package model

import (
	"time"
)

type (
	HistoryRequest struct {
		AccountID string
	}
	CreateTransactionRequest struct {
		ID          string
		AccountID   string
		Type        string
		ReferenceID string  `form:"reference_id" validate:"required,uuid" example:"49437636-fa79-40fb-b5cf-5f066235fddb"`
		Amount      float64 `form:"amount" validate:"required,gte=10000,lte=1000000000" example:"10000.00"`
		Status      string
	}

	WithdrawalResult struct {
		Withdrawal WithdrawalResultWithdrawal `json:"withdrawal"`
	}
	WithdrawalResultWithdrawal struct {
		ID          string    `json:"id"`
		WithdrawnBy string    `json:"withdrawn_by"`
		Status      string    `json:"status"`
		WithdrawnAt time.Time `json:"withdrawn_at"`
		Amount      float64   `json:"amount"`
		ReferenceID string    `json:"reference_id"`
	}

	DepositResult struct {
		Deposit DepositResultDeposit `json:"deposit"`
	}
	DepositResultDeposit struct {
		ID          string    `json:"id"`
		DepositedBy string    `json:"deposited_by"`
		Status      string    `json:"status"`
		DepositedAt time.Time `json:"deposited_at"`
		Amount      float64   `json:"amount"`
		ReferenceID string    `json:"reference_id"`
	}

	HistoryResult struct {
		Transactions []*HistoryResultTransaction `json:"transactions"`
	}
	HistoryResultTransaction struct {
		ID           string    `json:"id"`
		ReferenceID  string    `json:"reference_id"`
		Status       string    `json:"status"`
		TransactedAt time.Time `json:"transacted_at"`
		Type         string    `json:"type"`
		Amount       float64   `json:"amount"`
	}
)
