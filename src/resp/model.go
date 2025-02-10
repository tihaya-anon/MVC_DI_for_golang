package resp

import (
	"MVC_DI/global/enum"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IBody struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

type IResponse struct {
	ctx *gin.Context
}

func NewResponse(ctx *gin.Context) *IResponse {
	return &IResponse{ctx: ctx}
}

func (r *IResponse) Base(body IBody) {
	r.ctx.AbortWithStatusJSON(http.StatusOK, body)
}
func (r *IResponse) Success() {
	r.Base(IBody{
		Code: enum.CODE.SUCCESS,
		Msg:  enum.MSG.SUCCESS,
	})
}
func (r *IResponse) SuccessWithData(data any) {
	r.Base(IBody{
		Code: enum.CODE.SUCCESS,
		Msg:  enum.MSG.SUCCESS,
		Data: data,
	})
}