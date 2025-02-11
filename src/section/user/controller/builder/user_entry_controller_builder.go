package user_controller_builder

import (
  user_service "MVC_DI/section/user/service"
  user_controller "MVC_DI/section/user/controller"
)

func (b *UserEntryControllerBuilder) Build() *user_controller.UserEntryController {
  if b.userEntryController.UserEntryService == nil && b.isStrict {
    panic("`UserEntryService` is required")
  }
  return b.userEntryController
}

func (b *UserEntryControllerBuilder) WithUserEntryService(userEntryService user_service.UserEntryService) *UserEntryControllerBuilder {
  b.userEntryController.UserEntryService = userEntryService
  return b
}


// BUILDER
type UserEntryControllerBuilder struct {
  isStrict bool
  userEntryController *user_controller.UserEntryController
}

func NewUserEntryControllerBuilder() *UserEntryControllerBuilder {
  return &UserEntryControllerBuilder{
    isStrict: false,
    userEntryController: &user_controller.UserEntryController{},
  }
}

func (b *UserEntryControllerBuilder) UseStrict() *UserEntryControllerBuilder { 
  b.isStrict = true
  return b
}