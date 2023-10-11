package model

type Response struct {
	Status string      `json:"status" example:"Success"`
	Data   interface{} `json:"data"`
}
