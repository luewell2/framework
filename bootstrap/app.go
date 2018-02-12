package bootstrap

import (
	"fmt"

	"github.com/luewell/framework/Illuminate/Env"
)

func App() {
	fmt.Println("[bootstrap@App]: Initializing...")

	defer func() {
		fmt.Println("[bootstrap@App]: Ready!")
	}()

	Env.Load()
}
