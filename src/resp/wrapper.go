package resp

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerWrapper(fn func(ctx *gin.Context) *IResponse) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ResponseWrapper(ctx, fn(ctx))
	}
}

func ResponseWrapper(ctx *gin.Context, response *IResponse) {
	ctx.AbortWithStatusJSON(http.StatusOK, response)
}
