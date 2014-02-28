package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"syscall"
)

func getExitCode(err error) int {
	return 1
}
func findGo() string {
	goroot := os.Getenv("GOROOT")
	if goroot != "" {
		if runtime.GOOS == "windows" {
			return filepath.Join(goroot, "bin", "go.exe")
		} else {
			return filepath.Join(goroot, "bin", "go")
		}
	}

	p, err := filepath.Abs(os.Args[0])
	if err != nil {
		return ""
	}
	if runtime.GOOS == "windows" {
		pathext := strings.Split(os.Getenv("PATHEXT"),  string(filepath.ListSeparator))
		for _, path := range strings.Split(os.Getenv("PATH"), string(filepath.ListSeparator)) {
			for _, ext := range pathext {
				fullp := filepath.Join(path, "go." + ext)
				if fullp != p {
					if _, err = os.Stat(fullp); err == nil {
						return fullp
					}
				}
			}
		}
	} else {
		for _, path := range strings.Split(os.Getenv("PATH"), string(filepath.ListSeparator)) {
			fullp := filepath.Join(path, "go")
			if fullp != p {
				if _, err = os.Stat(fullp); err == nil {
					return fullp
				}
			}
		}
	}
	return ""
}

var replacer = strings.NewReplacer(
	"a", "ɐ",
	"b", "q",
	"c", "ɔ",
	"d", "p",
	"e", "ǝ",
	"f", "ɟ",
	"g", "ɓ",
	"h", "ɥ",
	"i", "ı",
	"j", "ɾ",
	"k", "ʞ",
	"l", "l",
	"m", "ɯ",
	"n", "u",
	"o", "o",
	"p", "d",
	"q", "b",
	"r", "ɹ",
	"s", "s",
	"t", "ʇ",
	"u", "n",
	"v", "ʌ",
	"w", "ʍ",
	"x", "x",
	"y", "ʎ",
	"z", "z",
	"1", "⇂",
	"2", "ᄅ",
	"3", "Ɛ",
	"4", "ㄣ",
	"5", "ގ",
	"6", "9",
	"7", "ㄥ",
	"8", "8",
	"9", "6",
)

func flip(name string) string {
	var s []string
	for _, c := range []rune(strings.ToLower(name)) {
		s = append(s, string(c))
	}
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        s[i], s[j] = s[j], s[i]
    }
	return replacer.Replace(strings.Join(s, ""))
}

func main() {
	nArgs := len(os.Args)
	if nArgs > 3 && os.Args[1] == "to" && os.Args[2] == "hell" {
		name := os.Args[3]
		var killed int
		if nArgs > 3 {
			killed = killall(name)
		}
		if killed <= 0 {
			fmt.Println("(；￣Д￣) . o O( It’s not very effective... )")
		} else {
			fmt.Printf("%s (x%d)", flip(name), killed)
		}
	} else {
		gocmd := findGo()
		if gocmd != "" {
			cmd := exec.Command(gocmd, os.Args[1:]...)
			cmd.Stdin = os.Stdin
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				if status, ok := cmd.ProcessState.Sys().(syscall.WaitStatus); ok {
					os.Exit(status.ExitStatus())
				}
				os.Exit(1)
			}
			os.Exit(0)
		}
	}
}
