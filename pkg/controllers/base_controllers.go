package controller

import (
	"github.com/eryajf/xirang/logic"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"

	"github.com/gin-gonic/gin"
)

type BaseController struct{}

// GetPasswd 生成加密密码
func (m *BaseController) GetPasswd(c *gin.Context) {
	req := new(request.GetOrderInfo)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Base.GetPasswd(c, req)
	})
}
