package router

import (
	"ege_platform/internal/config"
	"ege_platform/internal/pb"

	"github.com/pocketbase/pocketbase/core"
)

type Router struct {
	Pb  *pb.PB
	Cfg *config.Config
}

func NewRouter(p *pb.PB, cfg *config.Config) *Router {
	return &Router{
		Pb:  p,
		Cfg: cfg,
	}
}

func (r *Router) SetupRoutes() {
	r.Pb.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		e.Router.POST("/api/v1/auth_vk", r.AuthWithVK)
		e.Router.POST("/api/v1/auth_tg", r.AuthWithTG)
		e.Router.GET("/api/v1/access_token", r.GetAccessToken)
		e.Router.GET("/auth_vk_handle", r.InternalVKAuth)
		return nil
	})
}
