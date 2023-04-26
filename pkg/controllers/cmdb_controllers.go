package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/logic"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
)

type CmdbController struct{}

// GetUserResources workflow获取工单详细数据
func (m *CmdbController) GetUserResources(c *gin.Context) {
	req := new(request.UserResourceReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Cmdb.GetUserResource(c, req)
	})
}
