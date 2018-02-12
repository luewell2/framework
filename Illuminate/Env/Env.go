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

	fmt.Println(env)
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
	var lines []string

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	for _, line := range lines {
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
