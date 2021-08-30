package utils

import (
	"fmt"
	"runtime"
	"strings"
)

//returns the root path directory for the working application
func GetLocalEnv() string {
	_, filename, _, _ := runtime.Caller(0)
	tmp := strings.Split(filename, "/")
	var path string
	for i := 0; i < len(tmp)-2; i++ {
		path += tmp[i] + "/"
	}
	fmt.Println(path)
	return path
}
