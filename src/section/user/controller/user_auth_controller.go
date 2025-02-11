package user_controller

import (
	"github.com/gin-gonic/gin"
	user_service "MVC_DI/section/user/service"
	"MVC_DI/resp"
)

type UserAuthController struct {
	UserAuthService user_service.UserAuthService
}

func (c *UserAuthController) Hello(ctx *gin.Context) *resp.IResponse {
	return resp.NewResponse().SuccessWithData("hello `user-auth`")
}