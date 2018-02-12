package bootstrap

import (
	"fmt"
	"os"

	"github.com/luewell/framework/Illuminate/Env"
	"github.com/luewell/framework/Illuminate/Http"
)

func App() {
	fmt.Println("[bootstrap@App]: Initializing...")

	defer func() {
		fmt.Println("[bootstrap@App]: Ready!")
	}()

	Env.Load()
	Http.Router{}.Listen(os.Getenv("APP_URL"))
}
