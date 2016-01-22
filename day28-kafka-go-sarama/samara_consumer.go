package main

import (
	"fmt"
	"os"
	"log"
	"os/signal"
	"github.com/Shopify/sarama"
)


func main() {
	consumer, err := sarama.NewConsumer([]string{"192.168.59.103:9092"}, nil)
if err != nil {
    panic(err)
}

defer func() {
    if err := consumer.Close(); err != nil {
        log.Fatalln(err)
    }
}()

partitionConsumer, err := consumer.ConsumePartition("topicb", 0, sarama.OffsetNewest)
if err != nil {
    panic(err)
}

defer func() {
    if err := partitionConsumer.Close(); err != nil {
        log.Fatalln(err)
    }
}()

// Trap SIGINT to trigger a shutdown.
signals := make(chan os.Signal, 1)
signal.Notify(signals, os.Interrupt)

consumed := 0
ConsumerLoop:
for {
    select {
    case msg := <-partitionConsumer.Messages():
        fmt.Printf("Consumed message offset %d\n", msg.Offset)
        fmt.Println("msg", msg.)
        consumed++
    case <-signals:
        break ConsumerLoop
    }
}

fmt.Printf("Consumed: %d\n", consumed)

}