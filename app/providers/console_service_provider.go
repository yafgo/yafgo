package providers

import (
	"github.com/yafgo/framework/facades"

	"github.com/yafgo/yafgo/app/console"
)

type ConsoleServiceProvider struct {
}

func (receiver *ConsoleServiceProvider) Register() {
	kernel := console.Kernel{}
	// facades.Schedule.Register(kernel.Schedule())
	facades.Artisan.Register(kernel.Commands())
}

func (receiver *ConsoleServiceProvider) Boot() {

}
