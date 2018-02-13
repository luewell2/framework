package Http

import (
	"fmt"
	netHttp "net/http"
)

type Router struct {
	Routes []Route
}

func (this *Router) ServeHTTP(response netHttp.ResponseWriter, request *netHttp.Request) {
	routeFound := false
	http := NewContext(response, request)

	if http.FullUrl() != "/favicon.ico" {
		for _, route := range this.Routes {
			if route.Match(http.FullUrl()) {
				route.Handler(http)
				routeFound = true

				break
			}
		}

		if !routeFound {
			http.ThrowHttpCode(404)
		}
	}
}

func (this *Router) Listen(server string) {
	fmt.Printf("[Http\\Router@Listen]: Serving app at %s.\n", server)

	netHttp.ListenAndServe(server, this)
}

func Init() *Router {
	http := &Router{}

	return http
}

func (this *Router) pushRoute(pattern string, method string, handler Handler) {
	route := Route{
		Method:  method,
		Pattern: pattern,
		Handler: handler,
	}

	this.Routes = append(this.Routes, route)
}

/* Methods */
func (this *Router) Get(pattern string, handler Handler) {
	this.pushRoute(pattern, "GET", handler)
}

func (this *Router) Post(pattern string, handler Handler) {
	this.pushRoute(pattern, "POST", handler)
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

	return http
}
