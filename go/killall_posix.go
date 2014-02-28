// +build !windows

package main

import (
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func killall(name string) int {
	b, err := exec.Command("pgrep", name).Output()
	if err != nil {
		return 0
	}
	lines := strings.Split(string(b), "\n")
	for _, line := range lines {
		pid, err := strconv.Atoi(line)
		if err == nil {
			syscall.Kill(pid, syscall.SIGKILL)
		}
	}
	return len(lines) - 1
}
