package Config

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type Env struct{}

func (env Env) Load() {
	fmt.Println("[Env@Load]: Initializing...")

	env.loadFile()
}

func (env Env) loadFile() {
	env.readFile()
}

func (env Env) readFile() (envMap map[string]string) {
	file, err := os.Open(".env")

	defer file.Close()

	if err != nil {
		fmt.Println("[Env@readFile]: Error loading the file .env")
		return
	}

	return env.parseFile(file)
}

func (env Env) parseFile(file io.Reader) (envMap map[string]string) {
	envMap = make(map[string]string)
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	fmt.Println(lines)

	return
}
