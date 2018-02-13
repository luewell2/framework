package Http

import (
	"fmt"
	netHttp "net/http"
)

var httpCodeMessage = map[int]string{
	404: `Not Found`,
}

type Response struct {
	netHttp.ResponseWriter
}

func (this *Response) SetContentType(typ string) {
	this.Header().Set(`Content-Type`, typ)
}

func (this *Response) SetStatusCode(code string) {
	this.Header().Set(`Status Code`, code)
}

func (this *Response) Throw(code int, message string) {
	this.SetStatusCode(fmt.Sprintf("%s", code))

	fmt.Fprintf(this.ResponseWriter, fmt.Sprintf("%d : %s", code, message))
}

func (this *Response) ThrowHttpCode(code int) {
	msg := httpCodeMessage[code]

	if len(msg) == 0 {
		this.Throw(500, fmt.Sprintf("Unknow http code [%d]", code))
	} else {
		this.Throw(code, httpCodeMessage[code])
	}
}
