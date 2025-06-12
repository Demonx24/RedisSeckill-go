// 📦 Go Kafka DEMO: 消息传递 + 异常重试
// 使用 sarama Go 客户端连接 Kafka，模拟生产者 + 消费者 + 异常重试

package kafka

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Shopify/sarama"
)

const (
	topic         = "demo-topic"
	brokerAddress = "localhost:9092"
	retryLimit    = 3
)

// 模拟消息体
type Message struct {
	ID      string
	Content string
}

// Producer 发送消息
func produceMessage(msg string) {
	producer, err := sarama.NewSyncProducer([]string{brokerAddress}, nil)
	if err != nil {
		log.Fatalf("Failed to start producer: %v", err)
	}
	defer producer.Close()

	message := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(msg),
	}

	partition, offset, err := producer.SendMessage(message)
	if err != nil {
		log.Printf("Send error: %v", err)
		return
	}
	fmt.Printf("Sent to partition %d, offset %d\n", partition, offset)
}

// Consumer 消费者 + 重试机制
func consumeWithRetry() {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	consumer, err := sarama.NewConsumer([]string{brokerAddress}, config)
	if err != nil {
		log.Fatalf("Failed to start consumer: %v", err)
	}
	defer consumer.Close()

	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalf("Failed to consume partition: %v", err)
	}
	defer partitionConsumer.Close()

	// 优雅退出信号
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("🚀 Start consuming...")

	// 消费主循环
ConsumerLoop:
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			success := handleMessageWithRetry(msg.Value)
			if !success {
				log.Printf("❌ 消息处理失败超过上限: %s", string(msg.Value))
			}
		case err := <-partitionConsumer.Errors():
			log.Printf("❗ Kafka error: %v", err)
		case <-sigchan:
			fmt.Println("👋 Exit signal received")
			break ConsumerLoop
		}
	}
}

// 模拟处理逻辑 + 重试机制
func handleMessageWithRetry(value []byte) bool {
	var attempt int
	for attempt = 1; attempt <= retryLimit; attempt++ {
		success := handleMessage(value)
		if success {
			return true
		}
		log.Printf("⚠️ 第 %d 次重试失败: %s", attempt, string(value))
		time.Sleep(1 * time.Second)
	}
	return false
}

// 模拟业务逻辑，失败概率 30%
func handleMessage(value []byte) bool {
	if time.Now().UnixNano()%10 < 1 {
		return false // 模拟失败
	}
	fmt.Printf("✅ 成功处理消息: %s\n", string(value))
	return true
}

func main() {

	//go consumeWithRetry()
	//time.Sleep(2 * time.Second)
	//produceMessage("Kafka 消息 - 订单支付成功")
	//produceMessage("Kafka 消息 - 用户注册")
	//select {} // 阻塞主协程
	//producer, _ := sarama.NewSyncProducer([]string{"localhost:9092"}, nil)
	//msg := &sarama.ProducerMessage{Topic: "demo-topic", Value: sarama.StringEncoder("hello")}
	//producer.SendMessage(msg)
	//
	//config := sarama.NewConfig()
	//config.Version = sarama.V2_8_0_0
	//config.Consumer.Return.Errors = true
	//config.Consumer.Offsets.Initial = sarama.OffsetNewest
	//
	//group, _ := sarama.NewConsumerGroup([]string{"localhost:9092"}, "group-1", config)
	//select {} // 阻塞主协程
	//defer group.Close()
	//fmt.Println("<UNK> <UNK>")
	config := sarama.NewConfig()
	config.Version = sarama.V2_1_0_0
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	group, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "group-1", config)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	handler := MyConsumerGroupHandler{}
	for {
		err := group.Consume(ctx, []string{"demo-topic"}, handler)
		if err != nil {
			log.Fatal(err)
		}
	}
}
