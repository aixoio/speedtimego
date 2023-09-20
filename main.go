package main

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func main() {
	
	if len(os.Args) <= 1 {
		fmt.Println("Usage: <command>")
		return
	}

	fmt.Println("\nStarting...\n")

	starttime := time.Now().UnixNano()

	cmddat := os.Args[1:]

	cmdName := cmddat[0]
	cmdArgs := cmddat[1:]

	cmd := exec.Command(cmdName, cmdArgs...)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout

	cmd.Run()

	endtime := time.Now().UnixNano()

	fmt.Println("\nEnding...\n")

	time := endtime - starttime

	fmt.Printf("Command ran in %d nanoseconds!\n", time)
	fmt.Printf("Command ran in %d milliseconds!\n", time / 1000000)
	fmt.Printf("Start time: %d unix nanoseconds\n", starttime)
	fmt.Printf("End time: %d unix nanoseconds\n", endtime)

}
