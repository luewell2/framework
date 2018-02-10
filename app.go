package framework

import "github.com/luewell/framework/Illuminate/Http"

func Run() {
	Http.Server{}.Init()
}
