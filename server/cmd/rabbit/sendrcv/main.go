package main

import (
	"fmt"
	"time"

	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://myuser:mypass@localhost:5672/")
	if err != nil {
		panic(err)
	}

	// 虚拟连接
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}

	// 创建队列
	q, err := ch.QueueDeclare(
		"go_q1", true, false, false, false, nil,
	)
	if err != nil {
		panic(err)
	}

	go consume("c1", conn, q.Name)
	go consume("c2", conn, q.Name)

	i := 0
	for {
		i++
		// 发布消息到队列
		err := ch.Publish("", q.Name, false, false, amqp.Publishing{
			Body: []byte(fmt.Sprintf("message %d", i)),
		})
		if err != nil {
			panic(err)
		}
		time.Sleep(200 * time.Millisecond)
	}
}

// 消费方法
func consume(name string, conn *amqp.Connection, q string) {
	// 虚拟连接
	ch, err := conn.Channel()
	if err != nil {
		panic(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		q,
		"c1",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		panic(err)
	}

	for msg := range msgs {
		fmt.Printf("%s\n", msg.Body)
	}
}
