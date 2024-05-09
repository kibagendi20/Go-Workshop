package main

import "fmt"

var foo string = "bar" // declared at the package level

func main() {
	var buz string = "qux" // declared in the function level
	fmt.Println(foo, buz)
}
