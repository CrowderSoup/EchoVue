package server

import (
	"fmt"

	"github.com/CrowderSoup/EchoVue/server/config"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

// Server serves up our application
type server struct {
	assetsDir     string
	jwtPublicKey  string
	jwtPrivateKey string
}

// NewServerParams params for new server
type NewServerParams struct {
	fx.In

	Instance *echo.Echo
	Config   *config.Config
}

// NewServer returns a web server
func newServer(p NewServerParams) *server {
	return &server{
		assetsDir: p.Config.AssetsDir,
	}
}

func (s *server) config(e *echo.Echo) {
	// Remove Trailing Slashes from requests
	e.Pre(middleware.RemoveTrailingSlash())

	// Static dir
	e.Static("/css", fmt.Sprintf("%s/css", s.assetsDir))
	e.Static("/img", fmt.Sprintf("%s/img", s.assetsDir))
	e.Static("/js", fmt.Sprintf("%s/js", s.assetsDir))

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Front-end Entrypoint
	e.File("/", fmt.Sprintf("%s/index.html", s.assetsDir))

	// API Endpoints
	e.GET("/health", s.health)
}
