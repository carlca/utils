package dirx

import (
	"flag"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

var indent int
var err error

// Dirx is an early attempt at a useful util
func Dirx() {
	flag.String("path", "/users/carlca/code/go", "Starting path for recursive scan")
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
	dirList = append(dirList, scanDir(path, &level)...)
	for _, file := range dirList {
		fmt.Println(file)
	}
}

func scanDir(path string, level *int) []string {
	*level++
	pad := strings.Repeat(" ", *level*indent)
	result := []string{}
	files, _ := ioutil.ReadDir(path)
	for _, f := range files {
		if f.IsDir() {
			result = append(result, pad+f.Name()+" [dir]")
			result = append(result, scanDir(path+"/"+f.Name(), level)...)
		} else {
			result = append(result, pad+f.Name())
		}
	}
	*level--
	return result
}
