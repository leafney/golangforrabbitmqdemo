//测试pull模式的消费者
package main

import (
	"github.com/streadway/amqp"
	"log"
	// "os"
	"time"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s：%s", msg, err)
	}
}

func main() {
	// 创建连接
	conn, err := amqp.Dial("amqp://user:123456@192.168.5.107:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	// 创建一个通道
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	// 创建Exchange交换器
	err = ch.ExchangeDeclare(
		"goods_driect", //name 名称
		"direct",       //kind 方式
		true,           //durable 持久的
		false,          //autoDelete自动删除
		false,          //internal 内部的
		false,          //noWait不等待
		nil,            //args参数
	)
	failOnError(err, "Failed to declare an exchange")

	// 创建一个名为pageCont的消息队列，并和exchange绑定
	//创建消息队列
	q1, err := ch.QueueDeclare(
		"pageCont", //name
		true,       //durable 持久队列
		false,      //autoDelete
		false,      //exclusive
		false,      //noWait
		nil,        //args
	)
	failOnError(err, "Failed to declare a queue")

	//RouteKey绑定exchange和queue
	err = ch.QueueBind(
		q1.Name,        //queue name
		"html",         // route_key
		"goods_driect", //exchange
		false,          // noWait
		nil)
	failOnError(err, "Failed to bind a queue")

	// // 创建一个名为getOnePrice的消息队列，并和exchange绑定
	// 	//创建消息队列
	// 	q, err := ch.QueueDeclare(
	// 		"getOnePrice", //name
	// 		true,          //durable 持久队列
	// 		false,         //autoDelete
	// 		false,         //exclusive
	// 		false,         //noWait
	// 		nil,           //args
	// 	)
	// 	failOnError(err, "Failed to declare a queue")

	// 	//RouteKey绑定exchange和queue
	// 	err = ch.QueueBind(
	// 		"getOnePrice",  //queue name
	// 		"price",        // route_key
	// 		"goods_driect", //exchange
	// 		false,          // noWait
	// 		nil)
	// 	failOnError(err, "Failed to bind a queue")

	//实现消息Pull模式
	forever := make(chan bool)
	go func() {
		for {
			msg, ok, err := ch.Get(q1.Name, true)
			failOnError(err, "Failed to register a consumer")
			if ok {
				log.Printf("[x] 队列中有任务，任务内容：%s", msg.Body)
			} else {
				log.Printf("[x] 没有任务")
			}

			time.Sleep(10 * time.Second)
		}
	}()

	log.Printf(" [*] Waiting for logs. To exit press CTRL+C")
	<-forever
}
