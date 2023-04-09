package logic

import (
	"fmt"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	Order = &OrderLogic{}
)