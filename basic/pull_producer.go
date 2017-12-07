//测试pull模式的生产者
package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
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
	q, err := ch.QueueDeclare(
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
		q.Name,         //queue name
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

	// //定义消息内容
	// body := [6]string{"aaa", "bbb", "ccc", "ddd", "xxx", "yyy"}
	// for _, b := range body {
	// 	//创建消息生产者
	// 	err = ch.Publish(
	// 		"goods_driect", //exchange名称
	// 		"html",         //routing_key 路由关键字
	// 		true,           //mandatory
	// 		false,          //immediate
	// 		amqp.Publishing{
	// 			DeliveryMode: amqp.Persistent, //将消息标记为持久的
	// 			ContentType:  "text/plain",
	// 			Body:         []byte(b),
	// 		})
	// 	failOnError(err, "Failed to publish a message")

	// 	log.Printf("[x] Sent %s", b)
	// }

	// //消息内容 由用户输入
	// body := severityFrom(os.Args)
	// err = ch.Publish(
	// 	"goods_driect", //exchange名称
	// 	"html",         //routing_key 路由关键字
	// 	true,           //mandatory
	// 	false,          //immediate
	// 	amqp.Publishing{
	// 		DeliveryMode: amqp.Persistent, //将消息标记为持久的
	// 		ContentType:  "text/plain",
	// 		Body:         []byte(body),
	// 	})
	// failOnError(err, "Failed to publish a message")

	// log.Printf("[x] Sent %s", body)

}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 3) || os.Args[2] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[2:], " ")
	}
	return s
}

func severityFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = os.Args[1]
	}
	return s
}
