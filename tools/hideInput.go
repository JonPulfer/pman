package tools

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"
)

var sttyCmd string = "/bin/stty"
var sttyArgvEchoOff []string = []string{"stty", "-echo"}
var sttyArgvEchoOn []string = []string{"stty", "echo"}
var waitStatus syscall.WaitStatus = 0

func echoOff(pa syscall.ProcAttr) int {
	pid, err := syscall.ForkExec(sttyCmd, sttyArgvEchoOff, &pa)
	if err != nil {
		fmt.Printf("Error setting echo off: %s\n", err)
	}

	return pid
}

func echoOn(pa syscall.ProcAttr) {
	pid, err := syscall.ForkExec(sttyCmd, sttyArgvEchoOn, &pa)
	if err == nil {
		syscall.Wait4(pid, &waitStatus, 0, nil)
	} else {
		fmt.Printf("Error setting echo on: %s\n", err)
	}
}
func catchSignal(pa syscall.ProcAttr, sig chan os.Signal, brk chan bool) {
	select {
	case <-sig:
		echoOn(pa)
		os.Exit(-1)
	case <-brk:
	}
}

// HideInput sets echo off in the terminal to enable a string to be
// entered more securely.
func HideInput(prompt string) string {
	var secret string
	sig := make(chan os.Signal, 10)
	brk := make(chan bool)

	fmt.Printf(prompt)
	fd := []uintptr{
		os.Stdin.Fd(),
		os.Stdout.Fd(),
		os.Stderr.Fd(),
	}
	ProcessAttributes := syscall.ProcAttr{
		Dir:   "",
		Files: fd,
	}

	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGKILL,
		syscall.SIGQUIT, syscall.SIGTERM)
	go catchSignal(ProcessAttributes, sig, brk)

	pid := echoOff(ProcessAttributes)
	read := bufio.NewReader(os.Stdin)
	syscall.Wait4(pid, &waitStatus, 0, nil)

	input, err := read.ReadString('\n')
	if err == nil {
		secret = strings.TrimSpace(input)
	} else {
		fmt.Printf("Error reading string from stdin: %v\n", err)
	}

	defer close(brk)
	defer echoOn(ProcessAttributes)

	return secret

}
