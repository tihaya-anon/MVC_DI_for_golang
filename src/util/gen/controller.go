package gen

import (
	"MVC_DI/global"
	"MVC_DI/util"
	"path"
)

// GenerateGinController 生成基于 Gin 的 Controller
func GenerateGinController(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateGinController(pkg, basePath, entity, table)
	}
}

func _generateGinController(pkg, basePath, entity, table string) {
	controllerResourcePath := global.PATH.RESOURCE.TEMPLATE.CONTROLLER.DIR
	controllerDir := append([]string{basePath, entity}, global.PATH.CONTROLLER.DIR...)
	util.CreateDir(path.Join(controllerDir...))

	coreTemplatePath := append(controllerResourcePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.CORE...)
	corePath := append(controllerDir, global.PATH.CONTROLLER.CORE...)

	GenerateTemplate(pkg, path.Join(coreTemplatePath...), path.Join(corePath...), entity, table)

	builderTemplatePath := append(controllerResourcePath, global.PATH.RESOURCE.TEMPLATE.CONTROLLER.BUILDER...)
	builderPath := append(controllerDir, global.PATH.CONTROLLER.BUILDER...)

	GenerateTemplate(pkg, path.Join(builderTemplatePath...), path.Join(builderPath...), entity, table)

}
