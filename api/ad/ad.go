// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package ad

import (
	"context"

	"advertex/api/ad/v1"
)

type IAdV1 interface {
	GetAd(ctx context.Context, req *v1.GetAdReq) (res *v1.GetAdRes, err error)
	CreateAd(ctx context.Context, req *v1.CreateAdReq) (res *v1.CreateAdRes, err error)
}
