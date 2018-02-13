package Http

import (
	netHttp "net/http"
	"strings"
)

type Request struct {
	*netHttp.Request
}

func (this *Request) FullUrl() string {
	return this.URL.Path
}

func (this *Request) Split() []string {
	return strings.Split(this.FullUrl(), "/")
}
