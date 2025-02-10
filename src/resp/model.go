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

func (r *IResponse) Error(code string, error error) {
	var msg string
	if error != nil {
		msg = error.Error()
	} else {
		msg = enum.MSG.SYSTEM_ERROR
	}
	r.Base(IBody{
		Code: code,
		Msg:  msg,
	})
}

func (r *IResponse) SystemError(error error) {
	r.Error(enum.CODE.SYSTEM_ERROR, error)
}

func (r *IResponse) CustomerError(error error) {
	r.Error(enum.CODE.CUSTOMER_ERROR, error)
}

func (r *IResponse) ThirdPartyError(error error) {
	r.Error(enum.CODE.THIRD_PARTY_ERROR, error)
}
