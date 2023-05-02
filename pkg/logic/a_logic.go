package logic

import (
	"fmt"
	"github.com/yin-zt/itsm-workflow/pkg/utils/levelDB"
	"github.com/yin-zt/itsm-workflow/pkg/utils/loger"
	"github.com/yin-zt/itsm-workflow/pkg/utils/oa"
	"github.com/yin-zt/itsm-workflow/pkg/utils/tools"
)

var (
	ReqAssertErr = tools.NewRspError(tools.SystemErr, fmt.Errorf("请求异常"))

	Order     = &OrderLogic{}
	Cmdb      = &CmdbLogic{}
	OpeLoger  = loger.GetLoggerOperate()
	CmdbLoger = loger.GetCmdbLogger()
	Oa        = &oa.ApproValApi{}
	Leveldb   = levelDB.NewLevelDb()
)
