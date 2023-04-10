package main

import (
	"context"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/yin-zt/itsm-workflow/pkg/config"
	"github.com/yin-zt/itsm-workflow/pkg/routes"
	"github.com/yin-zt/itsm-workflow/pkg/utils/common"
	"github.com/yin-zt/itsm-workflow/pkg/utils/loger"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func init() {
	defer log.Flush()
	for _, dir := range config.ConstSysDirs {
		os.Mkdir(dir, 0777)
	}
	log.ReplaceLogger(loger.GetLoggerOperate())
	log.Info("test")
}

func main() {
	defer log.Flush()

	common.InitMysql()

	// 注册所有路由
	r := routes.InitRoutes()

	host := "0.0.0.0"
	port := config.ServicePort

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Errorf("listen: %s\n", err)
			return
		}
	}()

	log.Info(fmt.Sprintf("Server is running at %s:%d/%s", host, port, config.UrlPathPrefix))

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	// signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Info("Shutting down server...")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Errorf("Server forced to shutdown:", err)
		return
	}
	log.Info("Server exiting!")
}
