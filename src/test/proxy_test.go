package test

import (
	"MVC_DI/global/log"
	"MVC_DI/middleware"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type MockOrigin struct {
	ProxyMiddleware *middleware.LoggerProxyMiddleware
	Logger          *logrus.Logger
}

func (m *MockOrigin) Hello(ctx *gin.Context) {
	m.ProxyMiddleware.ProxyVoid(ctx, func() {
		m.Logger.Info("Hello")
	})
}

type MockOriginBuilder struct {
	mockOrigin *MockOrigin
}

func NewMockOriginBuilder() *MockOriginBuilder {
	return &MockOriginBuilder{
		mockOrigin: &MockOrigin{
			ProxyMiddleware: &middleware.LoggerProxyMiddleware{},
		},
	}
}

func (builder *MockOriginBuilder) WithProxy(proxy middleware.IProxy) *MockOriginBuilder {
	builder.mockOrigin.ProxyMiddleware.Proxy = proxy
	return builder
}

func (builder *MockOriginBuilder) WithLogger(logger *logrus.Logger) *MockOriginBuilder {
	builder.mockOrigin.Logger = logger
	builder.mockOrigin.ProxyMiddleware.Logger = logger
	return builder
}

func (builder *MockOriginBuilder) Build() *MockOrigin {
	return builder.mockOrigin
}

type MockProxy struct{}

func (proxy *MockProxy) Before(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("Before")
}
func (proxy *MockProxy) After(ctx *gin.Context, logger *logrus.Logger) {
	logger.Info("After")
}

func Test_Proxy(t *testing.T) {
	builder := NewMockOriginBuilder()
	origin := builder.WithProxy(&MockProxy{}).WithLogger(log.GetLogger()).Build()
	origin.Hello(nil)
}
