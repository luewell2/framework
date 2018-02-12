package Env

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func Load() {
	fmt.Println("[Env@Load]: Initializing...")

	loadFile()
}

func loadFile() {
	env := readFile()

	setEnv(env)
}

func readFile() (envMap map[string]string) {
	file, err := os.Open(".env")

	defer file.Close()

	if err != nil {
		fmt.Println("[Env@readFile]: Error loading the file .env")
		return
	}

	return parseFile(file)
}

func parseFile(file io.Reader) (envMap map[string]string) {
	envMap = make(map[string]string)

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()

		if !isIgnoreLine(line) {
			key, value := parseLine(line)

			envMap[key] = value
		}
	}

	return
}

func isIgnoreLine(line string) bool {
	line = strings.Trim(line, " \n\t")

	return (len(line) == 0) || strings.HasPrefix(line, "#")
}

func parseLine(line string) (key string, value string) {
	splitLine := strings.SplitN(line, "=", 2)

	key = strings.Trim(splitLine[0], " ")
	value = strings.Trim(strings.SplitN(strings.Trim(splitLine[1], " "), "#", 2)[0], " ")

	return
}

func setEnv(env map[string]string) {
	for key, value := range env {
		os.Setenv(key, value)
	}
}
