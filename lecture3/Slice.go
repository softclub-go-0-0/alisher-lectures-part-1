package main

import "fmt"

func main() {

	var n int
	fmt.Scan(&n)
	oddNumbers := make([]int, n)
	for i := 0; i < n; i++ {
		oddNumbers[i] = i*2 + 1
	}
	fmt.Println(oddNumbers)
}
