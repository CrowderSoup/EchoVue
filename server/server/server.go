package web

import (
	"context"
	"fmt"

	"github.com/CrowderSoup/EchoVue/server/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

// Server serves up our application
type Server struct {
	echo *echo.Echo
}

// NewServerParams params for new server
type NewServerParams struct {
	fx.In

	Instance *echo.Echo
	Config   *config.Config
}

// NewWebServer returns a web server
func NewWebServer(p NewServerParams) *Server {
	// Remove Trailing Slashes from requests
	p.Instance.Pre(middleware.RemoveTrailingSlash())

	// Static dir
	p.Instance.Static("/css", fmt.Sprintf("%s/css", p.Config.AssetsDir))
	p.Instance.Static("/img", fmt.Sprintf("%s/img", p.Config.AssetsDir))
	p.Instance.Static("/js", fmt.Sprintf("%s/js", p.Config.AssetsDir))

	// Middleware
	p.Instance.Use(middleware.Logger())
	p.Instance.Use(middleware.Recover())

	// Front-end Entrypoint
	p.Instance.File("/", fmt.Sprintf("%s/index.html", p.Config.AssetsDir))

	return &Server{
		echo: p.Instance,
	}
}

// InvokeServer starts up our web server
func InvokeServer(
	lc fx.Lifecycle,
	c *config.Config,
	server *Server,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				go server.echo.Start(fmt.Sprintf(":%d", c.Port))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return server.echo.Close()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		echo.New,
		NewWebServer,
	),
)
