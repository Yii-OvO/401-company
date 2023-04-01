package main

import (
	_ "github.com/SupenBysz/gf-admin-community"

	"401-company/internal/boot"

	_ "github.com/SupenBysz/gf-admin-company-modules"

	_ "401-company/internal/logic"
	_ "github.com/gogf/gf/contrib/drivers/pgsql/v2"

	"github.com/gogf/gf/v2/os/gctx"
)

func main() {
	boot.Main.Run(gctx.New())
}
