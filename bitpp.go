package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	num := 0
	for n > 0 {
		scanner.Scan()
		str := scanner.Text()
		if strings.Contains(str, "++") {
			num += 1
		} else {
			num -= 1
		}
		n--
	}
	fmt.Println(num)
}
