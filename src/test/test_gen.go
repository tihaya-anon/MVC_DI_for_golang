package test

import "MVC_DI/util/gen_mvc"

func TestGen() {
	gen_mvc.Generate("MVC_DI", []string{"user"})
}
