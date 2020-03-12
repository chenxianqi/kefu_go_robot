package grpcc

import (
	"context"
	"kefu_server/grpcs"
	"log"

	"google.golang.org/grpc"
)

// Run grpc client
func Run() {
	// 建立连接到gRPC服务
	conn, err := grpc.Dial("127.0.0.1:8028", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	// 函数结束时关闭连接
	defer conn.Close()

	// 创建Waiter服务的客户端
	t := grpcs.NewWaiterClient(conn)

	// 模拟请求数据
	res := "test123"

	// 调用gRPC接口
	tr, err := t.DoMD5(context.Background(), &grpcs.Req{JsonStr: res})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("服务端响应: %s", tr.BackJson)
}
