package test

import (
	"MVC_DI/util/gen"
)

func TestGen() {
	gen.Generate("MVC_DI", []string{"user"})
}

func TestGenController() {
	gen.GenerateGinController("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}

func TestGenService() {
	gen.GenerateService("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}

func TestGenMapper() {
	gen.GenerateMapper("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}
