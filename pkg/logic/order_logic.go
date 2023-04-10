package logic

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/models/order"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
	"github.com/yin-zt/itsm-workflow/pkg/utils/isql"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
)

type OrderLogic struct{}

// GetOrderInfo
func (o OrderLogic) GetOrderInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GetOrderInfo)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	return r, nil
}

func (o OrderLogic) AnalyOrderInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.OrderInfo)
	if !ok {
		return nil, ReqAssertErr
	}

	stepLists := r.StepList
	instanceId := ""
	for _, item := range stepLists {
		if item.Status == "running" {
			instanceId = item.InstanceId
		}
	}
	orderId := fmt.Sprintf("%s%s", r.ProcessInstance.InstanceId, instanceId)
	if isql.Order.Exist(tools.H{"apply_logic_id": orderId}) {
		return nil, tools.NewValidatorError(fmt.Errorf("工单已存在,请勿重复添加"))
	}

	order := order.Order{
		ParentId:  r.ParentId,
		GroupName: r.GroupName,
		Remark:    r.Remark,
		Creator:   ctxUser.Username,
		Source:    "platform", //默认是平台添加
	}

	// 然后在数据库中创建组
	err := isql.Order.Add(&order)
	if err != nil {
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建分组失败"))
	}

	return nil, nil

	fmt.Println(r.FormData)
	return nil, nil
}

// AddOrderRecord 添加工单数据
//func (o OrderLogic) AddOrderRecord(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
//	r, ok := req.(*request.OrderAddReq)
//	if !ok {
//		return nil, ReqAssertErr
//	}
//	_ = c
//
//	if isql.Order.Exist(tools.H{"order_name": r.OrderName}) {
//		return nil, tools.NewValidatorError(fmt.Errorf("该工单对应工单ID已存在"))
//	}
//
//	// 获取当前用户
//	ctxUser, err := isql.User.GetCurrentLoginUser(c)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("获取当前登陆用户信息失败"))
//	}
//
//	group := model.Group{
//		ParentId:  r.ParentId,
//		GroupName: r.GroupName,
//		Remark:    r.Remark,
//		Creator:   ctxUser.Username,
//		Source:    "platform", //默认是平台添加
//	}
//
//	// 然后在数据库中创建组
//	err = isql.Group.Add(&group)
//	if err != nil {
//		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建分组失败"))
//	}
//
//	return nil, nil
//}
