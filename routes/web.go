package routes

import (
	"github.com/yafgo/framework/contracts/http"
	"github.com/yafgo/framework/facades"

	"github.com/yafgo/yafgo/app/http/controllers"
)

func Web() {
	facades.Route.Get("/", func(ctx http.Context) {
		ctx.Response().Json(200, http.Json{
			"Hello": "Yafgo",
		})
	})

	userController := controllers.NewUserController()
	facades.Route.Get("/users/{id}", userController.Show)
}
