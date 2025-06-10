package config

import (
	"gorm.io/gorm/logger"
	"strconv"
	"strings"
)

type Redis struct {
	Address  string `json:"address" yaml:"address"`   // Redis 服务器的地址，通常为 "localhost:6379" 或其他主机和端口
	Password string `json:"password" yaml:"password"` // 连接 Redis 时的密码，如果没有设置密码则留空
	DB       int    `json:"db" yaml:"db"`             // 指定使用的数据库索引，单实例模式下可选择的数据库，默认为 0
}
type System struct {
	Host           string `json:"-" yaml:"host"`                          // 服务器绑定的主机地址，通常为 0.0.0.0 表示监听所有可用地址
	Port           int    `json:"-" yaml:"port"`                          // 服务器监听的端口号，通常用于 HTTP 服务
	Env            string `json:"-" yaml:"env"`                           // Gin 的环境类型，例如 "debug"、"release" 或 "test"
	RouterPrefix   string `json:"-" yaml:"router_prefix"`                 // API 路由前缀，用于构建 API 路径
	UseMultipoint  bool   `json:"use_multipoint" yaml:"use_multipoint"`   // 是否启用多点登录拦截，防止同一账户在多个地方同时登录
	SessionsSecret string `json:"sessions_secret" yaml:"sessions_secret"` // 用于加密会话的密钥，确保会话数据的安全性
	OssType        string `json:"oss_type" yaml:"oss_type"`               // 对应的对象存储服务类型，如 "local" 或 "qiniu"
}
type Mysql struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Config       string `yaml:"config" json:"config"`
	DBName       string `yaml:"db_name" json:"db_name"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns" json:"max_open_conns"`
	LogMode      string `yaml:"log_mode" json:"log_mode"`
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" +
		m.Host + ":" + strconv.Itoa(m.Port) + ")/" +
		m.DBName + "?" + m.Config
}

func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	}
	return logger.Info
}
