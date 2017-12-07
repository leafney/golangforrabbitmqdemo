package main

import (
	"github.com/streadway/amqp"
	"log"
)

// 处理错误
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://user:123456@192.168.5.107:5672/")
	failOnError(err, "连接 RabbitMQ 失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "打开通道失败")
	defer ch.Close()

	// 声明队列，存储消息
	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "声明队列失败")

	body := "我是刘备"

	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})

	log.Printf("[x] Sent %s", body)
	failOnError(err, "发送消息失败")
}
