package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/controllers"
)

// 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("ping", controllers.Demo)
		base.GET("getorderinfo", controllers.Base.GetPasswd)
	}
	return r
}
