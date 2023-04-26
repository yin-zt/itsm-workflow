package order

import (
	"time"
)

type T_Order struct {
	//gorm.Model
	Fid             uint      `gorm:"primary_key;AUTO_INCREMENT;column:F_id" json:"F_id"`
	FCreatedAt      time.Time `gorm:"column:F_created_at" json:"F_created_at"`
	FUpdatedAt      time.Time `gorm:"column:F_updated_at" json:"F_updated_at"`
	FDeletedAt      time.Time `gorm:"column:F_deleted_at" json:"F_deleted_at"`
	FApplyLogicId   string    `gorm:"type:varchar(128);comment:'申请ID';column:F_apply_logic_id" json:"F_apply_logic_id"`
	FTaskId         string    `gorm:"type:varchar(128);comment:'任务ID';column:F_task_id" json:"F_task_id"`
	FInstanceId     string    `gorm:"type:varchar(128);comment:'节点ID';column:F_instance_id" json:"F_instance_id"`
	FApplyUser      string    `gorm:"type:varchar(128);comment:'申请用户';column:F_apply_user" json:"F_apply_user"`
	FApplyType      string    `gorm:"type:varchar(128);comment:'申请类型';column:F_apply_type" json:"F_apply_type"`
	FOrderName      string    `gorm:"type:varchar(128);comment:'工单标题';column:F_order_name" json:"F_order_name"`
	FApplyStatus    uint      `gorm:"type:tinyint(1);default:0;comment:'提交状态(未提交/已提交, 默认未提交)';column:F_apply_status" json:"F_apply_status"`               // 0: 未提交审批   1: 已提交审批
	FApprovalUser   string    `gorm:"type:varchar(256);comment:'审批人员';column:F_approval_user" json:"F_approval_user"`                                     // 工单审批人员
	FExecuteStatus  uint      `gorm:"type:tinyint(1);default:0;comment:'审批状态(待审批/审批通过/审批拒绝, 默认待审批)';column:F_execute_status" json:"F_execute_status"`     // 0: 待审批, 60:审批通过， 20:审批拒绝
	FItsmOperation  uint      `gorm:"type:tinyint(1);default:0;comment:'回调itsm状态(等待中/回调成功/回调异常, 默认等待中)';column:F_itsm_operation" json:"F_itsm_operation"` // 0:等待中, 1: 回调成功, 2: 回调失败
	FDisplayContent string    `gorm:"type:varchar(256);comment:'详情内容';column:F_display_content" json:"F_display_content"`
	FItsmResult     string    `gorm:"type:varchar(256);comment:'调用itsm接口返回结果';column:F_itsm_result" json:"F_itsm_result"`
}
