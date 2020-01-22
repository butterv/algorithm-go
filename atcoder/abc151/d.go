package main

import "fmt"

func main() {
	var H, W int
	fmt.Scanf("%d %d", &H, &W)
	field := make([][]byte, H)
	for i := 0; i < H; i++ {
		var s string
		fmt.Scanf("%s", &s)
		field[i] = []byte(s)
	}
	d := make([][]int, H*W)
	for i := 0; i < H*W; i++ {
		d[i] = make([]int, H*W)
		for j := 0; j < H*W; j++ {
			d[i][j] = 1000
		}
	}
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			src := w + W*h
			if field[h][w] == '#' {
				continue
			}
			d[src][src] = 0
			if h-1 >= 0 && field[h-1][w] == '.' {
				dst := w + W*(h-1)
				d[src][dst] = 1
			}
			if h+1 < H && field[h+1][w] == '.' {
				dst := w + W*(h+1)
				d[src][dst] = 1
			}
			if w-1 >= 0 && field[h][w-1] == '.' {
				dst := w - 1 + W*h
				d[src][dst] = 1
			}
			if w+1 < W && field[h][w+1] == '.' {
				dst := w + 1 + W*h
				d[src][dst] = 1
			}
		}
	}
	for k := 0; k < W*H; k++ {
		for i := 0; i < W*H; i++ {
			for j := 0; j < W*H; j++ {
				if d[i][j] > d[i][k]+d[k][j] {
					d[i][j] = d[i][k] + d[k][j]
				}
			}
		}
	}
	max := -1
	for i := 0; i < W*H; i++ {
		for j := 0; j < i; j++ {
			if d[i][j] > max {
				if d[i][j] != 1000 {
					max = d[i][j]
				}
			}
		}
	}
	fmt.Println(max)
}
