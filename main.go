package main

import "fmt"

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
}
