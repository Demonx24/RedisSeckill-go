package database

import (
	"gorm.io/gorm"
	"time"
)

type MODEL struct {
	ID        uint           `json:"id" gorm:"primary_key"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" grom:"index"`
}

// 商品表
type Product struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string    `gorm:"type:varchar(255);not null" json:"name"`
	Description string    `gorm:"type:text" json:"description"`
	Price       float64   `gorm:"type:decimal(10,2);not null" json:"price"`
	StockCount  int       `gorm:"not null;default:0" json:"stock_count"`
	StartTime   time.Time `gorm:"not null" json:"start_time"`
	EndTime     time.Time `gorm:"not null" json:"end_time"`
	CreatedAt   time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt   time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 用户表
type User struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	Username  string    `gorm:"type:varchar(100);unique;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Email     string    `gorm:"type:varchar(255)" json:"email"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 订单表
type Order struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	OrderNo   string    `gorm:"type:varchar(64);unique;not null" json:"order_no"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	ProductID uint64    `gorm:"not null" json:"product_id"`
	Status    uint8     `gorm:"not null;default:0" json:"status"` // 0 未支付，1 已支付
	Amount    float64   `gorm:"type:decimal(10,2);not null" json:"amount"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}

// 秒杀成功用户记录表
type SeckillOrder struct {
	ID        uint64    `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    uint64    `gorm:"not null" json:"user_id"`
	ProductID uint64    `gorm:"not null" json:"product_id"`
	OrderID   uint64    `gorm:"not null" json:"order_id"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at"`
}
