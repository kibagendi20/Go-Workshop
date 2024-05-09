package main

import "fmt"

var helloList = []string{
	"Hello",
	"Hi",
	"Mambo",
	"Sasa",
	"Καλημέρα κόσμε",
	"‫مالس‬ ‫ایند‬‎",
}

func main() {
	fmt.Println(helloList[5])

}
