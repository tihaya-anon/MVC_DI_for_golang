package gen

import (
	"MVC_DI/util"
	"log"
	"os"
	"path"
	"text/template"
)

func GenerateTemplate(pkg, templatePath, targetPath, entity, table string) {
	templateFile, err := os.ReadFile(templatePath)
	if err != nil {
		log.Fatalf("read `%v` failed: %v", templatePath, err)
	}
	coreTemplate := string(templateFile)

	tmpl := template.Must(template.New("template").Parse(coreTemplate))
	util.CreateDir(targetPath)

	targetFile := path.Join(targetPath, table+".go")
	file, err := os.Create(targetFile)
	if err != nil {
		log.Fatalf("create `%v` file failed: %v", targetFile, err)
	}
	defer file.Close()

	if err := tmpl.Execute(file, map[string]interface{}{
		"entity_name": entity,
		"TableName":   util.SnakeToPascal(table),
		"tableName":   util.SnakeToCamel(table),
		"table_name":  table,
		"table_name_hyphen":  util.SnakeToHyphen(table),
		"pkg":         pkg,
	}); err != nil {
		log.Fatalf("generate `%v` failed: %v", targetPath, err)
	}
}
