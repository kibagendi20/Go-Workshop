package main

import (
	"fmt"
	"time"
)

func main() {
	//Declaring Multiple Variables with a Short Variable Declaration
	Debug, LogLevel, startUpTime := false, "info", time.Now()
	fmt.Println(Debug, LogLevel, startUpTime)

}
