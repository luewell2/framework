package Http

import netHttp "net/http"

type Request struct {
	*netHttp.Request
}
