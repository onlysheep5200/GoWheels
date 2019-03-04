package GoWheels

import (
	"github.com/apache/thrift/lib/go/thrift"
	"log"
	"./thrift_demo"
	"context"
)


func RunGreeterClient() {
	tsock, err := thrift.NewTSocket("localhost:9999")
	if err != nil {
		log.Fatal(err)
	}

	/**
	部分注意事项：
	1.server与client的transport factory需相同
	2.server与client的protocol需相同
	 */
	ttransportFactory := thrift.NewTTransportFactory()
	ttransportFactory = thrift.NewTFramedTransportFactory(ttransportFactory)
	ttransport, err := ttransportFactory.GetTransport(tsock)
	if err != nil {
		log.Fatal(err)
	}
	defer ttransport.Close()

	if err = ttransport.Open(); err != nil {
		log.Fatal(err)
	}

	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()
	iproto := protocolFactory.GetProtocol(ttransport)
	oproto := protocolFactory.GetProtocol(ttransport)

	client := thrift_demo.NewGreeterClientProtocol(ttransport, iproto, oproto)
	response, err := client.Greet(context.Background(), "hyd")
	if err != nil{
		log.Println("invoke error")
		log.Fatal(err)
	}

	log.Println(response)
}
