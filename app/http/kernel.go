package http

import (
	"github.com/yafgo/framework/contracts/http"
	"github.com/yafgo/framework/http/middleware"
)

type Kernel struct {
}

// The application's global HTTP middleware stack.
// These middleware are run during every request to your application.
func (kernel *Kernel) Middleware() []http.Middleware {
	return []http.Middleware{
		middleware.Cors(),
	}
}
