package main

import (
	"fmt"
	"github.com/WeixinX/examples/rpc/common"
	"net/rpc"
)

type TimeClient struct {
	c *rpc.Client
}

// 以下方法均基于 rpc.Client 包装了一层

func DialTimeServer(addr string) (*TimeClient, error) {
	c, err := rpc.DialHTTP("tcp", addr) // 连接到服务器监听地址
	return &TimeClient{c: c}, err
}

func (c *TimeClient) Close() error {
	return c.c.Close()
}

func (c *TimeClient) GetRawTime(req *common.GetRawTimeReq, res *common.GetRawTimeRes) error {
	return c.c.Call("TimeServer.GetRawTime", req, res)
}

func (c *TimeClient) GetFmtTime(req *common.GetFmtTimeReq, res *common.GetFmtTimeRes) error {
	return c.c.Call("TimeServer.GetFmtTime", req, res)

}

func (c *TimeClient) GetFmtTimeFromRaw(req *common.GetFmtTimeFromRawReq, res *common.GetFmtTimeFromRawRes) error {
	return c.c.Call("TimeServer.GetFmtTimeFromRaw", req, res)

}

func main() {
	client, err := DialTimeServer(":8010")
	if err != nil {
		panic(err)
	}
	defer client.Close()

	// 接口测试
	rawTime := new(common.GetRawTimeRes)
	err = client.GetRawTime(&common.GetRawTimeReq{}, rawTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("timestamp: ", rawTime.Timestamp)

	fmtTime := new(common.GetFmtTimeRes)
	err = client.GetFmtTime(&common.GetFmtTimeReq{}, fmtTime)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("format: ", fmtTime.FmtTime)

	fmtTimeFromRawReq := new(common.GetFmtTimeFromRawReq)
	fmtTimeFromRawRes := new(common.GetFmtTimeFromRawRes)
	fmtTimeFromRawReq.Timestamp = rawTime.Timestamp
	err = client.GetFmtTimeFromRaw(fmtTimeFromRawReq, fmtTimeFromRawRes)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("format from timestamp", rawTime.Timestamp, ": ", fmtTimeFromRawRes.FmtTime)
}
