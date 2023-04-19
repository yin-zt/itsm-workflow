package logic

import (
	"encoding/json"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/models/order"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
	"github.com/yin-zt/itsm-workflow/pkg/models/response"
	"github.com/yin-zt/itsm-workflow/pkg/utils/cmdb"
	"github.com/yin-zt/itsm-workflow/pkg/utils/isql"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
	"strconv"
	"strings"
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
	//displayStr := oneOrder.DisplayContent
	return response.OneOrderInfo{
		Title:    oneOrder.ApplyType,
		Type:     "default",
		Item:     []string{oneOrder.DisplayContent},
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
	if isql.Order.Exist(tools.H{"apply_logic_id": orderId}) {
		return nil, tools.NewValidatorError(fmt.Errorf("工单已存在,请勿重复添加"))
	}
	o.FindOutLabelKeys(r.UserTaskList, labelKeyMap)
	o.FindOutLabelVals(formData, labelValMap)

	dispalyContent := o.MakeDisplayContent(labelKeyMap, labelValMap)
	fmt.Println(labelValMap)
	fmt.Println(labelKeyMap)
	fmt.Println(dispalyContent)
	fmt.Println("mmmmmmmmmmmmmmmmmmmmmmmmm")

	order := order.Order{
		ApplyLogicId:   orderId,
		InstanceId:     instanceId,
		TaskId:         taskId,
		ApplyUser:      creator,
		ApplyType:      orderType,
		ApplyStatus:    0,
		ExecuteStatus:  0,
		DisplayContent: dispalyContent,
	}

	// 然后在数据库中创建组
	err := isql.Order.Add(&order)
	if err != nil {
		OpeLoger.Infof("向MySQL创建工单失败, 报错信息为：%v", err)
		return nil, tools.NewMySqlError(fmt.Errorf("向MySQL创建工单失败"))
	}
	return nil, nil
}

func (o OrderLogic) OaCallBack(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	var resp any
	defer func() {
		if err := recover(); err != resp {
			fmt.Println("捕获到了panic 产生的异常： ", err)
			fmt.Println("捕获到panic的异常了，recover并没有恢复回来")
			OpeLoger.Errorf("OaCallBack 捕获到panic异常，recover并没有恢复回来了，【err】为：%s", err)
		}
	}()

	r, ok := req.(*request.OaCallB)
	if !ok {
		return nil, ReqAssertErr
	}
	orderLogicId := c.PostForm("logic_id")
	orderFormCode := c.PostForm("code")
	orderCode, err := strconv.Atoi(orderFormCode)
	if err != nil {
		OpeLoger.Errorf("code 竟然不是整数, %v", c.PostForm("code"))
	}
	filter := tools.H{"apply_logic_id": orderLogicId}
	if !isql.Order.Exist(filter) {
		OpeLoger.Errorf("工单不存在，工单id为：%v", r.LogicId)
		return nil, tools.NewValidatorError(fmt.Errorf("工单不存在, 工单id为:%v", orderLogicId))
	}

	updateOrderInfo := order.Order{
		ApplyLogicId:  orderLogicId,
		ExecuteStatus: uint(orderCode),
	}
	err = isql.Order.Update(&updateOrderInfo)
	if err != nil {
		OpeLoger.Errorf("更新工单信息失败: %v", err)
		return nil, tools.NewMySqlError(fmt.Errorf("更新工单信息失败：%v", err))
	}
	callItsmRes, opeflag := o.SubmitItsmWorkflow(orderLogicId, orderFormCode)
	if opeflag {
		updateItsmInfo := order.Order{
			ApplyLogicId:  orderLogicId,
			ItsmOperation: 1,
			ItsmResult:    callItsmRes,
		}
		err = isql.Order.Update(&updateItsmInfo)
		if err != nil {
			OpeLoger.Errorf("调用itsm自动审批接口正常，更新工单信息失败: %v", err)
			return nil, tools.NewMySqlError(fmt.Errorf("调用itsm自动审批接口正常，更新工单信息失败：%v", err))
		} else {
			OpeLoger.Info("调用itsm自动审批接口正常，成功更新工单信息")
		}
	} else {
		updateItsmInfo := order.Order{
			ApplyLogicId:  orderLogicId,
			ItsmOperation: 2,
			ItsmResult:    callItsmRes,
		}
		err = isql.Order.Update(&updateItsmInfo)
		if err != nil {
			OpeLoger.Errorf("调用itsm自动审批接口异常，更新工单信息失败: %v", err)
			return nil, tools.NewMySqlError(fmt.Errorf("调用itsm自动审批接口异常，更新工单信息失败：%v", err))
		} else {
			OpeLoger.Info("调用itsm自动审批接口异常，成功更新工单信息")
		}
	}
	//isql.Order.Update()
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
		if task.FormDefinition == "" {
			continue
		}
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

// FindOutLabelVals 作用是解析工单任务中用户填写的字段值与字段id的关联关系，并以字典形式返回
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

// MakeDisplayContent 解析存在数据库中的字段id和字段值，并以字符串方式传回给oa详情接口
func (o OrderLogic) MakeDisplayContent(lableM map[string]map[string]string, contentM map[string][]map[string]interface{}) string {
	var resp any
	defer func() {
		if err := recover(); err != resp {
			fmt.Println("捕获到了panic 产生的异常： ", err)
			fmt.Println("捕获到panic的异常了，recover并没有恢复回来")
			OpeLoger.Errorf("MakeDisplayContent 捕获到panic异常，recover并没有恢复回来了，【err】为：%s", err)
		}
	}()
	var (
		//lableM    = make(map[string]interface{})
		//contentM  = make(map[string]interface{})
		resultStr = ""
	)
	//err := json.Unmarshal([]byte(labels), &lableM)
	//if err != nil {
	//	OpeLoger.Errorf("MakeDisplayContent 反序列化labels字符串异常，[err]: %v", err)
	//	return ""
	//}
	//err = json.Unmarshal([]byte(contents), &contentM)
	//if err != nil {
	//	OpeLoger.Errorf("MakeDisplayContent 反序列化labels字符串异常，[err]: %v", err)
	//	return ""
	//}
	for key, valMap := range lableM {
		contentLists, ok := contentM[key]
		if !ok {
			OpeLoger.Infof("此键: %v 并不存在于contentM字典: %v中", key, contentM)
			continue
		}
		if len(contentLists) == 0 {
			continue
		}
		for secKey, secValStr := range valMap {
			for _, itemInter := range contentLists {
				lableContent := itemInter[secKey]
				switch lableContentVal := lableContent.(type) {
				case string:
					resultStr = resultStr + "," + secValStr + ":" + lableContentVal
				case []string:
					if len(lableContentVal) == 0 {
						continue
					}
					listStr := strings.Join(lableContentVal, ";")
					//resultStr = resultStr + "\"" + secValStr + "\"" + ":" + "\"" + listStr + "\"" + ","
					resultStr = resultStr + "," + secValStr + ":" + listStr

				case []interface{}:
					var labelContent = ""
					for _, oneObj := range lableContentVal {
						oneObjVal, ok := oneObj.(map[string]interface{})
						if !ok {
							OpeLoger.Errorf("oneObj此变量并非字典，值为：%v", oneObj)
							continue
						}
						ipVal := oneObjVal["ip"]
						labelContent = labelContent + ";" + ipVal.(string)
					}
					resultStr = resultStr + "," + secValStr + ":" + strings.Trim(labelContent, ";")
				}
			}
		}
	}
	return strings.Trim(resultStr, ",")
}

func (o OrderLogic) SubmitItsmWorkflow(taskID, operateCode string) (string, bool) {
	var (
		tastinfos []string
	)
	tastinfos = strings.Split(taskID, "_")
	fmt.Println(tastinfos[0])
	fmt.Println(tastinfos[1])

	if operateCode == "20" {
		result, flag := cmdb.Cmdb.CloseItsmOrder(tastinfos[0], tastinfos[1])
		return result, flag
	} else if operateCode == "60" {
		result, flag := cmdb.Cmdb.AutoShenpi(tastinfos[0], tastinfos[1])
		fmt.Println(result, flag)
		return result, flag
	} else {
		OpeLoger.Errorf("operateCode 的值不是20或者60，而是错误的数值：%v", operateCode)
		return "", false
	}
}
