//go:build ignore

package main

import (
	"github.com/BiliGO/biz/dal/model"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库（运行前替换 DSN）
	db, err := gorm.Open(mysql.Open("root:password@tcp(127.0.0.1:3306)/biligo?charset=utf8mb4&parseTime=True&loc=Local"))
	if err != nil {
		panic(err)
	}

	g := gen.NewGenerator(gen.Config{
		OutPath:      "./biz/dal/query",
		Mode:         gen.WithDefaultQuery | gen.WithQueryInterface,
		FieldNullable: true,
	})

	g.UseDB(db)

	// 基于已有 model 生成 DAO
	g.ApplyBasic(
		model.User{},
		model.Video{},
		model.Comment{},
		model.Favorite{},
		model.Relation{},
	)

	g.Execute()
}
