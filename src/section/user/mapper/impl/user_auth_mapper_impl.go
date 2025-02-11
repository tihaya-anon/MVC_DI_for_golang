package user_mapper_impl

import (
	user_mapper "MVC_DI/section/user/mapper"
)

type UserAuthMapperImpl struct{}

func NewUserAuthMapperImpl() *UserAuthMapperImpl {
	return &UserAuthMapperImpl{}
}

// IMPLEMENT METHODS



// INTERFACE
var _ user_mapper.UserAuthMapper = (*UserAuthMapperImpl)(nil)