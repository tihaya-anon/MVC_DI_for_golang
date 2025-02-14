package middleware

import "github.com/gin-gonic/gin"

type IProxy interface {
	Before(ctx *gin.Context)
	After(ctx *gin.Context)
}

type ProxyMiddleware struct {
	Proxy IProxy
}

func (middleware *ProxyMiddleware) ProxyVoid(ctx *gin.Context, fn func()) {
	if middleware.Proxy == nil {
		fn()
		return
	}
	middleware.Proxy.Before(ctx)
	fn()
	middleware.Proxy.After(ctx)
}

func (middleware *ProxyMiddleware) ProxyReturn(ctx *gin.Context, fn func() any) any {
	if middleware.Proxy == nil {
		return fn()
	}
	middleware.Proxy.Before(ctx)
	defer middleware.Proxy.After(ctx)
	return fn()
}
