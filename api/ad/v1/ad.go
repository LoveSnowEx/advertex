package v1

import (
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

type Gender int
type Country string
type Platform int

const (
	Male Gender = iota
	Female
)

const (
	Android Platform = iota
	Ios
	Web
)

type Conditons struct {
	Age      [][]int  `json:"age" dc:"年齡" example:"[[18], [7, 14], [65, 90, 5]]"`
	Gender   Gender   `json:"gender" dc:"性別：0:男; 1:女"`
	Country  []string `json:"country" dc:"國家" example:"['TW', 'JP']"`
	Platform Platform `json:"platform" dc:"平台：0:android; 1:ios; 2:web"`
}

type GetAdReq struct {
	g.Meta `path:"/ad" tags:"Ad" method:"get" sm:"Get ad list"`
}

type GetAdRes struct {
	g.Meta
	Data string
}

type CreateAdReq struct {
	g.Meta    `path:"/ad" tags:"Ad" method:"post" sm:"Create ad"`
	Title     string    `json:"title" v:"required#請輸入廣告標題" dc:"廣告標題"`
	StartAt   time.Time `json:"start_at" v:"required#請輸入開始時間" dc:"開始時間" example:"2024-01-01T03:00:00.000Z"`
	EndAt     time.Time `json:"end_at" v:"required#請輸入結束時間" dc:"結束時間" example:"2024-12-31T16:00:00.000Z"`
	Conditons Conditons `json:"conditons" dc:"篩選條件"`
}

type CreateAdRes struct {
	g.Meta
}
