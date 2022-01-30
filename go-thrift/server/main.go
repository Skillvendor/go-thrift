package server

import (
	"fmt"

	"github.com/apache/thrift/lib/go/thrift"
)

func main() {

	var protocolFactory thrift.TProtocolFactory
	protocolFactory = thrift.NewTJSONProtocolFactory()

	var transportFactory thrift.TTransportFactory
	transportFactory = thrift.NewTFramedTransportFactory(transportFactory)

	// always run server here
	if err := runServer(transportFactory, protocolFactory, "localhost:1337", false); err != nil {
		fmt.Println("error running server:", err)
	}
}
