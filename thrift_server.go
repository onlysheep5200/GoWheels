package GoWheels

import (
	"./thrift_demo"
	"fmt"
	"context"
	"log"
	"github.com/apache/thrift/lib/go/thrift"
)

type GreeterHandler struct{}

func (self *GreeterHandler) Greet(ctx context.Context, name string) (*thrift_demo.Status, error){
	status := thrift_demo.Status{Code:0}
	extra := make(map[string]string)
	extra["msg"] = fmt.Sprintf("welcome %s", name)
	status.Extra = extra
	return &status, nil
}

func RunServer() {
	host := "127.0.0.1"
	port := 9999

	processor := thrift_demo.NewGreeterProcessor(&GreeterHandler{})
	socket, e := thrift.NewTServerSocket(fmt.Sprintf("%s:%d", host, port))
	if e != nil {
		log.Fatal(e)
	}

	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	tprotocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	server := thrift.NewTSimpleServer4(processor, socket, transportFactory, tprotocolFactory)
	server.Serve()
}


