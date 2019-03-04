package GoWheels

import (
	"github.com/nsqio/go-nsq"
	"log"
	"fmt"
)

type PrintHandler struct{}

func (self *PrintHandler) HandleMessage(message *nsq.Message) error {
	fmt.Printf("message is %v \n", string(message.Body))
	return nil
}

func NsqPubSub() {

	topic := "nsq_pub_sub_test"

	//启动producer
	go func() {
		config := nsq.NewConfig()
		producer, e := nsq.NewProducer("localhost:4150", config)

		if e != nil {
			log.Fatal(3)
		}

		e = producer.Ping()
		if e != nil {
			producer.Stop()
			log.Fatal(e)
		}

		defer producer.Stop()

		for i := 0; i < 10; i++ {
			err := producer.Publish(topic, []byte(fmt.Sprintf("message-%d", i)))
			if err != nil {
				log.Println("publish error", err)
			}
		}

		select{}

	}()

	//consumer
	go func() {
		config := nsq.NewConfig()
		consumer, e := nsq.NewConsumer(topic, "nsq_pub_sub_default", config)
		if e != nil {
			log.Fatal(e)
		}

		consumer.AddHandler(&PrintHandler{})
		consumer.ConnectToNSQD("localhost:4150")
		consumer.ConnectToNSQLookupd("http://localhost:4161")
		select{}

		//clean stop
		consumer.Stop()
		<-consumer.StopChan
	}()

	stopChan := make(chan int)
	<-stopChan

}
