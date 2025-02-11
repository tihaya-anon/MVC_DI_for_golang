package resp

import (
	"MVC_DI/global/enum"
	"MVC_DI/util/stream"
)

type IResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg,omitempty"`
	Data any    `json:"data,omitempty"`
}

func NewResponse() *IResponse {
	return &IResponse{}
}
func (r *IResponse) Success() *IResponse {
	r.Code = enum.CODE.SUCCESS
	r.Msg = enum.MSG.SUCCESS
	return r
}

func (r *IResponse) SuccessWithData(data any) *IResponse {
	r = r.Success()
	r.Data = data
	return r
}

func (r *IResponse) Error(code string, error error) *IResponse {
	var msg string
	if error != nil {
		msg = error.Error()
	} else {
		msg = enum.MSG.SYSTEM_ERROR
	}
	r.Code = code
	r.Msg = msg
	return r
}

func (r *IResponse) SystemError(error error) *IResponse {
	return r.Error(enum.CODE.SYSTEM_ERROR, error)
}

func (r *IResponse) CustomerError(error error) *IResponse {
	return r.Error(enum.CODE.CUSTOMER_ERROR, error)
}

func (r *IResponse) ThirdPartyError(error error) *IResponse {
	return r.Error(enum.CODE.THIRD_PARTY_ERROR, error)
}

func (r *IResponse) ValidationError(errorMap map[string]error) *IResponse {
	r.Code = enum.CODE.VALIDATION_ERROR
	r.Msg = enum.MSG.VALIDATION_ERROR
	r.Data = stream.NewMapStream(errorMap).
		Map(
			func(key string, val error) (string, any) {
				return key, val.Error()
			},
		).
		ToMap()
	return r
}
