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

func (ply *Player) Won(amount int) {
	ply.money += amount
}

func (ply *Player) Lost(amount int) { //this is how to create a method, very interesting architecture
	ply.money -= amount
}

func New() Player {
	return Player{
		money: 1000,
		wins:  0,
		loss:  0,
	}
}
