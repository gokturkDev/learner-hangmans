package coin

type Coin struct {
	heads, tails bool
}

func New(isHeads bool) Coin {
	if isHeads {
		return Coin{heads: true}
	}
	return Coin{tails: true}
}

func (coin *Coin) isHeads() bool {
	return coin.heads
}
