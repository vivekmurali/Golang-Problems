package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var tt int
	fmt.Fscanf(r, "%d\n", &tt)
	for tt > 0 {
		var a, b, c, p int64
		fmt.Fscanf(r, "%d %d %d %d\n", &p, &a, &b, &c)
		solve(p, a, b, c)
		tt--
	}

}

func solve(p, a, b, c int64) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	var A, B, C int64
	if p%a != 0 {
		A = a - (p % a)
	}
	if p%b != 0 {
		B = b - (p % b)
	}
	if p%c != 0 {
		C = c - (p % c)
	}
	// fmt.Fprintf(w, "hello")
	fmt.Fprintf(w, "%d\n", min(A, B, C))
	// fmt.Printf("%d\n", min(A, B, C))
}

func min(a ...int64) int64 {
	m := a[0]
	for _, v := range a {
		if v < m {
			m = v
		}
	}
	return m
}
