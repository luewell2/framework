package Http

import netHttp "net/http"

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
