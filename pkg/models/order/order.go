package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	ApplyLogicId   string `gorm:"type:varchar(128);comment:'申请ID'" json:"apply_logic_id"`
	TaskId         string `gorm:"type:varchar(128);comment:'任务ID'" json:"task_id"`
	InstanceId     string `gorm:"type:varchar(128);comment:'节点ID'" json:"instance_id"`
	ApplyUser      string `gorm:"type:varchar(128);comment:'申请用户'" json:"apply_user"`
	ApplyType      string `gorm:"type:varchar(128);comment:'申请类型'" json:"apply_type"`
	OrderName      string `gorm:"type:varchar(128);comment:'工单标题'" json:"order_name"`
	ApplyStatus    uint   `gorm:"type:tinyint(1);default:0;comment:'提交状态(未提交/已提交, 默认未提交)'" json:"apply_status"`             // 0: 未提交审批   1: 已提交审批
	ApprovalUser   string `gorm:"type:varchar(256);comment:'审批人员'" json:"approval_user"`                                    // 工单审批人员
	ExecuteStatus  uint   `gorm:"type:tinyint(1);default:0;comment:'审批状态(待审批/审批通过/审批拒绝, 默认待审批)'" json:"execute_status"`     // 0: 待审批, 60:审批通过， 20:审批拒绝
	ItsmOperation  uint   `gorm:"type:tinyint(1);default:0;comment:'回调itsm状态(等待中/回调成功/回调异常, 默认等待中)'" json:"itsm_operation"` // 0:等待中, 1: 回调成功, 2: 回调失败
	DisplayContent string `gorm:"type:varchar(256);comment:'详情内容'" json:"display_content"`
	ItsmResult     string `gorm:"type:varchar(256);comment:'调用itsm接口返回结果'" json:"itsm_result"`
}
