package main

import (
	"github.com/CrowderSoup/EchoVue/server/config"
	"github.com/CrowderSoup/EchoVue/server/server"

	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options(
		config.Module,
		server.Module,
	)
	app := fx.New(
		bundle,
		fx.Invoke(server.InvokeServer),
	)

	app.Run()

	<-app.Done()
}
