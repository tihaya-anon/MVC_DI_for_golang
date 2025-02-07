package test

import (
	"MVC_DI/config"
	"MVC_DI/util"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/google/uuid"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type CommonMethod struct {
	ID int64
}

func (commonMethod *CommonMethod) BeforeCreate(tx *gorm.DB) error {
	commonMethod.ID = int64(uuid.New().ID())
	return nil
}

// Generate 生成代码
func Generate(entities ...string) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to database failed: %v", err)
	}
	log.Printf("connect to: %v\n", config.Application.Database.Uri)

	for _, entity := range entities {
		basePath := "./src/section/" + entity
		// 生成 Mapper
		generateMapper(basePath, entity, gormDB)

		// 获取当前实体的所有表
		tables := getEntityTables(gormDB, entity)

		// 生成 Service 和 ServiceImpl
		generateService(basePath, tables)

		// 生成基于 Gin 的 Controller
		generateGinController(basePath, entity, tables)
	}
}

// getEntityTables 获取当前实体的所有表
func getEntityTables(db *gorm.DB, entity string) []string {
	var tables []string
	rows, err := db.Raw("SHOW TABLES LIKE '" + entity + "%'").Rows()
	if err != nil {
		log.Fatalf("get tables failed: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		var table string
		if err := rows.Scan(&table); err != nil {
			log.Fatalf("scan table failed: %v", err)
		}
		tables = append(tables, table)
	}
	return tables
}

// generateMapper 生成 Mapper
func generateMapper(basePath, entity string, gormDB *gorm.DB) {
	mapperPath := basePath + "/mapper"

	// 创建目录
	createDir(mapperPath)

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: mapperPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(gormDB)

	// 生成 Mapper 代码
	g.WithTableNameStrategy(func(tableName string) (targetTableName string) {
		if strings.HasPrefix(tableName, entity) {
			return tableName
		}
		return ""
	})
	// gen.WithMethod(CommonMethod{})
	g.ApplyBasic(g.GenerateAllTable(
		gen.FieldType("id", "int64"),
		gen.FieldJSONTag("id", "id"),
		gen.WithMethod(CommonMethod{}))...)
	g.Execute()
}

// generateService 生成 Service 和 ServiceImpl
func generateService(basePath string, tables []string) {
	for _, table := range tables {
		_generateService(basePath, table)
	}
}

func _generateService(basePath, table string) {
	servicePath := basePath + "/service"
	serviceImplPath := servicePath + "/impl"

	createDir(servicePath)
	createDir(serviceImplPath)
	// 生成 Service 接口
	serviceFile := servicePath + "/" + table + "_service.go"
	serviceTemplate := `package service

type {{.TableName}}Service interface {
	// 定义方法
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
	"forum_go/section/{{.Entity}}/service"
)

type {{.TableName}}ServiceImpl struct{}

var _ service.{{.TableName}}Service = (*{{.TableName}}ServiceImpl)(nil)

// 实现方法
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
	}); err != nil {
		log.Fatalf("generate service impl failed: %v", err)
	}
}

// generateGinController 生成基于 Gin 的 Controller
func generateGinController(basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateGinController(basePath, entity, table)
	}
}

func _generateGinController(basePath, entity, table string) {
	controllerPath := basePath + "/controller"
	createDir(controllerPath)

	controllerFile := controllerPath + "/" + entity + "_controller.go"
	controllerTemplate := `package controller

import (
	"github.com/gin-gonic/gin"
	"forum_go/section/{{.Entity}}/service"
	"forum_go/section/{{.Entity}}/service/impl"
	"net/http"
)

type {{.EntityName}}Controller struct {
	{{.TableName}}Service service.{{.TableName}}Service
}

func New{{.EntityName}}Controller() *{{.EntityName}}Controller {
	return &{{.EntityName}}Controller{
		{{.TableName}}Service: &impl.{{.TableName}}ServiceImpl{},
	}
}

// RegisterRoutes 注册路由
func (c *{{.EntityName}}Controller) RegisterRoutes(router *gin.Engine) {
	group := router.Group("/{{.Entity}}")
	{
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
		"Entity":     entity,
		"EntityName": cases.Title(language.English).String(entity),
		"TableName":  util.SnakeToPascal(table),
	}); err != nil {
		log.Fatalf("generate controller failed: %v", err)
	}
}

func createDir(path string) {
	if err := os.MkdirAll(path, os.ModePerm); err != nil {
		log.Fatalf("unable to create dir: %s, error: %v", path, err)
	}
}
