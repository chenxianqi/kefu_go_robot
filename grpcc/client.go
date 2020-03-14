package grpcc

import (
	"fmt"
	"kefu_server/grpcs"
	"sync"
	"sync/atomic"
	"unsafe"

	"google.golang.org/grpc"
)

var (
	globalClientConn unsafe.Pointer
	lck              sync.Mutex
)

// GrpcClient get grpc cline instance
func GrpcClient() grpcs.KefuClient {
	conn, err := initConn()
	if err != nil {
		fmt.Print("grpcClient grpc CONNECT err:" + err.Error())
		return (grpcs.NewKefuClient)(nil)
	}
	return grpcs.NewKefuClient(conn)
}

// initConn get connect
func initConn() (*grpc.ClientConn, error) {
	if atomic.LoadPointer(&globalClientConn) != nil {
		return (*grpc.ClientConn)(globalClientConn), nil
	}
	lck.Lock()
	defer lck.Unlock()
	if atomic.LoadPointer(&globalClientConn) != nil {
		return (*grpc.ClientConn)(globalClientConn), nil
	}
	cli, err := newGrpcConn()
	if err != nil {
		return nil, err
	}
	atomic.StorePointer(&globalClientConn, unsafe.Pointer(cli))
	return cli, nil
}

// newGrpcConn
func newGrpcConn() (*grpc.ClientConn, error) {
	conn, err := grpc.Dial(
		"127.0.0.1:8028",
		grpc.WithInsecure(),
	)
	if err != nil {
		return nil, err
	}
	fmt.Printf("grpcClient grpc success")
	return conn, nil
}

// Run grpc client
func Run() {
	initConn()
}
