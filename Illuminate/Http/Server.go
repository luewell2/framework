package Http

import "fmt"

type Server struct{}

func (this Server) Init() {
	fmt.Println("Initializing the server")
}
