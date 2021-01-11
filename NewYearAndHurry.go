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
	str := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(str[0])
	k, _ := strconv.Atoi(str[1])
	time := 240 - k
	count := 0
	sumTime := 0
	for i := 1; i <= n; i++ {
		sumTime += 5 * i
		if sumTime > time {
			break
		}
		count++
	}

	fmt.Println(count)
}
