package pb

import (
	"ege_platform/internal/api/middleware"
	"ege_platform/internal/config"

	mw "github.com/labstack/echo/v5/middleware"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

type PB struct {
	*pocketbase.PocketBase
}

func NewPB(cfg *config.Config) *PB {
	return &PB{
		PocketBase: pocketbase.NewWithConfig(
			&pocketbase.Config{
				DefaultDebug: cfg.PBDebug,
			},
		),
	}
}

func (p *PB) SetupMiddlewares() {
	p.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.Use(middleware.LoggerMiddleware())
		e.Router.Use(middleware.DBSessionMiddleware(p.Dao()))
		e.Router.Use(mw.Recover())
		return nil
	})
}

func (p *PB) Run() error {
	p.SetupMiddlewares()
	return p.Start()
}
