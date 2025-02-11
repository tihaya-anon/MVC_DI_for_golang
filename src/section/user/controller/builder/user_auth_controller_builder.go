package user_controller_builder

import (
  user_service "MVC_DI/section/user/service"
  user_controller "MVC_DI/section/user/controller"
)

func (b *UserAuthControllerBuilder) Build() *user_controller.UserAuthController {
  if b.userAuthController.UserAuthService == nil && b.isStrict {
    panic("`UserAuthService` is required")
  }
  return b.userAuthController
}

func (b *UserAuthControllerBuilder) WithUserAuthService(userAuthService user_service.UserAuthService) *UserAuthControllerBuilder {
  b.userAuthController.UserAuthService = userAuthService
  return b
}


// BUILDER
type UserAuthControllerBuilder struct {
  isStrict bool
  userAuthController *user_controller.UserAuthController
}

func NewUserAuthControllerBuilder() *UserAuthControllerBuilder {
  return &UserAuthControllerBuilder{
    isStrict: false,
    userAuthController: &user_controller.UserAuthController{},
  }
}

func (b *UserAuthControllerBuilder) UseStrict() *UserAuthControllerBuilder { 
  b.isStrict = true
  return b
}