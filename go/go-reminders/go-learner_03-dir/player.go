package player

type Playerer interface {
	Won()
	Loss()
}

type Player struct {
	money int
	wins  int
	loss  int
}

type Coin struct {
	heads, tails bool
}

func (ply Player) Won(amount int) {
	ply.money += amount
}

func (ply Player) Lost(amount int) { //this is how to create a method, very interesting architecture
	ply.money -= amount
}
