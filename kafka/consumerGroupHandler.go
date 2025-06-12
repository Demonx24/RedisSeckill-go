package kafka

import (
	"RedisSeckill-go/global"
	"encoding/json"
	"github.com/Shopify/sarama"
	"log"
)

type Order struct {
	ID        int    `json:"id"`
	ProductId string `json:"product_id"`
	UserId    string `json:"user_id"`
}
type MyConsumerGroupHandler struct{}

func (h MyConsumerGroupHandler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h MyConsumerGroupHandler) Cleanup(sarama.ConsumerGroupSession) error { return nil }

func (h MyConsumerGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// 这里写业务处理代码，比如解析订单，写数据库
		var order Order
		if err := json.Unmarshal(msg.Value, &order); err != nil {
			log.Printf("订单解析失败: %v", err)
			continue
		}

		// 幂等写入数据库
		if err := global.DB.Create(&order).Error; err != nil {
			log.Printf("写入数据库失败: %v", err)
			continue
		}

		// 标记该消息已经成功消费，准备提交offset
		sess.MarkMessage(msg, "")
	}
	return nil
}
