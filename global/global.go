package global

import (
	"RedisSeckill-go/config"
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
)

var (
	Config *config.Config
	Redis  redis.Client
	DB     *gorm.DB
)
