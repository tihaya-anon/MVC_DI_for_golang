package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type IProxy interface {
	Before(ctx *gin.Context, logger *logrus.Logger)
	After(ctx *gin.Context, logger *logrus.Logger)
}

type LoggerProxyMiddleware struct {
	Proxy  IProxy
	Logger *logrus.Logger
}

func (middleware *LoggerProxyMiddleware) ProxyVoid(ctx *gin.Context, fn func()) {
	if middleware.Proxy == nil {
		fn()
		return
	}
	middleware.Proxy.Before(ctx, middleware.Logger)
	fn()
	middleware.Proxy.After(ctx, middleware.Logger)
}

func (middleware *LoggerProxyMiddleware) ProxyReturn(ctx *gin.Context, fn func() (any, error)) (any, error) {
	if middleware.Proxy == nil {
		return fn()
	}
	middleware.Proxy.Before(ctx, middleware.Logger)
	ret, err := fn()
	middleware.Proxy.After(ctx, middleware.Logger)
	return ret, err
}
