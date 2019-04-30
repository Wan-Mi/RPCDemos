/**
官方安装命令无法安装的问题

go get google.golang.org/grpc

会报：
package google.golang.org/grpc: unrecognized import path "google.golang.org/grpc"(https fetch: Get https://google.golang.org/grpc?go-get=1: dial tcp 216.239.37.1:443: i/o timeout)


原因是这个代码已经转移到github上面了，但是代码里面的包依赖还是没有修改，还是google.golang.org这种，所以不能使用go get的方式安装。

需要：

git clone https://github.com/grpc/grpc-go.git $GOPATH/src/google.golang.org/grpc
git clone https://github.com/golang/net.git $GOPATH/src/golang.org/x/net
git clone https://github.com/golang/text.git $GOPATH/src/golang.org/x/text
go get -u github.com/golang/protobuf/{proto,protoc-gen-go}
git clone https://github.com/google/go-genproto.git $GOPATH/src/google.golang.org/genproto

cd $GOPATH/src/
go install google.golang.org/grpc

 */

/**
	本例拷贝自官方文档：https://www.grpc.io/docs/quickstart/go/
 	生成指令：protoc -I helloworld/ helloworld/helloworld.proto --go_out=plugins=grpc:helloworld
 */



package grpcDemo
