package Http

import (
	"fmt"
	netHttp "net/http"
)

type Router struct{}

func (router Router) Listen(server string) {
	fmt.Printf("[Http\\Router@Listen]: Serving app at %s.\n", server)

	netHttp.ListenAndServe(server, nil)
}
