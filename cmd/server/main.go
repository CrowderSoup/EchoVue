package main

import (
	"github.com/CrowderSoup/EchoVue/server/config"
	"github.com/CrowderSoup/EchoVue/server/server"
	"github.com/CrowderSoup/EchoVue/server/services"

	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options(
		config.Module,
		services.Module,
		server.Module,
	)
	app := fx.New(
		bundle,
		fx.Invoke(services.InvokeDB),
		fx.Invoke(server.InvokeServer),
	)

	app.Run()

	<-app.Done()
}
