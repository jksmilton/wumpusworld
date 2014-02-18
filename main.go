package main

import "fmt"
import "os"
import "strconv"
import "bufio"

func main() {
	dimensions := getDimensions()
	action := 0

	scanner := bufio.NewScanner(os.Stdin)

	theWorld := BuildWorld(dimensions)
	drawWorld(theWorld)
	percepts := theWorld.DoAction(action)
	fmt.Printf("%v\n", percepts)
	writeAction(action, percepts)
	fmt.Println("Press enter for next turn")
	scanner.Scan()

	continueRun := true

	for continueRun {
		action = DoAction()
		percepts = theWorld.DoAction(action)
		fmt.Printf("%v\n", percepts)
		drawWorld(theWorld)
		writeAction(action, percepts)
		fmt.Println("Press enter for next turn")
		scanner.Scan()

		if percepts[Gobble] || percepts[Fall] {
			fmt.Println("Game Over")
			continueRun = false
		} else if theWorld.hasGold && theWorld.playerLoc.x == 0 && theWorld.playerLoc.y == 0 {
			fmt.Println("Agent wins")
			continueRun = false
		}
	}

}

func getDimensions() coordinate {
	arguments := os.Args[1:]

	x := 4
	y := 4

	if len(arguments) == 1 {
		dimensionSize := parseArgument(arguments[0])
		x = dimensionSize
		y = dimensionSize
	} else if len(arguments) >= 2 {
		x = parseArgument(arguments[0])
		y = parseArgument(arguments[1])
	}

	return coordinate{x, y}

}

func parseArgument(arg string) int {
	i, err := strconv.Atoi(arg)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}
	return i
}

func writeAction(action int, percepts Senses) {
	if action == PickUpMask {
		fmt.Println("Attempted to pick up gold")
	} else if action > ShootMask {
		fmt.Print("Attempted to shoot ")
		switch action - ShootMask {
		case LeftMask:
			fmt.Printf("left\n")
		case RightMask:
			fmt.Printf("right\n")
		case UpMask:
			fmt.Printf("up\n")
		case DownMask:
			fmt.Printf("down\n")
		default:
			fmt.Printf("nowhere\n")
		}
	} else {
		fmt.Print("Attempted to move ")
		switch action {
		case LeftMask:
			fmt.Printf("left\n")
		case RightMask:
			fmt.Printf("right\n")
		case UpMask:
			fmt.Printf("up\n")
		case DownMask:
			fmt.Printf("down\n")
		default:
			fmt.Printf("nowhere\n")
		}
	}

	if percepts[Bump] {
		fmt.Println("Bumped into a wall")
	}
	if percepts[Fall] {
		fmt.Println("Fell down a hole")
	}
	if percepts[Smell] {
		fmt.Println("Smelled a nearby wumpus")
	}
	if percepts[Wind] {
		fmt.Println("Felt a breeze from a nearby hole")
	}
	if percepts[Glimmer] {
		fmt.Println("Saw a glimmer of gold")
	}
	if percepts[Gobble] {
		fmt.Println("Got eaten by a wumpus")
	}
	if percepts[Scream] {
		fmt.Println("Heard a scream")
	}
	if percepts[GotGold] {
		fmt.Println("Is carrying the gold")
	}

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
