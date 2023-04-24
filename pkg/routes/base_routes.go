package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/controllers"
)

// InitBaseRoutes 注册基础路由
func InitBaseRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/base")
	{
		base.GET("/ping", controllers.Demo)
		base.GET("/userinfo", controllers.Oa)
		base.GET("/userdepart", controllers.Department)
		base.GET("/gm", controllers.Gm)
	}
	return r
}
