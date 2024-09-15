// Order Generator
package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/IBM/sarama"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Kafka集群的地址
	brokers := []string{"127.0.0.1:10000", "127.0.0.1:10001", "127.0.0.1:10002"}

	// 创建生产者配置
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForLocal // 等待本地复制
	config.Producer.Return.Successes = true            // 成功交付的消息将在success channel返回

	// 连接Kafka集群
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		panic("Failed to start Sarama producer:" + err.Error())
	}
	defer producer.Close()

	// 模拟订单生成
	for i := 0; i < 10; i++ {
		orderID := fmt.Sprintf("Order%d", i)
		productID := rand.Intn(100)
		quantity := rand.Intn(5) + 1

		// 构造订单消息
		message := &sarama.ProducerMessage{
			Topic: "orders",
			//这里key是可选的。
			Value: sarama.StringEncoder(fmt.Sprintf("%s,%d,%d", orderID, productID, quantity)),
		}

		// 发送消息
		partition, offset, err := producer.SendMessage(message)
		if err != nil {
			fmt.Printf("Failed to send message: %s\n", err)
		} else {
			fmt.Printf("Message sent to topic(%s)/partition(%d)/offset(%d)\n", "orders", partition, offset)
		}

		time.Sleep(time.Second)
	}
}
