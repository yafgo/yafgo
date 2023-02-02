package console

import (
	"github.com/yafgo/framework/contracts/console"
	"github.com/yafgo/framework/contracts/schedule"
)

type Kernel struct {
}

func (kernel *Kernel) Schedule() []schedule.Event {
	return []schedule.Event{}
}

func (kernel *Kernel) Commands() []console.Command {
	return []console.Command{}
}
