package gen_mvc

import (
	"MVC_DI/util"
	"log"
	"os"
	"text/template"
)

// generateMapper 生成 Mapper 和 MapperImpl
func GenerateMapper(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateMapper(pkg, basePath, entity, table)
	}
}

func _generateMapper(pkg, basePath, entity, table string) {
	mapperPath := basePath + "/mapper"
	mapperImplPath := mapperPath + "/impl"

	util.CreateDir(mapperPath)
	util.CreateDir(mapperImplPath)
	// 生成 Mapper 接口
	mapperFile := mapperPath + "/" + table + "_mapper.go"
	mapperTemplate := `package mapper

type {{.TableName}}Mapper interface {
	// DEFINE METHODS
}
`
	tmpl := template.Must(template.New("mapper").Parse(mapperTemplate))
	file, err := os.Create(mapperFile)
	if err != nil {
		log.Fatalf("create mapper file failed: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]string{
		"TableName": util.SnakeToPascal(table),
	}); err != nil {
		log.Fatalf("generate mapper failed: %v", err)
	}

	// 生成 MapperImpl
	mapperImplFile := mapperImplPath + "/" + table + "_mapper_impl.go"
	mapperImplTemplate := `package impl

import (
	"{{.pkg}}/section/{{.entity}}/mapper"
)

type {{.TableName}}MapperImpl struct{}

func New{{.TableName}}MapperImpl() *{{.TableName}}MapperImpl {
	return &{{.TableName}}MapperImpl{}
}

var _ mapper.{{.TableName}}Mapper = (*{{.TableName}}MapperImpl)(nil)

// IMPLEMENT METHODS
`
	tmpl = template.Must(template.New("mapperImpl").Parse(mapperImplTemplate))
	file, err = os.Create(mapperImplFile)
	if err != nil {
		log.Fatalf("create mapper impl file failed: %v", err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]string{
		"TableName": util.SnakeToPascal(table),
		"entity":    entity,
		"pkg":       pkg,
	}); err != nil {
		log.Fatalf("generate mapper impl failed: %v", err)
	}
}
