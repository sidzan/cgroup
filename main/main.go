package main

import (
	"../cgroup"
	"../resources"
	"fmt"
)

func main() {
	fmt.Println("This is a package that manages the resources used by Docker container")
	cgroup.PrintThis()
	cgroup.PrintThis2()
	fmt.Println()

	resources.CpuCores()
	resources.Memory()
	resources.MemoryPath()

}
