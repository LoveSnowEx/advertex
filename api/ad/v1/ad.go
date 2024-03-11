package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAdReq struct {
	g.Meta `path:"/ad" tags:"Ad" method:"get" sm:"Get ad list"`
}

type GetAdRes struct {
	g.Meta
}

type CreateAdReq struct {
	g.Meta `path:"/ad" tags:"Ad" method:"post" sm:"Create ad"`
}

type CreateAdRes struct {
	g.Meta
}
