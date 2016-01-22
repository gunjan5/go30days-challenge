package main

import (
	"fmt"
	"github.com/Shopify/sarama" //GIT hash = b3d9702dd2d2cfe6b85b9d11d6f25689e6ef24b0
	"time"
)

//var groupName = "trash"
var topicName = "topica"
var partition int32 = 0
//var client *sarama.Client
//var err error

func clientStart() {
	fmt.Println("creating client")
	var err error
	client, err := sarama.NewClient([]string{"localhost:9092"}, nil)
	fmt.Printf("%T is the type of client", client)
	if err != nil {
		fmt.Println("ERROR:", err)
	}
}

// func consumer() {
// 	fmt.Println("creating consumer")
// 	consumer, err := sarama.NewConsumer(client, topicName, partition, nil, nil)
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}

// 	for e := range consumer.Events() {
// 		fmt.Println(string(e.Value))
// 	}
// }

// func producer() {
// 	fmt.Println("creating producer")
// 	producer, err := sarama.NewProducer(client, nil)
// 	if err != nil {
// 		fmt.Println("ERROR:", err)
// 	}
// 	for i := 0; i < 10; i++ {
// 		producer.SendMessage(topicName, nil, sarama.StringEncoder(fmt.Sprintf("A%d", i)))
// 	}
// }

func main() {

	clientStart()
    //consumer()

	time.Sleep(time.Second)

	//go producer()

	<-make(chan int)
}