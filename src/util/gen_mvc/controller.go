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
	"{{.pkg}}/section/{{.entity}}/service"
	"net/http"
)

type {{.TableName}}Controller struct {
	{{.TableName}}Service service.{{.TableName}}Service
}

func New{{.TableName}}Controller({{.tableName}}Service service.{{.TableName}}Service) *{{.TableName}}Controller {
	return &{{.TableName}}Controller{
		{{.TableName}}Service: {{.tableName}}Service,
	}
}

// RegisterRoutes
func (c *{{.TableName}}Controller) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/{{.table}}")
	{
		group.GET("/hello", func(ctx *gin.Context) {
			ctx.String(http.StatusOK, "hello {{.table}}")
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
		"entity":    entity,
		"TableName": util.SnakeToPascal(table),
		"tableName": util.SnakeToCamel(table),
		"table":     table,
		"pkg":       pkg,
	}); err != nil {
		log.Fatalf("generate controller failed: %v", err)
	}
}
