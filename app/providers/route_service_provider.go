package providers

import (
	"github.com/yafgo/framework/facades"

	"github.com/yafgo/yafgo/app/http"
	"github.com/yafgo/yafgo/routes"
)

type RouteServiceProvider struct {
}

func (receiver *RouteServiceProvider) Register() {
	//Add HTTP middlewares
	kernel := http.Kernel{}
	facades.Route.GlobalMiddleware(kernel.Middleware()...)
}

func (receiver *RouteServiceProvider) Boot() {
	//Add routes
	routes.Web()
}
