package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string = "This is my practice"
	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?\n", str, "Th")
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

