// 与えられた文字列でキャラクターを４方向に進めるプログラム
package main

import (
	"fmt"
)

func characterLocation(commands string) []int {
	ans := []int{0, 0}
	for _, v := range commands {
		// string()でキャスト
		if string(v) == "N" || string(v) == "S" {
			ans[1] += whatDirection(string(v))
		} else {
			ans[0] += whatDirection(string(v))
		}
	}
	return ans
}

func whatDirection(direction string) int {
	if direction == "N" || direction == "E" {
		return 1
	} else if direction == "S" || direction == "W" {
		return -1
	} else {
		return 0
	}
}

func main() {
	fmt.Println(characterLocation("NNNN"))   // [0 4]
	fmt.Println(characterLocation("NESW"))   // [0 0]
	fmt.Println(characterLocation("NW"))     // [-1 1]
	fmt.Println(characterLocation("AWFGMD")) // [-1 0]
}
