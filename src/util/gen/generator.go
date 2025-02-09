package gen

import (
	"MVC_DI/config"
	"MVC_DI/util"
	"log"
	"path"
	"strings"

	"github.com/google/uuid"

	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

type commonMethod struct {
	ID int64
}

func (commonMethod *commonMethod) BeforeCreate(tx *gorm.DB) error {
	commonMethod.ID = int64(uuid.New().ID())
	return nil
}

// Generate 生成代码
func Generate(pkg string, entities []string) {
	gormDB, err := gorm.Open(mysql.Open(config.Application.Database.Uri), &gorm.Config{})
	if err != nil {
		log.Fatalf("connect to database failed: %v", err)
	}
	log.Printf("connect to: %v\n", config.Application.Database.Uri)

	for _, entity := range entities {
		basePath := "./section/" + entity
		// 生成 Query
		generateQuery(basePath, entity, gormDB)

		// 获取当前实体的所有表
		tables := getEntityTables(gormDB, entity)

		// 生成 Service 和 ServiceImpl
		GenerateService(pkg, basePath, entity, tables)

		// 生成 Mapper 和 MapperImpl
		GenerateMapper(pkg, basePath, entity, tables)

		// 生成基于 Gin 的 Controller
		GenerateGinController(pkg, basePath, entity, tables)
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

// generateQuery 生成 Query
func generateQuery(basePath, entity string, gormDB *gorm.DB) {
	tmpPath := basePath + "/query"

	// 创建目录
	util.CreateDir(tmpPath)

	// 初始化生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: tmpPath,
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface,
	})
	g.UseDB(gormDB)

	// 生成 Query 代码
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
		gen.WithMethod(commonMethod{}))...)
	g.Execute()
	util.MoveDir(tmpPath, path.Join("database", entity))
}
