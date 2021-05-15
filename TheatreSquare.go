package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	r := bufio.NewReader(os.Stdin)
	var n, m, a int
	fmt.Fscanf(r, "%d %d %d\n", &n, &m, &a)
	if (n*m)%(a*a) == 0 {
		fmt.Println("Reached ehre")
		fmt.Fprintln(w, (n*m)/(a*a))
		return
	}
	fmt.Fprintln(w, (n*m)%(a*a))
}
