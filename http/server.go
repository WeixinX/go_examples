package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Request 接收 body 中的 json 格式数据
type Request struct {
	Timestamp int64 `json:"timestamp"`
}

// getNowRawTime 获取当前时间戳
func getNowRawTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Unix()
	fmt.Fprintln(w, "timestamp: ", t)
}

// getNowFmtTime 获取当前格式化时间
func getNowFmtTime(w http.ResponseWriter, r *http.Request) {
	t := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintln(w, "format: ", t)
}

// getFmtTimeFromRaw 将时间戳转化为格式化时间
func getFmtTimeFromRaw(w http.ResponseWriter, r *http.Request) {
	// 从 request body 获取数据
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	// 反序列化 body 获得 json 格式参数
	req := new(Request)
	err = json.Unmarshal(body, req)
	if err != nil {
		fmt.Fprintln(w, err)
	}
	// 将时间戳格式化
	t := time.Unix(req.Timestamp, 0).Format("2006-01-02 15:04:05")
	fmt.Fprintln(w, "format: ", t)
}

func main() {
	http.HandleFunc("/api/now_raw_time", getNowRawTime)
	http.HandleFunc("/api/now_fmt_time", getNowFmtTime)
	http.HandleFunc("/api/fmt_time_from_raw", getFmtTimeFromRaw)
	fmt.Println("server running...")
	err := http.ListenAndServe(":8008", nil)
	if err != nil {
		panic(err)
	}
}
