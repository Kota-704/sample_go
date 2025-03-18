package main

import (
	"fmt"
	"math/rand/v2"
)

// ガチャの種類を選ぶ
func choose_gacha() (n int) {
	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		return 1
	case 2:
		return 11
	default:
		fmt.Println("1か2を選択してください")
		return choose_gacha()
	}
}

// ガチャを回す
func pull_gacha(n int) {
	for i := 0; i < n; i++ {
		x := rand.N(100)

		switch {
		case x < 80:
			fmt.Println("ノーマル")
		case x < 93:
			fmt.Println("R")
		case x < 98:
			fmt.Println("SR")
		default:
			fmt.Println("SSR")
		}
	}
}

func main() {
	n := choose_gacha()
	pull_gacha(n)
}
