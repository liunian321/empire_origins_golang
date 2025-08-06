package db

import (
	"context"
	"log"
	"os"
	"strconv"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

func Init_Redis() {
	// 1. 获取环境变量并创建 redis 客户端
	db, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDR"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	// 2. 启动项目时检查连接是否正常
	if err := redisClient.Ping(context.Background()).Err(); err != nil {
		// 如果连接失败, 则退出程序
		log.Fatal("cannot connect to redis: %w", err)
	}
	// 3. 将 redis 客户端赋值给全局变量
	RedisClient = redisClient
}
