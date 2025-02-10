package util

import (
	"MVC_DI/resp"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IResponseFunc = func(ctx *gin.Context) *resp.IResponse

func responseWrapper(fn IResponseFunc) func(ctx *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.AbortWithStatusJSON(http.StatusOK, fn(ctx))
	}
}

type IRoutesWrapper struct {
	Routes gin.IRoutes
}

func RoutesWrapper(routes gin.IRoutes) *IRoutesWrapper {
	return &IRoutesWrapper{
		Routes: routes,
	}
}

func (r *IRoutesWrapper) Use(middleware ...gin.HandlerFunc) gin.IRoutes {
	return r.Routes.Use(middleware...)
}

func iResponse2Handler(fn []IResponseFunc) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, len(fn))
	for i, fn := range fn {
		handlers[i] = responseWrapper(fn)
	}
	return handlers
}

func (r *IRoutesWrapper) GET(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.GET(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) POST(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.POST(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) DELETE(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.DELETE(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) PATCH(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.PATCH(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) PUT(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.PUT(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) OPTIONS(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.OPTIONS(relativePath, iResponse2Handler(responseFunc)...)
}

func (r *IRoutesWrapper) HEAD(relativePath string, responseFunc ...IResponseFunc) gin.IRoutes {
	return r.Routes.HEAD(relativePath, iResponse2Handler(responseFunc)...)
}
