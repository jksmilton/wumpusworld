package main

import "math/rand"

type randAgent struct {
	location coordinate
	haveGold bool
}

func (agent *randAgent) DoAction() int {
	return rand.Intn(10)
}
