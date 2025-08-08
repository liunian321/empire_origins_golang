package main

import (
	"empire_origins_golang/config"
	"empire_origins_golang/db"

	"gorm.io/gen"
)

/*
用于生成 model 文件
*/
func main() {
	config.LoadEnv()
	// 创建生成器
	g := gen.NewGenerator(gen.Config{
		OutPath: "./model",           // 生成的 model 文件输出路径
		Mode:    gen.WithDefaultQuery, // 生成默认查询方法
	})

	// 设置数据库连接
	db.Init_Postgres()
	g.UseDB(db.DB)

	// 生成所有表的 model
	g.GenerateAllTable()

	// 执行生成
	g.Execute()
}
