package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	arrStr := strings.Split(scanner.Text(), " ")
	arr := make([]int, n)
	for i, v := range arrStr {
		k, _ := strconv.Atoi(v)
		arr[i] = k
	}
	var (
		x,
		y,
		a,
		b int
	)
	for i, v := range arr {
		if v%2 == 0 {
			x = i
			a++
		} else {
			y = i
			b++
		}
	}
	if a == 1 {
		fmt.Println(x + 1)
	} else {
		fmt.Println(y + 1)
	}

}
