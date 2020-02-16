package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"parking_lot/infra/cli"
	"strings"
)

func main() {
	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) > 0 {
		// read commands from file
		fileName := argsWithoutProg[0]
		ReadAndProcessFromFile(fileName)
	} else {
		//interactive session now
		ReadAndProcessStdIn()
	}
}

//ReadAndProcessFromInput Common Function with dependency of input injected via argument
func ReadAndProcessFromInput(input io.Reader) {
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		command := strings.ToLower(scanner.Text())
		cli.ProcessCommand(command)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

//ReadAndProcessStdIn Function to read from stdin and process the command with arguments
func ReadAndProcessStdIn() {
	ReadAndProcessFromInput(os.Stdin)
}

//ReadAndProcessFromFile Function to read from file line by line and process the command with arguments
func ReadAndProcessFromFile(filePath string) {
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	ReadAndProcessFromInput(f)
}
