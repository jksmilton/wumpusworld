package main

import "fmt"

func main() {
	theWorld := BuildWorld(coordinate{10, 10})
	drawWorld(theWorld)
	percepts := theWorld.DoAction(0)
	fmt.Printf("%v\n", percepts)
	fmt.Println("Press enter for next turn")
}

func drawWorld(theWorld world) {
	y := len(theWorld.worldMap)
	x := len(theWorld.worldMap[1])

	fmt.Printf("\n")

	for i := -1; i <= y; i++ {
		fmt.Printf("W")

		if i == -1 || i == y {
			for j := 0; j < x; j++ {
				fmt.Printf("W")
			}
		} else {
			for j := 0; j < x; j++ {
				crntLoc := coordinate{j, y - (i + 1)}
				if theWorld.playerLoc == crntLoc {
					fmt.Print(pickLetter(SymPlayer))
				} else {
					fmt.Print(pickLetter(theWorld.worldMap[y-(i+1)][j]))
				}
			}
		}

		fmt.Printf("W\n")

	}

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
