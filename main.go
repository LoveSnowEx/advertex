package main

import (
	_ "advertex/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"advertex/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
