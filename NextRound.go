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
	first := strings.Split(scanner.Text(), " ")
	n, _ := strconv.Atoi(first[0])
	k, _ := strconv.Atoi(first[1])
	scanner.Scan()
	second := strings.Split(scanner.Text(), " ")
	secondSlice := make([]int, 0)
	for _, v := range second {
		n, _ := strconv.Atoi(v)
		secondSlice = append(secondSlice, n)
	}
	score := secondSlice[k-1]
	count := 0
	for i := 0; i < n; i++ {
		if secondSlice[i] >= score && secondSlice[i] > 0 {
			count++
		}
	}
	fmt.Println(count)
}
