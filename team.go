package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var count int
var scanner = bufio.NewScanner(os.Stdin)

func main() {
	var n int
	scanner.Scan()
	n, _ = strconv.Atoi(scanner.Text())
	for i := 0; i < n; i++ {
		solve()
	}
	fmt.Print(count, "\n")
}

func solve() {
	scanner.Scan()
	if strings.Count(scanner.Text(), "1") > 1 {
		count++
	}
}
