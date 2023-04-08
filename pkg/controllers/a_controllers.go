package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
	"net/http"
)

func Demo(c *gin.Context) {
	CodeDebug()
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func CodeDebug() {
}
