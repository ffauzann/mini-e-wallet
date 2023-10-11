package model

type (
	InitRequest struct {
		CustomerXID string `json:"customer_xid" form:"customer_xid" validate:"required,uuid" example:"a25dc5b4-35bc-4cc5-8864-8eb41d83f2bf"`
	}

	InitResult struct {
		Token string `json:"token"`
	}
)

// func (r *InitResult) ToResponse() *Response {
// 	return &Response{
// 		Status: constant.ResponseStatusSuccess,
// 		Data:   r,
// 	}
// }
