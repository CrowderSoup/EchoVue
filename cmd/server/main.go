package main

import (
	"github.com/CrowderSoup/EchoVue/server/config"
	web "github.com/CrowderSoup/EchoVue/server/server"

	"go.uber.org/fx"
)

func main() {
	bundle := fx.Options(
		config.Module,
		web.Module,
	)
	app := fx.New(
		bundle,
		fx.Invoke(web.InvokeServer),
	)

	app.Run()

	<-app.Done()
}
