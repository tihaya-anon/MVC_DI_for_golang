package test

import (
	"MVC_DI/util/gen"
	"testing"
)

func Test_Gen(t *testing.T) {
	gen.Generate("MVC_DI", []string{"user"})
}

func Test_GenController(t *testing.T) {
	gen.GenerateGinController("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}

func Test_GenService(t *testing.T) {
	gen.GenerateService("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}

func Test_GenMapper(t *testing.T) {
	gen.GenerateMapper("MVC_DI", "section", "user", []string{"user_auth", "user_entry"})
}
