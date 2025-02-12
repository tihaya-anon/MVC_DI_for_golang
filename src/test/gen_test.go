package test

import (
	"MVC_DI/config"
	"MVC_DI/util/gen"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Test_Gen(t *testing.T) {
	config.InitConfig()
	gen.Generate("MVC_DI", []string{"user"})
}
func Test_GenQuery(t *testing.T) {
	config.InitConfig()
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		t.Errorf("connect to database failed: %v", err)
	}
	gen.GenerateQuery([]string{"user"}, gormDB)
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
