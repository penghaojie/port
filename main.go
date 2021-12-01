package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) == 1 {
		List(0)
	} else {
		if isNumber(args[1]) {
			port, _ := strconv.Atoi(args[1])
			List(uint32(port))
		} else if args[1] == "-c" || args[1] == "-C" {
			if len(args) < 3 || !isNumber(args[2]) {
				errorCommand()
			}
			port, _ := strconv.Atoi(args[2])
			KillProcessByPort(uint32(port))
		} else {
			errorCommand()
		}
	}
}

func isNumber(s string) bool {
	for i := 0; i < len(s); i++ {
		if s[i] < '0' || s[i] > '9' {
			return false
		}
	}
	return true
}

func errorCommand() {
	fmt.Println("unrecognizable command formatter")
}
