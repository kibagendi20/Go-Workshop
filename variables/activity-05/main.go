package main

import (
	"fmt"
	"time"
)

// Declaring Multiple Variables from a Function
func getConfig() (bool, string, time.Time) {
	return false, "info", time.Now()
}
func main() {
	Debug, LogLevel, startUpTime := getConfig()
	fmt.Println(Debug, LogLevel, startUpTime)
}
