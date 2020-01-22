package main

import "fmt"

func main() {
	var C int
	fmt.Scanf("%c", &C)
	fmt.Printf("%c\n", rune(C+1))
}
