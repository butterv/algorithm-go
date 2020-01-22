package main

import "fmt"

func main() {
	var N, K, M int
	fmt.Scanf("%d %d %d", &N, &K, &M)

	var total int
	for i := 0; i < N-1; i++ {
		var A int
		fmt.Scanf("%d", &A)
		total += A
	}
	must := N*M - total

	if must <= K {
		if must < 0 {
			must = 0
		}
		fmt.Println(must)
	} else {
		fmt.Println(-1)
	}
}
