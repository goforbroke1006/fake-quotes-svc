package main

import (
	"fmt"

	"github.com/goforbroke1006/fake-quotes-svc/cmd"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func main() {
	fmt.Println("Hello, World!!!")
	fmt.Println("version:", version)
	fmt.Println("commit:", commit)
	fmt.Println("date:", date)

	cmd.Execute()
}
