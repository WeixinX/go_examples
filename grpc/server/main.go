package main

import (
	"context"
	"fmt"
	"net"
	"time"

	pb "github.com/WeixinX/examples/grpc/proto"

	"google.golang.org/grpc"
)

// timeServer 定义一个结构体用于实现 .proto文件中定义的方法
// 新版本 gRPC 要求必须嵌入 pb.UnimplementedGreeterServer 结构体
type timeServer struct {
	pb.UnimplementedTimeServerServer
}

// 实现 .proto 文件中定义的接口

func (t *timeServer) GetRawTime(ctx context.Context, req *pb.GetRawTimeReq) (*pb.GetRawTimeResp, error) {
	return &pb.GetRawTimeResp{Timestamp: time.Now().Unix()}, nil
}

func (t *timeServer) GetFmtTime(ctx context.Context, req *pb.GetFmtTimeReq) (*pb.GetFmtTimeResp, error) {
	return &pb.GetFmtTimeResp{FmtTime: time.Now().Format("2006-01-02 15:04:05")}, nil
}

func (t *timeServer) GetFmtTimeFromRaw(ctx context.Context, req *pb.GetFmtTimeFromRawReq) (*pb.GetFmtTimeFromRawResp, error) {
	timestamp := req.Timestamp
	return &pb.GetFmtTimeFromRawResp{FmtTime: time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")}, nil
}

func main() {
	// 绑定监听端口
	lsn, err := net.Listen("tcp", ":8002")
	if err != nil {
		panic(err)
	}
	svr := grpc.NewServer()                         // 创建 gRPC server
	pb.RegisterTimeServerServer(svr, &timeServer{}) // 把 timeServer 注册到 gRPC 中
	fmt.Println("timeserver running on :8002")
	// 启动服务
	if err = svr.Serve(lsn); err != nil {
		panic(err)
	}
}
