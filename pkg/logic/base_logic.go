package logic

import (
	"github.com/gin-gonic/gin"
	"github.com/yin-zt/itsm-workflow/pkg/models/request"
)

type OrderLogic struct{}

// GetOrderInfo
func (l OrderLogic) GetOrderInfo(c *gin.Context, req interface{}) (data interface{}, rspError interface{}) {
	r, ok := req.(*request.GetOrderInfo)
	if !ok {
		return nil, ReqAssertErr
	}
	_ = c

	return r, nil
}
