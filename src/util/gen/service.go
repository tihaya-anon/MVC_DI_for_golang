package gen

import (
	"MVC_DI/global"
	"MVC_DI/util"
	"path"
)

// GenerateService 生成 Service 和 ServiceImpl
func GenerateService(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateService(pkg, basePath, entity, table)
	}
}

func _generateService(pkg, basePath, entity, table string) {
	serviceTemplatePath := global.PATH.RESOURCE.TEMPLATE.SERVICE.DIR
	serviceDir := append([]string{basePath, entity}, global.PATH.SERVICE.DIR...)
	util.CreateDir(path.Join(serviceDir...))

	interfaceTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.INTERFACE...)
	interfacePath := append(serviceDir, global.PATH.SERVICE.INTERFACE...)

	GenerateTemplate(pkg, path.Join(interfaceTemplatePath...), path.Join(interfacePath...), entity, table)

	builderTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.BUILDER...)
	builderPath := append(serviceDir, global.PATH.SERVICE.BUILDER...)

	GenerateTemplate(pkg, path.Join(builderTemplatePath...), path.Join(builderPath...), entity, table)

	implTemplatePath := append(serviceTemplatePath, global.PATH.RESOURCE.TEMPLATE.SERVICE.IMPL...)
	implPath := append(serviceDir, global.PATH.SERVICE.IMPL...)

	GenerateTemplate(pkg, path.Join(implTemplatePath...), path.Join(implPath...), entity, table)
}
