package main

import (
	"fmt"
	"os"
)

func main() {
	condition, _ := exists("/home/server/xyz")
	fmt.Println(condition)
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
