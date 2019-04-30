package main

import (
	"fmt"
	"github.com/Wan-Mi/RPCDemos/thriftDemo/hello"
	"git.apache.org/thrift.git/lib/go/thrift"
)

type HelloServer struct {
}

func (e *HelloServer)SayHello(userName string, userAge int32) (r *hello.User, err error){
	usr := hello.User{
		Name:userName,
		Age:userAge,
	}

	fmt.Println(usr)

	return &usr,nil
}

func main()  {

	transport, err := thrift.NewTServerSocket(":8988")
	if err != nil {
		panic(err)
	}

	handler := &HelloServer{}
	processor := hello.NewHelloServiceProcessor(handler)

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()
	server := thrift.NewTSimpleServer4(
		processor,
		transport,
		transportFactory,
		protocolFactory,
	)

	fmt.Println("server start")

	if err := server.Serve(); err != nil {
		panic(err)
	}

}