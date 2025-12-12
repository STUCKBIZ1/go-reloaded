package main

import (
	"fmt"
	"os"
	"strings"

	go_reloaded "go_reloaded/servers"
)

func main() {
	if len(os.Args) < 3 || len(os.Args) > 3 {
		fmt.Println("usage: program input.txt output.txt")
		return
	}

	input := os.Args[1]
	output := os.Args[2]

	if !strings.HasSuffix(output, ".txt") {
		fmt.Println("Output file must end with .txt")
		return
	}

	content, err := os.ReadFile(input)
	if err != nil {
		fmt.Println(err)
		return
	}

	lines := strings.Split(string(content), "\n")
	result := go_reloaded.Manager(lines)

	err = os.WriteFile(output, []byte(result), 0o775)
	if err != nil {
		fmt.Println(err)
		return
	}
}
