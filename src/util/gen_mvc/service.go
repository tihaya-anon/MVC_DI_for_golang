package gen_mvc

import (
	"MVC_DI/util"
	"log"
	"os"
	"strings"
	"text/template"
)

// generateService 生成 Service 和 ServiceImpl
func GenerateService(pkg, basePath string, tables []string) {
	for _, table := range tables {
		_generateService(pkg, basePath, table)
	}
}

func _generateService(pkg, basePath, table string) {
	servicePath := basePath + "/service"
	serviceImplPath := servicePath + "/impl"

	util.CreateDir(servicePath)
	util.CreateDir(serviceImplPath)
	// 生成 Service 接口
	serviceFile := servicePath + "/" + table + "_service.go"
	serviceTemplate := `package service

type {{.TableName}}Service interface {
	// DEFINE METHODS
}
`
	tmpl := template.Must(template.New("service").Parse(serviceTemplate))
	file, err := os.Create(serviceFile)
	if err != nil {
		log.Fatalf("create service file failed: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]string{
		"TableName": util.SnakeToPascal(table),
	}); err != nil {
		log.Fatalf("generate service failed: %v", err)
	}

	// 生成 ServiceImpl
	serviceImplFile := serviceImplPath + "/" + table + "_service_impl.go"
	serviceImplTemplate := `package impl

import (
	"{{.Pkg}}/section/{{.Entity}}/service"
)

type {{.TableName}}ServiceImpl struct{}

var _ service.{{.TableName}}Service = (*{{.TableName}}ServiceImpl)(nil)

// IMPLEMENT METHODS
`
	tmpl = template.Must(template.New("serviceImpl").Parse(serviceImplTemplate))
	file, err = os.Create(serviceImplFile)
	if err != nil {
		log.Fatalf("create service impl file failed: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]string{
		"TableName": util.SnakeToPascal(table),
		"Entity":    strings.Split(table, "_")[0],
		"Pkg":       pkg,
	}); err != nil {
		log.Fatalf("generate service impl failed: %v", err)
	}
}
