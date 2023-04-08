package controllers

import (
	"github.com/yin-zt/itsm-workflow/pkg/logic"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"

	"github.com/gin-gonic/gin"
)

type OrderController struct{}

// GetPasswd 生成加密密码
func (m *OrderController) GetPasswd(c *gin.Context) {
	req := new(request.GetOrderInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Order.GetOrderInfo(c, req)
	})
}
