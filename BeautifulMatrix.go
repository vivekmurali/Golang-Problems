package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

var scanner = bufio.NewScanner(os.Stdin)

func main() {
	flag := []int{0, 0}
	var mat [5][5]int
	for i, v := range mat {
		scanner.Scan()
		str := strings.Split(scanner.Text(), " ")
		for j, _ := range v {
			val, _ := strconv.Atoi(str[j])
			if val == 1 {
				flag[0] = i
				flag[1] = j

			}
			mat[i][j] = val
		}
	}
	fmt.Println(math.Abs(2-float64(flag[0])) + math.Abs(2-float64(flag[1])))
}
