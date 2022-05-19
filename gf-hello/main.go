package main

import (
	_ "gf-hello/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"gf-hello/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
