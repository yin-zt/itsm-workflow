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

}
