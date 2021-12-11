package coin

import (
	"math/rand"
	"time"
)

type Coin struct {
	Heads, Tails bool
}

func flipACoin() Coin {
	rand.NewSource(time.Now().Unix())
	randNum := rand.Intn(1)
	if randNum < 1 {
		return Coin{Heads: true}
	}
	return Coin{Tails: true}
}
