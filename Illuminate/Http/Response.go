package Http

import netHttp "net/http"

type Response struct {
	netHttp.ResponseWriter
}

func (this *Response) SetContentType(typ string) {
	this.Header().Set("Content-Type", typ)
}

func (this *Response) SetStatusCode(code string) {
	this.Header().Set("Status Code", code)
}
