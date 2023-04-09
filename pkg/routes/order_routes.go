package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/controllers"
)

// InitOrderRoutes 注册工单路由
func InitOrderRoutes(r *gin.RouterGroup) gin.IRoutes {
	base := r.Group("/order")
	{
		base.GET("getorderinfo", controllers.Order.GetOrderInfo)
		base.POST("collectorderinfo", controllers.Order.CollectOrderInfo)
	}
	return r
}
