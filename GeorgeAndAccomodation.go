// https://codeforces.com/problemset/problem/467/A
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	count := 0
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	for n > 0 {
		scanner.Scan()
		str := strings.Split(scanner.Text(), " ")
		n1, _ := strconv.Atoi(str[0])
		n2, _ := strconv.Atoi(str[1])
		if n2-n1 >= 2 {
			count += 1
		}
		n--
	}

	fmt.Println(count)

}
