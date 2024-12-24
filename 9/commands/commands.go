package commands

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func Cd(args string) {
	path := args[3:]
	err := os.Chdir(path)
	if err != nil {
		fmt.Println(err)
	}
}

func Pwd() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(dir)

	return dir
}

func Echo(args string) {
	fmt.Println(args[5:])
}

func Kill(args string) {
	arg := args[5:]
	if args == "" {
		fmt.Println("Usage: kill <PID>")
		return
	}

	pid, err := strconv.Atoi(arg)
	if err != nil {
		fmt.Println(err)
		return
	}

	proc, err := os.FindProcess(pid)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = proc.Kill()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("Process %s killed\n", args)
	}
}

func Ps() string {
	procs, err := exec.Command("ps").Output()
	if err != nil {
		fmt.Println(err)
		return ""
	}

	fmt.Println(string(procs))
	return string(procs)
}

func Pipe(input string) {
	// ...
}

func External(input string) {
	args := strings.Fields(input)
	cmd := exec.Command(args[0], args[1:]...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Println(err)
	}
}
