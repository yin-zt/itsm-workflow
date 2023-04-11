package logic

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
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

	//order := order.Order{
	//	ParentId:  r.ParentId,
	//	GroupName: r.GroupName,
	//	Remark:    r.Remark,
	//	Creator:   ctxUser.Username,
	//	Source:    "platform", //默认是平台添加
	//}

	// 然后在数据库中创建组
	//err := isql.Order.Add(&order)
	//if err != nil {
	//	return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建分组失败"))
	//}

	fmt.Println(r.FormData)
	fmt.Println("pppppppppppppp")
	fmt.Println(r.UserTaskList)
	fmt.Println("ooooooooooooooooooooo")
	fmt.Println(r.UserTaskList[0].FormDefinition)
	fmt.Println("mmmmmmmmmmmmmmmmm")

	taskForms := r.UserTaskList
	res := o.FindOutLabelKeyVal(taskForms)
	fmt.Println(res)

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

// FindOutLabelKeyVal 作用是解析工单任务中字段名称与字段id的关联关系，并以字典的形式返回。
// 字典最外层的key是容器的id，里面的字典key为字段id，值为 字段的中文名
func (o OrderLogic) FindOutLabelKeyVal(tasks []request.UserTask) map[string]map[string]string {
	var retMap = map[string]map[string]string{}
	for _, task := range tasks {
		var m2 []map[string]interface{}
		json.Unmarshal([]byte(task.FormDefinition), &m2)
		if len(m2) == 0 {
			continue
		}
		for _, innerMap := range m2 {
			var mapKey = ""
			var bigMap = map[string]string{}
			fmt.Println(123)
			if labelName, ok := innerMap["key"]; ok {
				mapKey = labelName.(string)
				retMap[mapKey] = map[string]string{}
			} else {
				continue
			}
			fmt.Println(456)
			if propertys, ok := innerMap["propertys"]; ok {
				if propertysLists, ok := propertys.([]interface{}); ok {
					for _, oneProperty := range propertysLists {
						oneData := oneProperty.(map[string]interface{})
						labelName := oneData["label"]
						modelField := oneData["modelField"]
						bigMap[modelField.(string)] = labelName.(string)
					}
					fmt.Println(567)
				} else {
					log.Error("propertys 竟然不是列表")
				}
			}
			if mapKey != "" && bigMap != nil {
				retMap[mapKey] = bigMap
			}
		}
	}
	return retMap
}
