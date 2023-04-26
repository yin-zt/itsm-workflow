package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/controllers"
)

// InitCmdbRoutes 注册CMDB路由
func InitCmdbRoutes(r *gin.RouterGroup) gin.IRoutes {
	order := r.Group("/cmdb5")
	{
		order.POST("/get_user_resources", controllers.Cmdb.GetUserResources)
		//order.POST("/collect", controllers.Order.CollectOrderInfo)
		//order.POST("/callback", controllers.Order.OaCallBack)
		//order.POST("/add", controllers.Order.AddOrder)
	}
	return r
}
