package gen_mvc

import (
	"MVC_DI/util"
	"log"
	"os"
	"text/template"
)

// generateGinController 生成基于 Gin 的 Controller
func GenerateGinController(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateGinController(pkg, basePath, entity, table)
	}
}

func _generateGinController(pkg, basePath, entity, table string) {
	controllerPath := basePath + "/controller"
	util.CreateDir(controllerPath)

	controllerFile := controllerPath + "/" + table + "_controller.go"
	controllerTemplate := `package controller

import (
	"github.com/gin-gonic/gin"
	"{{.Pkg}}/section/{{.Entity}}/service"
	"{{.Pkg}}/section/{{.Entity}}/service/impl"
	"net/http"
)

type {{.TableName}}Controller struct {
	{{.TableName}}Service service.{{.TableName}}Service
}

func New{{.TableName}}Controller() *{{.TableName}}Controller {
	return &{{.TableName}}Controller{
		{{.TableName}}Service: &impl.{{.TableName}}ServiceImpl{},
	}
}

// RegisterRoutes
func (c *{{.TableName}}Controller) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/{{.Table}}")
	{
		group.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "hello {{.Table}}")
		})
	}
}
`
	tmpl := template.Must(template.New("controller").Parse(controllerTemplate))
	file, err := os.Create(controllerFile)
	if err != nil {
		log.Fatalf("create controller file failed: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]interface{}{
		"Entity":    entity,
		"TableName": util.SnakeToPascal(table),
		"Table":     table,
		"Pkg":       pkg,
	}); err != nil {
		log.Fatalf("generate controller failed: %v", err)
	}
}
