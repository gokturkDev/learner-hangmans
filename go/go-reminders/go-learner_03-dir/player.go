package player

type Playerer interface {
	FlipCoin()
}

type Player struct {
	money int
}

type Coin struct {
	heads, tails bool
}

}
