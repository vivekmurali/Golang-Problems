package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var tt int
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	fmt.Fscanf(r, "%d\n", &tt)

	for tt > 0 {
		var a, b int
		fmt.Fscanf(r, "%d %d\n", &a, &b)
		if a <= b {
			if 2*a <= b {
				fmt.Fprintln(w, b*b)
			} else {
				fmt.Fprintln(w, 4*a*a)
			}
		} else {
			if 2*b <= a {
				fmt.Fprintln(w, a*a)
			} else {
				fmt.Fprintln(w, 4*b*b)
			}
		}
		tt--
	}
}
