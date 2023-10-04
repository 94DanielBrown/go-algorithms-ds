package main

import (
	"fmt"
	"github.com/94DanielBrown/helpers"
)

func main() {
	fmt.Println(reverse("abc"))
}

func reverse(s string) (result string) {
	return helpers.Reverse(s)
	//	return "", nil
}
