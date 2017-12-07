### Golang RabbitMQ Demo


#### install:

```
go get github.com/streadway/amqp
```

#### 

##### 连接到RabbitMQ服务器

```
	// 连接到RabbitMQ服务器
	conn, err := amqp.Dial("amqp://user:123456@192.168.5.107:5672/")
	failOnError(err, "连接 RabbitMQ 失败")
	defer conn.Close()
```

##### 打开通道

```
	ch, err := conn.Channel()
	failOnError(err, "打开通道失败")
	defer ch.Close()
```


##### 声明队列，存储消息

```
	// 声明队列，存储消息
	q, err := ch.QueueDeclare(
		"task_queue",
		true,
		false,
		false,
		false,
		nil,
	)
```

##### 发布消息

```
	err = ch.Publish(
		"",
		q.Name,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(body),
		})
```

##### 接收消息

```
	msgs, err := ch.Consume(
		q.Name,
		"",
		false, //auto-ack
		false,
		false,
		false,
		nil,
	)
```

##### Bind

```
err = ch.QueueBind（
  q.Name，     //队列名称
  “black”，    //路由键
  “logs”，     //交换名称
  false，
   nil）
```

#### basic list

##### HelloWorld

* one_send.go
* one_receive.go

##### Work queues

* two_task.go
* two_worker.go

##### Publish/Subscribe Push Mode

* push_producer.go
* push_consumer.go

##### Publish/Subscribe Pull Mode

* pull_producer.go
* pull_consumer.go

![producer](basic/pull_producer.jpg)  ![consumer_1](basic/pull_consumer_1.jpg)  ![consumer_2](basic/pull_consumer_2.jpg)

#### 参考

* [RabbitMQ - RabbitMQ tutorial](https://www.rabbitmq.com/tutorials/tutorial-one-go.html)
