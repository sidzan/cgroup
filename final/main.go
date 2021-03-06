package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"os/exec"
	//"os/user"
	"syscall"
)

func main() {

	condition := Exists("/tmp/test")
	if condition == false {
		copyTEMPfile()
	}

	cgcreateExists := Exists("/usr/bin/cgcreate")
	if cgcreateExists == false {

		installCgcreate()
	} else {
		fmt.Println("cgcreate already installed ")
	}

	mounted := DefaultMountScript()
	fmt.Println(mounted)

	fmt.Println("at the  end of the main")
}
func RunInBackgrond() {
	fmt.Printf("hello world")

}
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
func installCgcreate() {
	binary, lookErr := exec.LookPath("sh")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "/tmp/test/installCgcreate.sh"}
	env := os.Environ()
	execErr1 := syscall.Exec(binary, args, env)
	if execErr1 != nil {
		panic(execErr1)
	}
}

func checkpid() {
	for {
		fmt.Println("This will check and execute the shell file")

		binary, lookErr := exec.LookPath("sh")
		if lookErr != nil {
			panic(lookErr)
		}
		args := []string{"sh", "/tmp/test/checkpid.sh"}
		env := os.Environ()
		execErr1 := syscall.Exec(binary, args, env)
		if execErr1 != nil {
			panic(execErr1)
		}
	}

}
func DefaultMountScript() int {
	fmt.Println("This will mount if groups has not been mounted file")
	binary, lookErr := exec.LookPath("sh")
	if lookErr != nil {
		panic(lookErr)
	}
	args := []string{"sh", "/tmp/test/mount.sh"}
	env := os.Environ()
	execErr1 := syscall.Exec(binary, args, env)
	if execErr1 != nil {
		panic(execErr1)
	}
	fmt.Println("end of Mount script")
	return 0
}

func CopyFile(source string, dest string) (err error) {
	sourcefile, err := os.Open(source)
	if err != nil {
		return err
	}

	defer sourcefile.Close()

	destfile, err := os.Create(dest)
	if err != nil {
		return err
	}

	defer destfile.Close()

	_, err = io.Copy(destfile, sourcefile)
	if err == nil {
		sourceinfo, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, sourceinfo.Mode())
		}

	}

	return
}

func CopyDir(source string, dest string) (err error) {

	// get properties of source dir
	sourceinfo, err := os.Stat(source)
	if err != nil {
		return err
	}

	// create dest dir

	err = os.MkdirAll(dest, sourceinfo.Mode())
	if err != nil {
		return err
	}

	directory, _ := os.Open(source)

	objects, err := directory.Readdir(-1)

	for _, obj := range objects {

		sourcefilepointer := source + "/" + obj.Name()

		destinationfilepointer := dest + "/" + obj.Name()

		if obj.IsDir() {
			// create sub-directories - recursively
			err = CopyDir(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			// perform copy
			err = CopyFile(sourcefilepointer, destinationfilepointer)
			if err != nil {
				fmt.Println(err)
			}
		}

	}
	return
}

func copyTEMPfile() {

	flag.Parse()           // get the source and destination directory
	source_dir := "./test" // get the source directory from 1st argument

	dest_dir := "/tmp/test" // get the destination directory from the 2nd argument

	fmt.Println("Source :" + source_dir)

	// check if the source dir exist
	src, err := os.Stat(source_dir)
	if err != nil {
		panic(err)
	}

	if !src.IsDir() {
		fmt.Println("Source is not a directory")
		os.Exit(1)
	}

	// create the destination directory
	fmt.Println("Destination :" + dest_dir)

	_, err = os.Open(dest_dir)
	if !os.IsNotExist(err) {
		fmt.Println("Destination directory already exists. Abort!")
		os.Exit(1)
	}

	err = CopyDir(source_dir, dest_dir)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Directory copied")
	}

}
