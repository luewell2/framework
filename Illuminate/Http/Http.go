package Http

import (
	"fmt"
	netHttp "net/http"
)

type Router struct{}

func (this *Router) ServeHTTP(response netHttp.ResponseWriter, request *netHttp.Request) {
	http := NewContext(response, request)

	fmt.Println(http.FullUrl(), http.Split())
}

func (this *Router) Listen(server string) {
	fmt.Printf("[Http\\Router@Listen]: Serving app at %s.\n", server)

	netHttp.ListenAndServe(server, this)
}

func Init() *Router {
	http := &Router{}

	return http
}

/* Http.Context */
type Context struct {
	Request
	Response
}

func NewContext(response netHttp.ResponseWriter, request *netHttp.Request) *Context {
	http := &Context{}

	http.Request.Request = request
	http.Response.ResponseWriter = response

	http.Response.SetContentType("text/html")
	http.Response.SetStatusCode("200")

	http.Init()

	return http
}

func (http *Context) Init() {}
