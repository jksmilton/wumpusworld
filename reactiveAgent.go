package main 

import "math/rand"

type reactiveAgent struct {
	location coordinate
	hasgold bool
}

func (agent *reactiveAgent) DoAction(percepts Senses) int {
	if percepts[Glimmer] {
		return PickUpMask
	} else if percepts[Smell] {
		direction := rand.Intn(4) + 1
		return ShootMask + direction
	} else {
		direction := rand.Intn(4)
		return MovMask + direction + 1
	}
}