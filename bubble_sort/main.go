package main

import "fmt"

func main() {
	fmt.Println(bubbleSort([]int{1, 2, 3, 4, 5}))

}

func bubbleSort(slice []int) []int {
	n := len(slice)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if slice[j] > slice[j+1] {
				lesser := slice[j+1]
				slice[j+1] = slice[j]
				slice[j] = lesser
			}
		}
	}

	return slice
}
