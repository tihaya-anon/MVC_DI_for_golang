package user_service_impl

import (
	user_service "MVC_DI/section/user/service"
	user_mapper "MVC_DI/section/user/mapper"
)

type UserEntryServiceImpl struct{
	UserEntryMapper user_mapper.UserEntryMapper
}

// IMPLEMENT METHODS



// INTERFACE
var _ user_service.UserEntryService = (*UserEntryServiceImpl)(nil)

