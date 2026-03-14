//go:build ignore

package main

import (
	"log"
	"os"

	"github.com/BiliGO/biz/dal/model"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	_ = godotenv.Load()
	dsn := os.Getenv("MYSQL_DSN")
	if dsn == "" {
		log.Fatal("MYSQL_DSN is required")
	}
	db, err := gorm.Open(mysql.Open(dsn))
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
