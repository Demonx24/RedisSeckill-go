package main

import (
	"RedisSeckill-go/core"
	"RedisSeckill-go/global"
	"RedisSeckill-go/initialize"
)

func main() {
	global.Config = core.InitConf()
	global.Redis = initialize.InitRedis()
	global.DB = initialize.InitGorm()
}
