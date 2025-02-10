package test

import (
	"MVC_DI/security"
	"fmt"
)

func Test_JWT() {
	type Claims struct {
		Name string
		Age  int
	}
	token, _ := security.GenerateJWT(
		Claims{
			Name: "test",
			Age:  10,
		},
	)
	fmt.Println(token)
	result, _ := security.ParseJWT[Claims](token)
	fmt.Println(result)
	fmt.Println(security.CheckJWT(token))
	token = token + "abcd"
	fmt.Println(token)
	fmt.Println(security.CheckJWT(token))
}
