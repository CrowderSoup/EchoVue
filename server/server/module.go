package server

import (
	"context"
	"fmt"

	"github.com/CrowderSoup/EchoVue/server/config"
	"github.com/labstack/echo/v4"
	"go.uber.org/fx"
)

// InvokeServer starts up our web server
func InvokeServer(
	lc fx.Lifecycle,
	c *config.Config,
	instance *echo.Echo,
	server *server,
) {
	lc.Append(
		fx.Hook{
			OnStart: func(ctx context.Context) error {
				server.config(instance)

				go instance.Start(fmt.Sprintf(":%d", c.Port))
				return nil
			},
			OnStop: func(ctx context.Context) error {
				return instance.Close()
			},
		},
	)
}

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		echo.New,
		newServer,
	),
)
