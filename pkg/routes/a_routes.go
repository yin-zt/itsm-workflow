package routes

import (
	log "github.com/cihub/seelog"
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/config"
)

// 初始化
func InitRoutes() *gin.Engine {
	defer log.Flush()
	//设置模式
	//gin.SetMode(config.Conf.System.Mode)

	// 创建带有默认中间件的路由:
	// 日志与恢复中间件
	r := gin.Default()
	// 创建不带中间件的路由:
	// r := gin.New()
	// r.Use(gin.Recovery())

	// 启用限流中间件
	// 默认每50毫秒填充一个令牌，最多填充200个

	// 启用操作日志中间件
	//r.Use(middleware.OperationLogMiddleware())

	// 路由分组
	apiGroup := r.Group("/" + config.UrlPathPrefix)

	// 注册路由
	InitBaseRoutes(apiGroup)  // 注册基础路由
	InitOrderRoutes(apiGroup) // 注册工单路由

	log.Info("初始化路由完成！")
	return r
}
