package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ApplyLogicId   string `gorm:"type:varchar(128);comment:'工单申请ID'" json:"apply_logic_id"`
	TaskId         string `gorm:"type:varchar(128);comment:'任务ID'" json:"task_id"`
	InstanceId     string `gorm:"type:varchar(128);comment:'节点ID'" json:"instance_id"`
	ApplyUser      string `gorm:"type:varchar(128);comment:'工单申请用户'" json:"apply_user"`
	ApplyStatus    string `gorm:"type:varchar(128);comment:'工单申请状态';default:'0'" json:"apply_status"` // 0: 未提交审批   1: 已提交审批
	ApplyType      string `gorm:"type:varchar(128);comment:'工单申请类型'" json:"apply_type"`
	ExecuteStatus  string `gorm:"type:varchar(128);comment:'工单执行状态'" json:"execute_status"` // 0: 待审批, 60:审批通过， 20:审批拒绝
	DisplayLabel   string `gorm:"type:varchar(128);comment:'工单详情字段'" json:"display_label"`
	DisplayContent string `gorm:"type:varchar(256);comment:'工单详情'" json:"display_content"`
}
