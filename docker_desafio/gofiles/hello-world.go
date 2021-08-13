package main

import "fmt"

// func main() {
// 	fmt.Println("CODE EDUCATION ROCKS!")
// }

const (
	blue   = "\033[1;34m%s\033[0m"
	yellow = "\033[1;33m%s\033[0m"
	red    = "\033[1;31m%s\033[0m"
	green  = "\033[0;32m%s\033[0m"
)

func main() {
	fmt.Println("\033[2J")
	fmt.Printf(green, "#########################################\n")
	fmt.Printf(blue, "\tCODE")
	fmt.Printf(yellow, "\tEDUCATION")
	fmt.Printf(red, "\tROCKS!\n")
	fmt.Printf(green, "#########################################\n\n\n")
}
