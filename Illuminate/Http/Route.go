package Http

type Route struct {
	Pattern string
	Method  string
	Handler Handler
}

type Handler func(*Context)
