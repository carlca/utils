package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
	//"path/filepath"
	//"github.com/carlca/types"
)

var indent int
var err error

func main() {
	flag.String("path", "c:/go", "Starting path for recursive scan")
	flag.Int("indent", 2, "Indents per directory level")
	flag.Parse()
	path := flag.Args()[0]
	indent, err = strconv.Atoi(flag.Args()[1])
	if err != nil {
		panic(err)
	}
	fmt.Println(path)
	dirList := []string{}
	level := 0
	dirList = append(dirList, ScanDir(path, &level)...)
	for _, file := range dirList {
		fmt.Println(file)
	}
}

func ScanDir(path string, level *int) []string {
	*level += 1
	pad := strings.Repeat(" ", *level*indent)
	result := []string{}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			result = append(result, pad+f.Name()+" [dir]")
			result = append(result, ScanDir(path+"/"+f.Name(), level)...)
		} else {
			result = append(result, pad+f.Name())
		}
	}
	*level -= 1
	return result
}
