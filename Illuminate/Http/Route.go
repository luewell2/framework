package Http

type Route struct {
	Pattern string
	Method  string
	Handler Handler
}

func (this *Route) Match(pattern string) bool {
	return pattern == this.Pattern
}

type Handler func(*Context)
