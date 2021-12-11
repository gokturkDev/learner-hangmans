package coin

type Coin struct {
	heads, tails bool
}

func NewHeadsCoin() Coin {
	return Coin{heads: true}
}

func NewTailsCoin() Coin {
	return Coin{tails: true}
}

func (coin *Coin) IsHeads() bool {
	return coin.heads
}
