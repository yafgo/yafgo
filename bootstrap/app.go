package bootstrap

import (
	"github.com/yafgo/framework/foundation"

	"github.com/yafgo/yafgo/config"
)

func Boot() {
	app := foundation.Application{}

	//Bootstrap the application
	app.Boot()

	//Bootstrap the config.
	config.Boot()
}
