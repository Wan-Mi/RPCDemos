package main

import (
	"fmt"
	"github.com/Wan-Mi/RPCDemos/thriftDemo/hello"
	"net"
	"os"

	"git.apache.org/thrift.git/lib/go/thrift"

)

func main()  {

	transportFactory := thrift.NewTBufferedTransportFactory(8192)
	protocolFactory := thrift.NewTCompactProtocolFactory()

	transport, err := thrift.NewTSocket(net.JoinHostPort("127.0.0.1", "8988"))
	if err != nil {
		fmt.Fprintln(os.Stderr, "error resolving address:", err)
		os.Exit(1)
	}

	useTransport := transportFactory.GetTransport(transport)
	client := hello.NewHelloServiceClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Fprintln(os.Stderr, "Error opening socket to 127.0.0.1:8988", " ", err)
		os.Exit(1)
	}
	defer transport.Close()

	res, err := client.SayHello("jack",2)
	if err != nil {
		fmt.Println("Hello failed:", err)
		return
	}

	fmt.Println("response: ", res)
	fmt.Println("well done")

}