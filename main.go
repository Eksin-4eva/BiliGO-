package main

import (
	"log"
	"os"
	"strconv"
	"time"

	"github.com/BiliGO/biz/dal/mysql"
	"github.com/joho/godotenv"
	"github.com/cloudwego/hertz/pkg/app/server"
)

func main() {
	// 加载 .env（生产环境可直接注入环境变量，不依赖此文件）
	_ = godotenv.Load()

	if err := mysql.Init(&mysql.Config{
		DSN:             mustEnv("MYSQL_DSN"),
		MaxOpenConns:    envInt("MYSQL_MAX_OPEN_CONNS", 100),
		MaxIdleConns:    envInt("MYSQL_MAX_IDLE_CONNS", 10),
		ConnMaxLifetime: time.Hour,
	}); err != nil {
		log.Fatalf("mysql init failed: %v", err)
	}

	h := server.Default()
	register(h)
	h.Spin()
}

func mustEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		log.Fatalf("env %s is required", key)
	}
	return v
}

func envInt(key string, defaultVal int) int {
	v := os.Getenv(key)
	if v == "" {
		return defaultVal
	}
	n, err := strconv.Atoi(v)
	if err != nil {
		return defaultVal
	}
	return n
}
