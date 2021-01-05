package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	scanner.Scan()
	str1 := strings.ToLower(scanner.Text())
	scanner.Scan()
	str2 := strings.ToLower(scanner.Text())
	fmt.Println(strings.Compare(str1, str2))
}
