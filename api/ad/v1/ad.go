package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type GetAdReq struct {
	g.Meta `path:"/ad" tags:"Ad" method:"get" summary:"Get ad list"`
}

type GetAdRes struct {
	g.Meta `mime:"application/json"`
}

type CreateAdReq struct {
	g.Meta `path:"/ad" tags:"Ad" method:"post" summary:"Create ad"`
}

type CreateAdRes struct {
	g.Meta `mime:"application/json"`
}
