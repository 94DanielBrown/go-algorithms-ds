package main

func main() {
}

func maxChars(s string) (result string) {
	var maxFrequency int
	var maxChar rune

	charMap := make(map[rune]int)

	for _, char := range s {
		charMap[char]++
	}

	for char, count := range charMap {
		if count > maxFrequency {
			maxChar = char
			maxFrequency = count
		}
	}

	result = string(maxChar)
	return result
}
