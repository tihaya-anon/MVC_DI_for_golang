package gen

import (
	"MVC_DI/global"
	"MVC_DI/util"
	"path"
)

// GenerateMapper 生成 Mapper 和 MapperImpl
func GenerateMapper(pkg, basePath, entity string, tables []string) {
	for _, table := range tables {
		_generateMapper(pkg, basePath, entity, table)
	}
}

func _generateMapper(pkg, basePath, entity, table string) {
	mapperTemplatePath := global.PATH.RESOURCE.TEMPLATE.MAPPER.DIR
	mapperDir := append([]string{basePath, entity}, global.PATH.MAPPER.DIR...)
	util.CreateDir(path.Join(mapperDir...))

	interfaceTemplatePath := append(mapperTemplatePath, global.PATH.RESOURCE.TEMPLATE.MAPPER.INTERFACE...)
	interfacePath := append(mapperDir, global.PATH.MAPPER.INTERFACE...)

	GenerateTemplate(pkg, path.Join(interfaceTemplatePath...), path.Join(interfacePath...),"_mapper", entity, table)

	implTemplatePath := append(mapperTemplatePath, global.PATH.RESOURCE.TEMPLATE.MAPPER.IMPL...)
	implPath := append(mapperDir, global.PATH.MAPPER.IMPL...)

	GenerateTemplate(pkg, path.Join(implTemplatePath...), path.Join(implPath...),"_mapper_impl", entity, table)
}
