package resources

import (
	"fmt"
	"runtime"
)

func CpuCores() {
	fmt.Printf("\n %d \n", runtime.NumCPU())
}
