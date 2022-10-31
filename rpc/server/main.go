package main

import (
	"fmt"
	"net/http"
	"net/rpc"
	"time"

	"github.com/WeixinX/examples/rpc/common"
)

type TimeServer struct {
}

// 定义接口，签名需满足 func (t *T) MethodName(argType T1, replyType *T2) error

func (t *TimeServer) GetRawTime(req *common.GetRawTimeReq, res *common.GetRawTimeRes) error {
	res.Timestamp = time.Now().Unix()
	return nil
}

func (t *TimeServer) GetFmtTime(req *common.GetFmtTimeReq, res *common.GetFmtTimeRes) error {
	res.FmtTime = time.Now().Format("2006-01-02 15:04:05")
	return nil
}

func (t *TimeServer) GetFmtTimeFromRaw(req *common.GetFmtTimeFromRawReq, res *common.GetFmtTimeFromRawRes) error {
	timestamp := req.Timestamp
	res.FmtTime = time.Unix(timestamp, 0).Format("2006-01-02 15:04:05")
	return nil
}

func main() {
	server := new(TimeServer)
	rpc.Register(server) // 注册接口方法
	rpc.HandleHTTP()     // 注册 HTTP 路由
	fmt.Println("timeserver running...")
	if err := http.ListenAndServe(":8010", nil); err != nil { // 启动 HTTP 服务
		panic(err)
	}
}
