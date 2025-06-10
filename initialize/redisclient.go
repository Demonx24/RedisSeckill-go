package initialize

import (
	"RedisSeckill-go/global"
	"context"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var RedisClient *redis.Client

func InitRedis() redis.Client {
	RedisClient = redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Address,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	_, err := RedisClient.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}
	return *RedisClient
}
