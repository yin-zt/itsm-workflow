package logic

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/models/order"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
	"github.com/yin-zt/itsm-workflow/pkg/models/response"
	"github.com/yin-zt/itsm-workflow/pkg/utils/isql"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
)

type OrderLogic struct{}

// GetOrderInfo
func (o OrderLogic) GetOrderInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	var resp any
	defer func() {
		if err := recover(); err != resp {
			fmt.Println("捕获到了panic 产生的异常： ", err)
			fmt.Println("捕获到panic的异常了，recover并没有恢复回来")
			OpeLoger.Errorf("GetOrderInfo 捕获到panic异常，recover并没有恢复回来了，【err】为：%s", err)
		}
	}()
	r, ok := req.(*request.GetOrderInfo)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c
	filter := tools.H{"apply_logic_id": r.LogicId}

	if !isql.Order.Exist(filter) {
		OpeLoger.Errorf("工单信息不存在，工单ID为：%v", r.LogicId)
		return nil, tools.NewMySqlError(fmt.Errorf("工单信息不存在"))
	}

	oneOrder := new(order.Order)
	err := isql.Order.Find(filter, oneOrder)
	if err != nil {
		OpeLoger.Error("获取工单详细信息失败: %s", err.Error())
		return nil, tools.NewMySqlError(fmt.Errorf("获取工单详细信息失败: %s", err.Error()))
	}
	labelVal := oneOrder.DisplayLabel
	contentVal := oneOrder.DisplayContent
	fmt.Println(labelVal)
	fmt.Println(contentVal)
	return response.OneOrderInfo{
		Title:    oneOrder.ApplyType,
		Type:     "default",
		IsExpand: true,
	}, nil
}

func (o OrderLogic) AnalyOrderInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	var resp any
	defer func() {
		if err := recover(); err != resp {
			fmt.Println("捕获到了panic 产生的异常： ", err)
			fmt.Println("捕获到panic的异常了，recover并没有恢复回来")
			OpeLoger.Errorf("AnalyOrderInfo 捕获到panic异常，recover并没有恢复回来了，【err】为：%s", err)
		}
	}()
	var (
		taskId      string
		instanceId  string
		creator     string
		labelKeyMap = make(map[string]map[string]string)
		labelValMap = make(map[string][]map[string]interface{})
	)
	r, ok := req.(*request.OrderInfo)
	if !ok {
		return nil, ReqAssertErr
	}
	stepLists := r.StepList
	for _, item := range stepLists {
		if item.Status == "running" {
			taskId = item.InstanceId
		}
	}

	if r.ProcessInstance.Status != "running" {
		log.Errorf("此工单状态为非流转,工单ID:%s, 任务ID:%s", instanceId, taskId)
		return nil, nil
	}
	instanceId = r.ProcessInstance.InstanceId
	creator = r.ProcessInstance.Creator
	orderType := r.ProcessInstance.Name
	orderId := fmt.Sprintf("%s_%s", instanceId, taskId)
	formData := r.FormData
	fmt.Println(orderId)
	if isql.Order.Exist(tools.H{"apply_logic_id": orderId}) {
		return nil, tools.NewValidatorError(fmt.Errorf("工单已存在,请勿重复添加"))
	}
	o.FindOutLabelKeys(r.UserTaskList, labelKeyMap)
	jsonKey, err := json.Marshal(labelKeyMap)
	if err != nil {
		OpeLoger.Infof("labelKeyMap 解析异常，值为：%v", labelKeyMap)
		jsonKey = nil
	}

	o.FindOutLabelVals(formData, labelValMap)
	jsonVal, err := json.Marshal(labelValMap)
	if err != nil {
		OpeLoger.Infof("labelValMap 解析异常，值为：%v", labelValMap)
		jsonVal = nil
	}

	order := order.Order{
		ApplyLogicId:   orderId,
		InstanceId:     instanceId,
		TaskId:         taskId,
		ApplyUser:      creator,
		ApplyType:      orderType,
		ApplyStatus:    0,
		ExecuteStatus:  0,
		DisplayLabel:   string(jsonKey),
		DisplayContent: string(jsonVal),
	}

	// 然后在数据库中创建组
	err = isql.Order.Add(&order)
	if err != nil {
		OpeLoger.Infof("向MySQL创建工单失败, 报错信息为：%v", err)
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建工单失败"))
	}
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

// FindOutLabelKeys 作用是解析工单任务中字段名称与字段id的关联关系，并以字典的形式返回。
// 字典最外层的key是容器的id，里面的字典key为字段id，值为 字段的中文名
func (o OrderLogic) FindOutLabelKeys(tasks []request.UserTask, retMap map[string]map[string]string) {
	for _, task := range tasks {
		var m2 []map[string]interface{}
		err := json.Unmarshal([]byte(task.FormDefinition), &m2)
		if err != nil {
			log.Errorf("解析task.FormDefinition字符串内容失败，错误：%s", err)
			return
		}
		if len(m2) == 0 {
			continue
		}

		for _, innerMap := range m2 {
			var mapKey = ""
			var bigMap = map[string]string{}
			if labelName, ok := innerMap["key"]; ok {
				mapKey = labelName.(string)
				retMap[mapKey] = map[string]string{}
			} else {
				continue
			}
			if propertys, ok := innerMap["propertys"]; ok {
				if propertysLists, ok := propertys.([]interface{}); ok {
					for _, oneProperty := range propertysLists {
						oneData := oneProperty.(map[string]interface{})
						labelName := oneData["label"]
						modelField := oneData["modelField"]
						bigMap[modelField.(string)] = labelName.(string)
					}
				} else {
					log.Error("propertys 竟然不是列表")
				}
			}
			if mapKey != "" && bigMap != nil {
				retMap[mapKey] = bigMap
			}
		}
	}
}

// FindOutLabelVals
func (o OrderLogic) FindOutLabelVals(formData string, kepMap map[string][]map[string]interface{}) {
	var m3 []map[string]interface{}
	fmt.Println(formData)
	err := json.Unmarshal([]byte(formData), &m3)
	if err != nil {
		log.Errorf("解析formData内容失败，报错为:%s", err)
		return
	}
	for _, oneFormVal := range m3 {
		var (
			mapList = []map[string]interface{}{}
		)
		keyStrId := oneFormVal["key"]
		keyValues := oneFormVal["values"]
		keyValuesLists := keyValues.([]interface{})

		for _, item := range keyValuesLists {
			itemVal := item.(map[string]interface{})
			mapList = append(mapList, itemVal)
		}
		kepMap[keyStrId.(string)] = mapList
	}
}
