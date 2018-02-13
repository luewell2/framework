package Http

import (
	"encoding/json"
	"encoding/xml"
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
	this.Header().Set("Content-Type", typ)
}

func (this *Response) SetStatusCode(code string) {
	this.Header().Set("Status Code", code)
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

/* Methods Response */

func (this *Response) ResponseString(content string) {
	this.Write([]byte(content))
}

func (this *Response) ResponseJSON(content interface{}) {
	jsonContent, _ := json.Marshal(content)

	this.SetContentType("application/json")
	this.Write(jsonContent)
}

func (this *Response) ResponseXML(content interface{}) {
	xmlContent, _ := xml.Marshal(content)

	this.SetContentType("application/xml")
	this.Write(xmlContent)
}
