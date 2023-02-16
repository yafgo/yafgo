package config

import (
	"github.com/yafgo/framework/auth"
	"github.com/yafgo/framework/cache"
	"github.com/yafgo/framework/console"
	"github.com/yafgo/framework/contracts"
	"github.com/yafgo/framework/database"
	"github.com/yafgo/framework/event"
	"github.com/yafgo/framework/facades"
	"github.com/yafgo/framework/filesystem"
	"github.com/yafgo/framework/grpc"
	"github.com/yafgo/framework/http"
	"github.com/yafgo/framework/log"
	"github.com/yafgo/framework/mail"
	"github.com/yafgo/framework/queue"
	"github.com/yafgo/framework/route"
	"github.com/yafgo/framework/schedule"
	"github.com/yafgo/framework/validation"

	"github.com/yafgo/yafgo/app/providers"
)

// Boot Start all init methods of the current folder to bootstrap all config.
func Boot() {}

func init() {
	config := facades.Config
	config.Add("app", map[string]any{
		// Application Name
		//
		// This value is the name of your application. This value is used when the
		// framework needs to place the application's name in a notification or
		// any other location as required by the application or its packages.
		"name": config.Env("APP_NAME", "Yafgo"),

		// Application Environment
		//
		// This value determines the "environment" your application is currently
		// running in. This may determine how you prefer to configure various
		// services the application utilizes. Set this in your ".env" file.
		"env": config.Env("APP_ENV", "production"),

		// Application Debug Mode
		"debug": config.Env("APP_DEBUG", false),

		// Application Timezone
		//
		// Here you may specify the default timezone for your application, which
		// will be used by the PHP date and date-time functions. We have gone
		// ahead and set this to a sensible default for you out of the box.
		"timezone": "UTC",

		// Encryption Key
		//
		// 32 character string, otherwise these encrypted strings
		// will not be safe. Please do this before deploying an application!
		"key": config.Env("APP_KEY", ""),

		// Application URL
		"url": config.Env("APP_URL", "http://localhost"),

		// Application host, http server listening address.
		"host": config.Env("APP_HOST", "0.0.0.0:3000"),

		// Autoload service providers
		//
		// The service providers listed here will be automatically loaded on the
		// request to your application. Feel free to add your own services to
		// this array to grant expanded functionality to your applications.
		"providers": []contracts.ServiceProvider{
			&log.ServiceProvider{},
			&console.ServiceProvider{},
			&database.ServiceProvider{},
			&cache.ServiceProvider{},
			&http.ServiceProvider{},
			&route.ServiceProvider{},
			&schedule.ServiceProvider{},
			&event.ServiceProvider{},
			&queue.ServiceProvider{},
			&grpc.ServiceProvider{},
			&mail.ServiceProvider{},
			&auth.ServiceProvider{},
			&filesystem.ServiceProvider{},
			&validation.ServiceProvider{},

			&providers.AppServiceProvider{},
			&providers.AuthServiceProvider{},
			&providers.RouteServiceProvider{},
			&providers.GrpcServiceProvider{},
			&providers.ConsoleServiceProvider{},
			&providers.QueueServiceProvider{},
			&providers.EventServiceProvider{},
			&providers.ValidationServiceProvider{},
		},
	})
}
