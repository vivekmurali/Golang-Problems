//https://codeforces.com/problemset/problem/959/A
package main

import (
	"fmt"
)

func main() {
	a := 0
	fmt.Scanf("%d", &a)
	if a%2 == 0 {
		fmt.Println("Mahmoud")
	} else {
		fmt.Println("Ehab")
	}
}
