package main

import "fmt"

func main() {
	word := ""
	n := 0
	fmt.Scan(&n)
	for n > 0 {
		fmt.Scan(&word)
		length := len(word)
		if length > 10 {
			fmt.Printf("%s%d%s\n", word[0:1], length-2, word[length-1:])
		} else {
			fmt.Print(word, "\n")
		}
		n--
	}
}
