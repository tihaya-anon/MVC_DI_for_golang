package controller_uitl

import (
	"MVC_DI/vo/req"

	"github.com/gin-gonic/gin"
)

func BindPageReq(ctx *gin.Context) (*req.TPageReq, error) {
	pageReq := &req.TPageReq{}
	if err := ctx.ShouldBindQuery(pageReq); err != nil {
		return nil, err
	}
	return pageReq, nil
}
