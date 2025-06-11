package service

import (
	"RedisSeckill-go/global"
	"context"
	"os"
)

func DoSeckill(userID, productID string) (int, error) {
	script, err := os.ReadFile("script/lua/seckill.lua")
	if err != nil {
		return -1, err
	}

	keys := []string{
		"seckill:stock:" + productID,
		"seckill:users:" + productID,
	}

	result, err := global.Redis.Eval(context.Background(), string(script), keys, userID).Int()
	return result, err
}
