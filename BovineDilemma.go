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
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())
	for t > 0 {
		scanner.Scan()
		n, _ := strconv.Atoi(scanner.Text())
		arr := make([]int, n)
		scanner.Scan()
		str := strings.Split(scanner.Text(), " ")
		for i, v := range str {
			k, _ := strconv.Atoi(v)
			// fmt.Println("Debug: ", k)
			arr[i] = k
		}
		fmt.Println(solve(arr, n))
		t--
	}
}

func solve(arr []int, n int) int {
	areas := make([]float64, 0)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			area := calcArea(arr[i], arr[j])
			if !check(area, areas) {
				areas = append(areas, area)
			}
		}
	}
	return len(areas)
}

func check(area float64, areas []float64) bool {
	for _, v := range areas {
		if area == v {
			return true
		}
	}
	return false
}

func calcArea(a, b int) float64 {
	// a = float64(a)
	// b = float64(b)
	base := math.Abs(float64(a) - float64(b))
	return 0.5 * base
}
