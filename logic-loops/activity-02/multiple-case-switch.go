package main

import (
	"fmt"
	"time"
)

func main() {
	var dayBorn = time.Sunday

	switch dayBorn {
	case time.Monday, time.Tuesday, time.Wednesday, time.Thursday, time.Friday:
		fmt.Println("Norn on a weekday")
	case time.Saturday, time.Sunday:
		fmt.Println("Born on a weekend")
	default:
		fmt.Println("Error, day not valid")

	}
}
