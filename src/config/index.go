package config

import "MVC_DI/config/model"

var Application = &model.Application{}

func init() {
	Parse("application.yaml", Application)
	Resolve(Application)
}