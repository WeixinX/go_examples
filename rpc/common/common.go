package common

// 获取时间戳

type GetRawTimeReq struct {
}

type GetRawTimeRes struct {
	Timestamp int64
}

// 获取格式化时间

type GetFmtTimeReq struct {
}

type GetFmtTimeRes struct {
	FmtTime string
}

// 通过时间戳获取格式化时间

type GetFmtTimeFromRawReq struct {
	Timestamp int64
}

type GetFmtTimeFromRawRes struct {
	FmtTime string
}
