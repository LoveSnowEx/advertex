package ad

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"advertex/api/ad/v1"
)

func (c *ControllerV1) CreateAd(ctx context.Context, req *v1.CreateAdReq) (res *v1.CreateAdRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
