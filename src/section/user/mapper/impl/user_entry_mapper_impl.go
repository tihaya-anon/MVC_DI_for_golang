package user_mapper_impl

import (
	user_mapper "MVC_DI/section/user/mapper"
)

type UserEntryMapperImpl struct{}

func NewUserEntryMapperImpl() *UserEntryMapperImpl {
	return &UserEntryMapperImpl{}
}

// IMPLEMENT METHODS



// INTERFACE
var _ user_mapper.UserEntryMapper = (*UserEntryMapperImpl)(nil)