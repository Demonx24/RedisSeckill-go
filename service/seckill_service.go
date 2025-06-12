package service

import (
	"RedisSeckill-go/global"
	"context"
	"go.opentelemetry.io/otel"
	"os"
)

var tracer = otel.Tracer("seckill-service")

func DoSeckill(userID, productID string) (int, error) {
	_, span := tracer.Start(context.Background(), "DoSeckillLogic")
	defer span.End()
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
