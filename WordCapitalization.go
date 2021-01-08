// Codeforces Word Capitaliztion 281/A
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()
	fmt.Println(strings.Title(str))
}
