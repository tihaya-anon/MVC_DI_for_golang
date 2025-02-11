package user_controller

import (
	"github.com/gin-gonic/gin"
	user_service "MVC_DI/section/user/service"
	"MVC_DI/resp"
)

type UserEntryController struct {
	UserEntryService user_service.UserEntryService
}

func (c *UserEntryController) Hello(ctx *gin.Context) *resp.IResponse {
	return resp.NewResponse().SuccessWithData("hello `user-entry`")
}