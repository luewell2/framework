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
	readFile()
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
			// fmt.Println(line)
		}
	}

	return
}

func isIgnoreLine(line string) bool {
	trim := strings.Trim(line, " \n\t")

	fmt.Println(trim)

	return true
}
