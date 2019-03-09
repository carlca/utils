package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	pathEnv := os.Getenv("PATH")
	paths := strings.Split(pathEnv, ":")
	sort.Strings(paths)
	for _, path := range paths {
		if len(os.Args) == 1 || filterNeeded(path) {
			fmt.Println(path)
		}            
	}
}

func filterNeeded(path string) bool {
	return strings.Contains(strings.ToLower(path), strings.ToLower(os.Args[1]))
}