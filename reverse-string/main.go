package main

import (
	"fmt"
	"slices"
	"strings"
)

func main() {
	fmt.Println(reverse("abc"))
	fmt.Println(reverse2("abc"))
}

func reverse(s string) (result string) {
	slice := strings.Split(s, "")
	reversed := ""
	for _, char := range slice {
		reversed = char + reversed
	}
	return reversed
}

func reverse2(s string) (result string) {
	slice := strings.Split(s, "")
	slices.Reverse(slice)
	result = strings.Join(slice, "")
	return result
}
