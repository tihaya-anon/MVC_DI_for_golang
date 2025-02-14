package util

import (
	"MVC_DI/resp"

	"github.com/gin-gonic/gin"
)

type IHandlerFunc = func(ctx *gin.Context) *resp.IResponse

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

func iHandler2GinHandler(fn []IHandlerFunc) []gin.HandlerFunc {
	handlers := make([]gin.HandlerFunc, len(fn))
	for i, fn := range fn {
		handlers[i] = resp.HandlerWrapper(fn)
	}
	return handlers
}

func (r *IRoutesWrapper) GET(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.GET(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) POST(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.POST(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) DELETE(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.DELETE(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) PATCH(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.PATCH(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) PUT(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.PUT(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) OPTIONS(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.OPTIONS(relativePath, iHandler2GinHandler(responseFunc)...)
}

func (r *IRoutesWrapper) HEAD(relativePath string, responseFunc ...IHandlerFunc) gin.IRoutes {
	return r.Routes.HEAD(relativePath, iHandler2GinHandler(responseFunc)...)
}
