package api

import (
	"RedisSeckill-go/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var resultadd = [3]int{0, 0, 0}
var number = 0

func SeckillHandler(c *gin.Context) {
	number++
	userID := c.Query("user_id")
	productID := c.Query("product_id")

	result, err := service.DoSeckill(userID, productID)
	if result == 1 {
		orderMsg := fmt.Sprintf("{\"user_id\": \"%s\", \"product_id\": \"%s\"}", userID, productID)
		service.SendKafkaMessage(orderMsg)
	}
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "抢购失败", "error": err.Error()})
		return
	}

	switch result {
	case 0:
		c.JSON(http.StatusOK, gin.H{"message": "商品已售罄"})
		resultadd[0]++ //500
	case 1:
		c.JSON(http.StatusOK, gin.H{"message": "抢购成功"})
		resultadd[1]++
	case 2:
		c.JSON(http.StatusOK, gin.H{"message": "不能重复抢购"})
		resultadd[2]++ //403
	default:
		c.JSON(http.StatusOK, gin.H{"message": "未知结果"}) //400
	}
	fmt.Println(resultadd)
	fmt.Println(number)
}
