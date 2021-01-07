package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	arr := strings.Split(scanner.Text(), "+")
	sort.Strings(arr)
	fmt.Println(strings.Join(arr, "+"))
}
