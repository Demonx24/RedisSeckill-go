package main

import (
	"RedisSeckill-go/api"
	"RedisSeckill-go/core"
	"RedisSeckill-go/global"
	"RedisSeckill-go/initialize"
	"github.com/gin-gonic/gin"
)

func main() {
	global.Config = core.InitConf()
	global.Redis = initialize.InitRedis()
	global.DB = initialize.InitGorm()
	r := gin.Default()
	r.GET("/seckill", api.SeckillHandler)

	r.Run(":8080")
}
