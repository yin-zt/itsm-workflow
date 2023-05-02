package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/logic"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
)

type CmdbController struct{}

// GetUserAllResources workflow获取用户的关联模型字段
func (m *CmdbController) GetUserAllResources(c *gin.Context) {
	req := new(request.UserResourceReq)
	Run(c, req, func() (interface{}, interface{}) {
		return logic.Cmdb.GetUserAllResources(c, req)
	})
}
