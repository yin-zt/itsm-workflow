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

// OA系统回调推送数据
type OaCallB struct {
	LogicId        string `json:"logic_id"`         // 业务id
	Type           string `json:"type"`             // 审批类型
	Code           uint   `json:"code"`             // 60审批通过，20审批驳回
	AprCurrentName string `json:"apr_current_name"` // 审批人
	AprCurrentMid  uint   `json:"apr_current_mid"`  // 审批人min
}

type OrderInfo struct {
	ProcessInstance ProcessInstance `json:"processInstance" validate:"required"`
	UserTaskList    []UserTask      `json:"userTaskList"`
	UserTaskInfo    []UserTaskInfo  `json:"userTaskInfo"`
	StepList        []StepList      `json:"stepList"`
	FormData        string          `json:"formData"`
}

type ProcessInstance struct {
	InstanceId string `json:"instanceId" validate:"required"`
	Status     string `json:"status" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Category   string `json:"category" validate:"required"`
	Creator    string `json:"creator" validate:"required"`
}

type UserTaskInfo struct {
	Assignee   []string `json:"assignee"`
	Status     string   `json:"status"`
	UserTaskId string   `json:"userTaskId"`
}

type StepList struct {
	InstanceId string `json:"instanceId"`
	Status     string `json:"status"`
	//FormData   string `json:"formData"`
}

type UserTask struct {
	Type           string `json:"type"`
	FormDefinition string `json:"formDefinition"`
}
