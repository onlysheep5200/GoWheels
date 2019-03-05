package GoWheels

import (
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"fmt"
)

var topic = "test-topic"

//configMap的配置值 ： https://github.com/edenhill/librdkafka/blob/master/CONFIGURATION.md


func pub(doneChan chan interface{}) {


	producer, e := kafka.NewProducer(&kafka.ConfigMap{"bootstrap.servers": "localhost"})
	if e != nil {
		log.Fatal(e)
	}
	defer producer.Close()

	go func() {
		for ev := range producer.Events(){
			switch eve :=  ev.(type){
			case *kafka.Message :
				if eve.TopicPartition.Error != nil{
					fmt.Printf("deliver error %v\n", eve.TopicPartition)
				}
			}
		}

		close(doneChan)
	}()

	for i := 0; i < 10; i++ {
		message := fmt.Sprintf("message-%d", i)
		producer.Produce(&kafka.Message{
			TopicPartition:kafka.TopicPartition{Topic:&topic, Partition:kafka.PartitionAny},
			Value:[]byte(message),
		}, nil)
	}
}

func sub(doneChan chan interface{}) {
	consumer, e := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers" : "localhost",
		"group.id" : "myGroup",
		"auto.offset.reset" : "earliest",
	})

	if e != nil {
		log.Fatal(e)
	}

	consumer.Subscribe(topic, nil)

	for{
		message, e := consumer.ReadMessage(-1)
		if e != nil{
			log.Printf("consume message error %v", e)
		}else{
			log.Println(string(message.Value))
		}
	}

	close(doneChan)
}



