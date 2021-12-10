package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	buff := []string{}
	stdin := bufio.NewScanner(os.Stdin)
	for stdin.Scan() {
		text := stdin.Text()
		switch {
		case strings.HasPrefix(text, "package "):
			buff = append(buff, "# "+text)

		case strings.HasPrefix(text, "VARIABLES"):
			buff = append(buff, "## "+text)

		case strings.HasPrefix(text, "FUNCTIONS"):
			buff = append(buff, "## "+text)

		case strings.HasPrefix(text, "TYPES"):
			buff = append(buff, "## "+text)

		case strings.HasPrefix(text, "func "):
			buff = append(buff, "### "+text)

		case strings.HasPrefix(text, "var ("):
			buff = append(buff, "```go")
			buff = append(buff, text)

		case strings.HasPrefix(text, "type "):
			buff = append(buff, "```go")
			buff = append(buff, text)

		case strings.HasPrefix(text, ")"):
			buff = append(buff, text)
			buff = append(buff, "```")

		case strings.HasPrefix(text, "}"):
			buff = append(buff, text)
			buff = append(buff, "```")

		default:
			buff = append(buff, text)
		}
	}

	fmt.Fprint(os.Stdout, strings.Join(buff, "\n"))
}
