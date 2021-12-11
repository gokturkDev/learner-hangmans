package player

import (
	"math/rand"
	"time"
)

type Playerer interface {
	FlipCoin()
}

type Player struct {
	money int
}

type Coin struct {
	heads, tails bool
}

func (ply Player) FlipCoin(amount int) Coin {
	rand.NewSource(time.Now().Unix())
	randNum := rand.Intn(1)
	ply.money -= amount
	if randNum < 1 {
		return Coin{heads: true}
	}
	return Coin{tails: true}
}
