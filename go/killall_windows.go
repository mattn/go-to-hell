package main

import (
	"os/exec"
	"strings"
)

func killall(name string) int {
	b, err := exec.Command("taskkill", "/im", name + "*").Output()
	if err != nil {
		return 0
	}
	return len(strings.Split(string(b), "\n")) - 1
}
