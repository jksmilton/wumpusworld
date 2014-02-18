package main

//World Representation
const SymWall int = 1
const SymMon int = 2
const SymHole int = 3
const SymGold int = 4
const SymEmpty int = 0
const SymPlayer int = 5

//Action Representation
const MovMask = 0
const ShootMask = 4
const StillMask = 0
const LeftMask = 1
const UpMask = 2
const RightMask = 3
const DownMask = 4
const PickUpMask = 9

//Percept input type
type Senses []bool

//Percept Representation bit
const Bump = 1
const Wind = 2
const Smell = 3
const Glimmer = 4
const Scream = 5
const GotGold = 6
const Gobble = 7
const Fall = 8
