package request

// GetOrderInfo
type GetOrderInfo struct {
	LogicId string `json:"logic_id" form:"logic_id" validate:"required"`
}
