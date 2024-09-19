package database

import (
	"gomailer/internal/domain/campaign"

	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func NewDatabase() *gorm.DB {
	options := map[string]string{
		"CONNECTION TIMEOUT": "90",
		"LANGUAGE":           "SIMPLIFIED CHINESE",
		"TERRITORY":          "CHINA",
		"SSL":                "false",
	}
	url := oracle.BuildUrl("localhost", 1521, "xe", "demo", "12345", options)
	dialector := oracle.New(oracle.Config{
		DSN:                     url,
		IgnoreCase:              false, // query conditions are not case-sensitive
		NamingCaseSensitive:     true,  // whether naming is case-sensitive
		VarcharSizeIsCharLength: true,  // whether VARCHAR type size is character length, defaulting to byte length

		// RowNumberAliasForOracle11 is the alias for ROW_NUMBER() in Oracle 11g, defaulting to ROW_NUM
		RowNumberAliasForOracle11: "ROW_NUM",
	})
	db, err := gorm.Open(dialector, &gorm.Config{
		SkipDefaultTransaction:                   true, // 是否禁用默认在事务中执行单次创建、更新、删除操作
		DisableForeignKeyConstraintWhenMigrating: true, // 是否禁止在自动迁移或创建表时自动创建外键约束
		NamingStrategy: schema.NamingStrategy{
			NoLowerCase:         true, // 是否不自动转换小写表名
			IdentifierMaxLength: 30,   // Oracle: 30, PostgreSQL:63, MySQL: 64, SQL Server、SQLite、DM: 128
		},
		PrepareStmt:     false, // 创建并缓存预编译语句，启用后可能会报 ORA-01002 错误
		CreateBatchSize: 50,    // 插入数据默认批处理大小
	})

	if err != nil {
		panic("fail to connect to database")
	}

	db.AutoMigrate(&campaign.Campaign{}, &campaign.Contact{})

	return db
}
