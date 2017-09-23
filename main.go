// main.go
package main

import "fmt"

var version string = "undefined"
var commit string = "000000000"
var date string = "01.01.1972"

func main() {
	fmt.Printf("Version:    %v\n", version)
	fmt.Printf("Commit:     %v\n", commit)
	fmt.Printf("Build Date: %v\n", date)
}
