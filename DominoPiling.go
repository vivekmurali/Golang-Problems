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
	nums := strings.Split(scanner.Text(), " ")
	num1, _ := strconv.Atoi(nums[0])
	num2, _ := strconv.Atoi(nums[1])
	fmt.Println(num2 * num1 / 2)
}
