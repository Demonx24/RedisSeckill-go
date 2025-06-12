package service

import (
	"RedisSeckill-go/kafka"
	"context"
	"github.com/Shopify/sarama"
	"log"
)

func SendKafkaMessage(msg string) {
	producer, err := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	if err != nil {
		log.Fatal("Kafka 生产者创建失败:", err)
	}
	defer producer.Close()

	_, _, err = producer.SendMessage(&sarama.ProducerMessage{
		Topic: "seckill-order",
		Value: sarama.StringEncoder(msg),
	})
	if err != nil {
		log.Println("发送 Kafka 消息失败:", err)
	}
}
func StartOrderConsumer() {
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "group-1", config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	handler := kafka.MyConsumerGroupHandler{}
	for {
		err := group.Consume(ctx, []string{"seckill-order"}, handler)
		if err != nil {
			log.Fatal(err)
		}
	}
}
