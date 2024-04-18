package main

import (
	_ "advertex/internal/packed"
	"context"
	"os"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite"
	_ "github.com/golang-migrate/migrate/v4/source/file"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/golang-migrate/migrate/v4"

	"advertex/internal/cmd"
)

func main() {
	sourcePath := g.Cfg().MustGet(context.Background(), "migration.sourcePath").String()
	databaseUrl := g.Cfg().MustGet(context.Background(), "migration.databaseUrl").String()
	g.Log().Info(context.Background(), "running migration")
	m, err := migrate.New(sourcePath, databaseUrl)
	if err != nil {
		g.Log().Error(context.Background(), "error creating migration", err)
		os.Exit(1)
	}
	if err := m.Up(); err != nil {
		if err.Error() == "no change" {
			g.Log().Info(context.Background(), "no migration needed")
		} else {
			g.Log().Error(context.Background(), "error running migration", err)
			os.Exit(1)
		}
	}
	cmd.Main.Run(gctx.GetInitCtx())
}
