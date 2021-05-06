// https://codeforces.com/contest/1520/problem/A
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var t int
	// fmt.Scanf("%d\n", &t)
	fmt.Fscanf(r, "%d\n", &t)
	for t > 0 {
		var length int
		// fmt.Scanf("%d\n", &length)
		fmt.Fscanf(r, "%d\n", &length)
		var word string
		fmt.Fscanf(r, "%s\n", &word)
		solve(length, word)
		t--
	}
}

func solve(l int, w string) {
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()
	s := strings.Split(w, "")
	sl := make([]string, 0)
	for i := 0; i < len(s); i++ {
		if i == len(s)-1 {
			if ok := strings.Contains(strings.Join(sl, ""), s[i]); ok {
				fmt.Fprintln(wr, "NO")
				return
			}
			break
		}
		// fmt.Println(s[i])
		if s[i] == s[i+1] {
			continue
		} else {
			if ok := strings.Contains(strings.Join(sl, ""), s[i]); ok {
				fmt.Fprintln(wr, "NO")
				return
			}
			sl = append(sl, s[i])
		}
	}
	fmt.Fprintln(wr, "YES")
}
