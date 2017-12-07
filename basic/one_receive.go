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
	conn, err := amqp.Dial("amqp://user:123456@192.168.5.107:5672/")
	failOnError(err, "连接 RabbitMQ 失败")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "打开通道失败")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello",
		false,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "声明队列失败")

	msgs, err := ch.Consume(q.Name,
		"", true, false, false, false, nil)
	failOnError(err, "注册消费者失败")

	fovever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for messages. To exit press CTRL+C")
	<-fovever

}
