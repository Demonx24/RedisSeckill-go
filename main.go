package main

import (
	"RedisSeckill-go/api"
	"RedisSeckill-go/core"
	"RedisSeckill-go/global"
	"RedisSeckill-go/initialize"
	"RedisSeckill-go/service"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"log"
	"runtime"
)

const number = 0

func main() {
	shutdown := initialize.InitTracer("redis-seckill-go")
	defer func() {
		if err := shutdown(context.Background()); err != nil {
			log.Fatalf("Error shutting down tracer: %v", err)
		}
	}()
	global.Config = core.InitConf()
	global.Redis = initialize.InitRedis()
	global.DB = initialize.InitGorm()
	go service.StartOrderConsumer()
	r := gin.New()
	runtime.GOMAXPROCS(runtime.NumCPU())
	r.Use(otelgin.Middleware("seckill-server"))
	r.GET("/seckill", api.SeckillHandler)
	r.GET("/seckill/users", func(c *gin.Context) {
		ctx := context.Background()
		userKey := "seckill:users"
		members, err := global.Redis.SMembers(ctx, userKey).Result()
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		fmt.Println("查询到的用户集合:", members) // 打印日志
		c.JSON(200, gin.H{"users": members})
	})
	r.Run(":8081")

}
