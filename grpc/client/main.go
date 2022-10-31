package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"time"

	pb "github.com/WeixinX/examples/grpc/proto"

	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	// 使用 grpc.Dial() 与 gRPC 建立连接
	conn, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// 使用 pb.NewXXX(conn) 获取客户端
	c := pb.NewTimeServerClient(conn)

	// 请求接口测试
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	rawTime, err := c.GetRawTime(ctx, &pb.GetRawTimeReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("timestamp:", rawTime.Timestamp)

	fmtTime, err := c.GetFmtTime(ctx, &pb.GetFmtTimeReq{})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("format:", fmtTime.FmtTime)

	resp, err := c.GetFmtTimeFromRaw(ctx, &pb.GetFmtTimeFromRawReq{Timestamp: rawTime.Timestamp})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("format from timestamp", rawTime.Timestamp, ":", resp.FmtTime)
}
