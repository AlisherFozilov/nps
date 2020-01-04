package main

import "fmt"

func main() {
	score1 := 10
	//score2 := 10
	//score3 := 10
	promoters := 0
	detractorov := 0
	if score1 >= 0 {
		promoters = promoters + 1
	}

	nps := (promoters-detractorov) / 3*100
	fmt.Println(nps)
}