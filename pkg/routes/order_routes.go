package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/controllers"
)

// InitOrderRoutes 注册工单路由
func InitOrderRoutes(r *gin.RouterGroup) gin.IRoutes {
	order := r.Group("/order")
	{
		order.GET("/getorderinfo", controllers.Order.GetOrderInfo)
		order.POST("/collect", controllers.Order.CollectOrderInfo)
		order.POST("/callback", controllers.Order.OaCallBack)
		//order.POST("/add", controllers.Order.AddOrder)
	}
	return r
}
