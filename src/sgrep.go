package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	argLength := len(os.Args[1:])

	if argLength < 1 {
		fmt.Println("Usage: sgrep text filename")
		return
	}

	text := os.Args[1]
	stackAtInit := 1

	var scanner *bufio.Scanner

	if argLength > 1 {
		fileName := os.Args[2]

		f, err := os.Open(fileName)
		if err != nil {
			log.Fatal(err)
		}
		scanner = bufio.NewScanner(f)
		defer f.Close()
	} else {
		scanner = bufio.NewScanner(os.Stdin)
	}

	stackAt := stackAtInit
	stack := false
	for scanner.Scan() {
		line := scanner.Text()
		if !stack && !strings.Contains(line, text) {
			continue
		} else if !stack || (stack && strings.Contains(line, text)) {
			fmt.Printf("%s\n", line)
			stack = true
			stackAt = stackAtInit
		} else if stack && stackAt > 0 {
			fmt.Printf("%s\n", line)
			stackAt--
		} else if stack && stackAt == 0 {
			if len(line) == 0 || (line[0] != 9 && strings.Index(line, "Caused by:") != 0 && line[0] != 32) {
				stack = false
				stackAt = stackAtInit
			} else {
				fmt.Printf("%s\n", line)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
