package user_service_builder

import (
	user_service "MVC_DI/section/user/service"
	user_service_impl "MVC_DI/section/user/service/impl"
	user_mapper "MVC_DI/section/user/mapper"
)

func (b *UserEntryServiceBuilder) Build() user_service.UserEntryService {
	if b.userEntryServiceImpl.UserEntryMapper == nil && b.isStrict {
		panic("`UserEntryMapper` is required")
	}
	return b.userEntryServiceImpl
}

func (b *UserEntryServiceBuilder) WithUserEntryMapper(mapper user_mapper.UserEntryMapper) *UserEntryServiceBuilder {
	b.userEntryServiceImpl.UserEntryMapper = mapper
	return b
}

// BUILDER
type UserEntryServiceBuilder struct {
  isStrict bool
	userEntryServiceImpl *user_service_impl.UserEntryServiceImpl
}

func NewUserEntryServiceBuilder() *UserEntryServiceBuilder {
	return &UserEntryServiceBuilder{
		userEntryServiceImpl: &user_service_impl.UserEntryServiceImpl{},
	}
}

func (b *UserEntryServiceBuilder) UseStrict() *UserEntryServiceBuilder { 
  b.isStrict = true
  return b
}