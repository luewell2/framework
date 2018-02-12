package Http

import (
	"fmt"
	netHttp "net/http"
)

type Router struct {
	netHttp.Handler
}

var methods = []string{
	netHttp.MethodGet,
	netHttp.MethodHead,
	netHttp.MethodPost,
	netHttp.MethodPut,
	netHttp.MethodPatch,
	netHttp.MethodDelete,
	netHttp.MethodConnect,
	netHttp.MethodOptions,
	netHttp.MethodTrace,
}

func (this Router) Listen(server string) {
	fmt.Printf("[Http\\Router@Listen]: Serving app at %s.\n", server)

	netHttp.ListenAndServe(server, this)
}
