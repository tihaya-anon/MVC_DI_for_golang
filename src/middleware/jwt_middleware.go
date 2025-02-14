package middleware

import (
	"MVC_DI/global/enum"
	"MVC_DI/resp"
	"MVC_DI/security"
	"strings"

	"github.com/gin-gonic/gin"
)

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := resp.NewResponse()
		token := ctx.Request.Header.Get("Authorization")
		if token == "" {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.MISSING_TOKEN, enum.MSG.MISSING_TOKEN, nil))
			return
		}
		token = strings.Split(token, " ")[1]
		if !security.CheckJWT(token) {
			resp.ResponseWrapper(ctx, response.AllArgsConstructor(enum.CODE.INVALID_TOKEN, enum.MSG.INVALID_TOKEN, nil))
			return
		}
		ctx.Next()
	}
}
