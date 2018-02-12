package bootstrap

import (
	"fmt"

	"github.com/luewell/framework/Illuminate/Config"
)

func App() {
	fmt.Println("[bootstrap@App]: Initializing...")

	defer func() {
		fmt.Println("[bootstrap@App]: Ready!")
	}()

	Config.Env{}.Load()
}
