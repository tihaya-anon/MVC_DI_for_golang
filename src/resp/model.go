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
func (response *IResponse) Success() *IResponse {
	response.Code = enum.CODE.SUCCESS
	response.Msg = enum.MSG.SUCCESS
	return response
}

func (response *IResponse) SuccessWithData(data any) *IResponse {
	response = response.Success()
	response.Data = data
	return response
}

func (response *IResponse) Error(code string, error error) *IResponse {
	var msg string
	if error != nil {
		msg = error.Error()
	} else {
		msg = enum.MSG.SYSTEM_ERROR
	}
	response.Code = code
	response.Msg = msg
	return response
}

func (response *IResponse) SystemError(error error) *IResponse {
	return response.Error(enum.CODE.SYSTEM_ERROR, error)
}

func (response *IResponse) CustomerError(error error) *IResponse {
	return response.Error(enum.CODE.CUSTOMER_ERROR, error)
}

func (response *IResponse) ThirdPartyError(error error) *IResponse {
	return response.Error(enum.CODE.THIRD_PARTY_ERROR, error)
}

func (response *IResponse) ValidationError(errorMap map[string]error) *IResponse {
	response.Code = enum.CODE.VALIDATION_ERROR
	response.Msg = enum.MSG.VALIDATION_ERROR
	response.Data = stream.NewMapStream(errorMap).
		Map(
			func(key string, val error) (string, any) {
				return key, val.Error()
			},
		).
		ToMap()
	return response
}

func (response *IResponse) AllArgsConstructor(code string, msg string, data any) *IResponse {
	response.Code = code
	response.Msg = msg
	response.Data = data
	return response
}
