package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(reverse("abc"))
}

func reverse(s string) (result string) {
	slice := strings.Split(s, "")
	reversed := ""
	for _, char := range slice {
		reversed = char + reversed
	}
	fmt.Println(reversed)
	return reversed
}
