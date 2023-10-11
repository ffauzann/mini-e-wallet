package model

import "time"

type TokenRequest struct {
	Token string `header:"token"`
}

type AccountDetail struct {
	ID        string
	Balance   float64
	IsEnabled bool
	UpdatedAt time.Time
}
