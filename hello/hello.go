package main

import "fmt"
import "github.com/murugandineshj2ee/GoLangHelloWorld/hello/morestrings"
import "github.com/google/go-cmp/cmp"

func main() {
	fmt.Println("Hello, world.")
	fmt.Println(morestrings.ReverseRunes("!oG ,olleH"))
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}