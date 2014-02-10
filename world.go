package main

import "math/rand"

const propWall int = 10
const propMon int = 10 + propWall
const propHole int = 10 + propMon

func BuildWorld(size coordinate) [][]int {
	world := make([][]int, size.y)

	for i := 0; i < size.y; i++ {
		world[i] = make([]int, size.x)

		for j := 0; j < size.x; j++ {
			picker := rand.Intn(100)
			if picker < propWall {
				world[i][j] = SymWall
			} else if picker < propMon {
				world[i][j] = SymMon
			} else if picker < propHole {
				world[i][j] = SymHole
			} else {
				world[i][j] = SymEmpty
			}
		}
	}
	world[0][0] = SymPlayer
	goldPlaced := false

	for !goldPlaced {
		x := rand.Intn(size.x)
		y := rand.Intn(size.y)
		if world[y][x] == 0 {
			world[y][x] = SymGold
			goldPlaced = true
		}
	}

	return world
}
