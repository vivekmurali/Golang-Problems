package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var tt int
	fmt.Fscanf(r, "%d\n", &tt)
	for tt > 0 {
		var n int
		fmt.Fscanf(r, "%d\n", &n)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscanf(r, "%d ", &arr[i])
		}
		fmt.Fscanf(r, "\n")

		fmt.Println(solve(n, arr) + 1)
		tt--
	}
}

func solve(n int, arr []int) int {
	b := make([]int, len(arr))
	copy(b, arr)
	sort.Ints(b)
	if b[0] == b[1] {
		return (check(b[len(b)-1], arr))
	} else {
		return (check(b[0], arr))
	}
}

func check(n int, arr []int) int {
	for i, v := range arr {
		if v == n {
			return i
		}
	}
	return -1
}
