package api

import (
	"RedisSeckill-go/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

func SeckillHandler(c *gin.Context) {
	userID := c.Query("user_id")
	productID := c.Query("product_id")

	result, err := service.DoSeckill(userID, productID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "抢购失败", "error": err.Error()})
		return
	}

	switch result {
	case 0:
		c.JSON(http.StatusOK, gin.H{"message": "商品已售罄"})
	case 1:
		c.JSON(http.StatusOK, gin.H{"message": "抢购成功"})
	case 2:
		c.JSON(http.StatusOK, gin.H{"message": "不能重复抢购"})
	default:
		c.JSON(http.StatusOK, gin.H{"message": "未知结果"})
	}
}
