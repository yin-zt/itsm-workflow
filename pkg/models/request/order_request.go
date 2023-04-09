package request

// GetOrderInfo
type GetOrderInfo struct {
	LogicId string `json:"logic_id" form:"logic_id" validate:"required"`
}

// OrderAddReq 添加资源结构体
type OrderAddReq struct {
	OrderId   string `json:"orderId" validate:"required"`
	OrderName string `json:"orderName" validate:"required,min=1,max=20"`
	ParentId  uint   `json:"parentId" validate:"omitempty,min=0"`
	Remark    string `json:"remark" validate:"min=0,max=100"` // 分组的中文描述
}
