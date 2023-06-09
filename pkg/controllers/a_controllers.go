package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/yin-zt/itsm-workflow/pkg/utils/oa"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
	"net/http"
)

var (
	Order = &OrderController{}
	Cmdb  = &CmdbController{}

	validate = validator.New()
	trans    ut.Translator
)

func Run(c *gin.Context, req interface{}, fn func() (interface{}, interface{})) {
	var err error
	// bind struct
	err = c.Bind(req)
	if err != nil {
		tools.Err(c, tools.NewValidatorError(err), nil)
		return
	}
	// 校验
	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			tools.Err(c, tools.NewValidatorError(fmt.Errorf(err.Translate(trans))), nil)
			return
		}
	}
	data, err1 := fn()
	if err1 != nil {
		tools.Err(c, tools.ReloadErr(err1), data)
		return
	}
	tools.Success(c, data)
}

func RunOnce(c *gin.Context, req interface{}, fn func() (interface{}, interface{})) {
	var err error
	// bind struct
	err = c.Bind(req)
	if err != nil {
		tools.Err(c, tools.NewValidatorError(err), nil)
		return
	}
	// 校验
	err = validate.Struct(req)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			tools.Err(c, tools.NewValidatorError(fmt.Errorf(err.Translate(trans))), nil)
			return
		}
	}
	data, err1 := fn()
	if err1 != nil {
		tools.Err(c, tools.ReloadErr(err1), data)
		return
	}
	tools.SuccessOnce(c, data)
}

func Demo(c *gin.Context) {
	CodeDebug()
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": "pong"})
}

func Oa(c *gin.Context) {
	commonOa := oa.ApproValApi{}
	responseData := commonOa.GetCompanyId("mailongyin")
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": responseData})
}

func Department(c *gin.Context) {
	commonOa := oa.ApproValApi{}
	responseData := commonOa.GetPersonDepart("mailongyin")
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": responseData})
}

func Gm(c *gin.Context) {
	commonOa := oa.ApproValApi{}
	responseData := commonOa.GetGmInfo("52434")
	c.JSON(http.StatusOK, tools.H{"code": 200, "msg": "ok", "data": responseData})
}

func CodeDebug() {
}
