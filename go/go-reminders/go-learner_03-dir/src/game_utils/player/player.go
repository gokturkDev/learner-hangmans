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
	ply.wins += 1
}

func (ply *Player) Lost(amount int) { //this is how to create a method, very interesting architecture
	ply.money -= amount
	ply.loss += 1
}

func (ply *Player) GetMoney() int {
	return ply.money
}

func (ply *Player) GetWinCount() int {
	return ply.wins
}

func (ply *Player) GetLossCount() int {
	return ply.loss
}

func New() Player {
	return Player{
		money: 1000,
		wins:  0,
		loss:  0,
	}
}
