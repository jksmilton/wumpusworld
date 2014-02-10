package main

import "fmt"

func main() {
	world := BuildWorld(coordinate{4, 4})
	drawWorld(world)
}

func drawWorld(world [][]int) {
	y := len(world)
	x := len(world[1])

	fmt.Printf("\n")

	for i := -1; i <= y; i++ {
		fmt.Printf("W")

		if i == -1 || i == y {
			for j := 0; j < x; j++ {
				fmt.Printf("W")
			}
		} else {
			for j := 0; j < x; j++ {
				fmt.Print(pickLetter(world[y-(i+1)][j]))
			}
		}

		fmt.Printf("W\n")

	}

	fmt.Println("Press enter for next turn")
}

func pickLetter(sym int) string {
	if sym == SymWall {
		return "W"
	} else if sym == SymPlayer {
		return "P"
	} else if sym == SymMon {
		return "M"
	} else if sym == SymGold {
		return "G"
	} else if sym == SymHole {
		return "H"
	} else {
		return " "
	}
}
