package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscanf(r, "%d\n", &n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscanf(r, "%d ", &arr[i])
	}
	fmt.Fscanf(r, "\n")
	solve(arr)
}

func solve(arr []int) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	e := 0
	o := 0
	for _, v := range arr {
		if v%2 == 0 {
			e++
		} else {
			o++
		}
	}
	if e == len(arr) || o == len(arr) {
		for _, v := range arr {
			fmt.Fprintf(w, "%d ", v)
		}
		return

	}
	sort.Ints(arr)
	for _, v := range arr {
		fmt.Fprintf(w, "%d ", v)
	}
}
