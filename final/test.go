package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func main() {
	DefaultMountScript()
}
func DefaultMountScript() int {
	fmt.Println("This will mount if groups has not been mounted file")
	binary, lookErr := exec.LookPath("sh")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "/tmp/test/mount.sh"}
	env := os.Environ()
	syscall.Exec(binary, args, env)

	fmt.Println("end of Mount script")
	return 0
}
