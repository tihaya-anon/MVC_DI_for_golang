package user_service_builder

import (
	user_service "MVC_DI/section/user/service"
	user_service_impl "MVC_DI/section/user/service/impl"
	user_mapper "MVC_DI/section/user/mapper"
)

func (b *UserAuthServiceBuilder) Build() user_service.UserAuthService {
	if b.userAuthServiceImpl.UserAuthMapper == nil && b.isStrict {
		panic("`UserAuthMapper` is required")
	}
	return b.userAuthServiceImpl
}

func (b *UserAuthServiceBuilder) WithUserAuthMapper(mapper user_mapper.UserAuthMapper) *UserAuthServiceBuilder {
	b.userAuthServiceImpl.UserAuthMapper = mapper
	return b
}

// BUILDER
type UserAuthServiceBuilder struct {
  isStrict bool
	userAuthServiceImpl *user_service_impl.UserAuthServiceImpl
}

func NewUserAuthServiceBuilder() *UserAuthServiceBuilder {
	return &UserAuthServiceBuilder{
		userAuthServiceImpl: &user_service_impl.UserAuthServiceImpl{},
	}
}

func (b *UserAuthServiceBuilder) UseStrict() *UserAuthServiceBuilder { 
  b.isStrict = true
  return b
}