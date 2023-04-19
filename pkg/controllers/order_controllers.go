package controllers

import (
	"github.com/yin-zt/itsm-workflow/pkg/logic"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

// GetOrderInfo workflow获取工单详细数据
func (m *OrderController) GetOrderInfo(c *gin.Context) {
	req := new(request.GetOrderInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Order.GetOrderInfo(c, req)
	})
}

// CollectOrderInfo 用于收集cmdb-itsm推送的工单信息
func (m *OrderController) CollectOrderInfo(c *gin.Context) {
	req := new(request.OrderInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Order.AnalyOrderInfo(c, req)
	})
}

// OaCallBack 用于接收OA系统
func (m *OrderController) OaCallBack(c *gin.Context) {
	req := new(request.OaCallB)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Order.OaCallBack(c, req)
	})
}

// AddOrder 用于在数据库中增加一条工单记录
//func (m *OrderController) AddOrder(c *gin.Context) {
//	req := new(request.OrderAddReq)
//	Run(c, req, func() (interface{}, interface{}) {
//		return logic.Order.AddOrderRecord(c, req)
//	})
//}
