package main

import (
	"fmt"
	"os"
)

func main() {
	a := Exists("/usr/bin/cgcreate")
	fmt.Println(a)
}
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
