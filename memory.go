package resources

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

func MemoryPath() {
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"ls", "/home/sijan/cgroup/resources/"}
	env := os.Environ()
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}

}

func Memory() {
	fmt.Println("This should detail the memory used")

	fmt.Println("This will execute sh file")
	binary, lookErr := exec.LookPath("sh")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "/home/sijan/cgroup/resources/memory.sh"}
	env := os.Environ()
	execErr1 := syscall.Exec(binary, args, env)
	if execErr1 != nil {
		panic(execErr1)
	}

}
