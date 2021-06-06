package services

import (
	"github.com/CrowderSoup/EchoVue/server/config"
	"github.com/jinzhu/gorm"
)

// Database holds our gorm db
type database struct {
	connection *gorm.DB
	config     *config.DBConfig
}

func newDatabase(config *config.Config) *database {
	return &database{
		config: &config.DBConfig,
	}
}

func (d *database) connect() {
	db, err := gorm.Open(d.config.Dialect, d.config.ConnectionString)
	if err != nil {
		panic(err)
	}

	d.connection = db
}

func (d *database) migrate() {
	d.connection.AutoMigrate()
}
