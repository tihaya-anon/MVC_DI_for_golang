package user_service_impl

import (
	user_service "MVC_DI/section/user/service"
	user_mapper "MVC_DI/section/user/mapper"
)

type UserAuthServiceImpl struct{
	UserAuthMapper user_mapper.UserAuthMapper
}

// IMPLEMENT METHODS



// INTERFACE
var _ user_service.UserAuthService = (*UserAuthServiceImpl)(nil)

