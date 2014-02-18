package main

import "math/rand"
import "fmt"

const propWall int = 10
const propMon int = 10 + propWall
const propHole int = 10 + propMon

type world struct {
	worldMap  [][]int
	playerLoc coordinate
	hasGold   bool
}

func (thisWorld *world) handleShoot(direction int, crntSenses Senses) Senses {
	return crntSenses
}

func (thisWorld *world) handlePickup() {
	if thisWorld.worldMap[thisWorld.playerLoc.y][thisWorld.playerLoc.x] == SymGold {
		thisWorld.hasGold = true
		thisWorld.worldMap[thisWorld.playerLoc.y][thisWorld.playerLoc.x] = SymEmpty
	}

}

func (thisWorld *world) moveTo(loc coordinate, crntSenses Senses) Senses {
	y := len(thisWorld.worldMap)
	x := len(thisWorld.worldMap[0])

	if loc.x >= x || loc.x < 0 || loc.y >= y || loc.y < 0 {
		fmt.Println("Bumped into a wall")
		crntSenses[Bump] = true
	} else if thisWorld.worldMap[loc.y][loc.x] == SymWall {
		fmt.Println("Bumped into a wall")
		crntSenses[Bump] = true
	} else {
		thisWorld.playerLoc = loc
	}
	return crntSenses
}

func (thisWorld *world) handleMov(direction int, crntSenses Senses) Senses {
	attemptedMov := coordinate{thisWorld.playerLoc.x, thisWorld.playerLoc.y}
	if direction == LeftMask {
		attemptedMov.x = attemptedMov.x - 1
		crntSenses = thisWorld.moveTo(attemptedMov, crntSenses)
	} else if direction == UpMask {
		attemptedMov.y += 1
		crntSenses = thisWorld.moveTo(attemptedMov, crntSenses)
	} else if direction == RightMask {
		attemptedMov.x += 1
		crntSenses = thisWorld.moveTo(attemptedMov, crntSenses)
	} else if direction == DownMask {
		attemptedMov.y = attemptedMov.y - 1
		crntSenses = thisWorld.moveTo(attemptedMov, crntSenses)
	}
	return crntSenses
}

func (thisWorld *world) DoAction(action int) Senses {
	crntSenses := make(Senses, 9)

	if action >= PickUpMask {
		thisWorld.handlePickup()
	} else if action >= ShootMask {
		crntSenses = thisWorld.handleShoot(action-ShootMask, crntSenses)
	} else {
		crntSenses = thisWorld.handleMov(action-MovMask, crntSenses)
	}
	return thisWorld.getOtherPercepts(crntSenses)
}

func (thisWorld *world) getOtherPercepts(crntSenses Senses) Senses {

	x := thisWorld.playerLoc.x
	y := thisWorld.playerLoc.y
	maxY := len(thisWorld.worldMap)
	maxX := len(thisWorld.worldMap[0])

	if thisWorld.worldMap[y][x] == SymGold {
		crntSenses[Glimmer] = true
	} else if thisWorld.worldMap[y][x] == SymMon {
		crntSenses[Gobble] = true
	} else if thisWorld.worldMap[y][x] == SymHole {
		crntSenses[Fall] = true
	}
	if thisWorld.hasGold {
		crntSenses[GotGold] = true
	}
	if x > 0 {
		crntSenses = thisWorld.checkSquare(crntSenses, coordinate{x - 1, y})
	}
	if y > 0 {
		crntSenses = thisWorld.checkSquare(crntSenses, coordinate{x, y - 1})
	}
	if x < maxX-1 {
		crntSenses = thisWorld.checkSquare(crntSenses, coordinate{x + 1, y})
	}
	if y < maxY-1 {
		crntSenses = thisWorld.checkSquare(crntSenses, coordinate{x, y + 1})
	}
	return crntSenses
}

func (thisWorld *world) checkSquare(crntSenses Senses, loc coordinate) Senses {
	if thisWorld.worldMap[loc.y][loc.x] == SymMon {
		crntSenses[Smell] = true
	} else if thisWorld.worldMap[loc.y][loc.x] == SymHole {
		crntSenses[Wind] = true
	}
	return crntSenses
}

func BuildWorld(size coordinate) world {
	newWorld := make([][]int, size.y)

	for i := 0; i < size.y; i++ {
		newWorld[i] = make([]int, size.x)

		for j := 0; j < size.x; j++ {
			picker := rand.Intn(100)
			if picker < propWall {
				newWorld[i][j] = SymWall
			} else if picker < propMon {
				newWorld[i][j] = SymMon
			} else if picker < propHole {
				newWorld[i][j] = SymHole
			} else {
				newWorld[i][j] = SymEmpty
			}
		}
	}
	newWorld[0][0] = SymEmpty
	goldPlaced := false

	for !goldPlaced {
		x := rand.Intn(size.x)
		y := rand.Intn(size.y)
		if newWorld[y][x] == 0 {
			newWorld[y][x] = SymGold
			goldPlaced = true
		}
	}

	return world{newWorld, coordinate{0, 0}, false}
}
