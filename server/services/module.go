package services

import (
	"context"

	"go.uber.org/fx"

	// Driver for gorm
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Module provided to fx
var Module = fx.Options(
	fx.Provide(
		newDatabase,
	),
)

// InvokeDB opens / manages our database connection
func InvokeDB(lc fx.Lifecycle, d *database) {
	lc.Append(
		fx.Hook{
			OnStart: func(context context.Context) error {
				d.connect()
				d.migrate()

				// Autoload Relationships
				d.connection.Set("gorm:auto_preload", true)

				return nil
			},
			OnStop: func(context context.Context) error {
				// Ensure we close our connection on app shutdown
				d.connection.Close()
				return nil
			},
		},
	)
}
