package main

import (
	"fmt"
	"math/rand/v2"
)

var n int

func choose_gacha() {
	fmt.Println("1: 単発ガチャ 2: 11連ガチャ")
	fmt.Scanln(&n)

	switch n {
	case 1:
		n = 1

	case 2:
		n = 11

	default:
		fmt.Println("1か2")
	}
}


	func pull_gacha() {
		 i := 0
	LOOP:
		for {
			if i >= n {
				break LOOP
			}

			x :=	rand.N(100)
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
			i++
		}
	}


func main () {
	choose_gacha()
	pull_gacha()
}
