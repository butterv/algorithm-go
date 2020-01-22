package main

import "fmt"

type Cnt struct {
	AC int
	WA int
}

func main() {
	var N, M int
	fmt.Scanf("%d %d", &N, &M)

	m := make(map[int]Cnt)
	for i := 0; i < M; i++ {
		var p int
		var S string
		fmt.Scanf("%d %s", &p, &S)

		v, ok := m[p]
		if !ok {
			v = Cnt{}
		}
		if S == "AC" {
			v.AC++
		}
		if S == "WA" && v.AC == 0 {
			v.WA++
		}
		m[p] = v
	}

	var acCnt, waCnt int
	for _, v := range m {
		if v.AC > 0 {
			acCnt++
			waCnt += v.WA
		}
	}

	fmt.Printf("%d %d", acCnt, waCnt)
}
