package main

import "fmt"

func main() {
	input := 5
	fmt.Scan(&input)
	if (input%2) < 1 && input > 2 {
		fmt.Print("YES")
	} else {
		fmt.Print("NO")
	}
}
