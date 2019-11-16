package currency

import "strconv"

type Pouch struct {
	coins map[CoinType]uint64
}

func NewPouch() *Pouch {
	coins := make(map[CoinType]uint64)
	return &Pouch{coins: coins}
}

func (p Pouch) Add(coinType CoinType, quantity uint64) {
	if !isCoinTypeAllowed(coinType) {
		return
	}

	if _, ok := p.coins[coinType]; ok {
		p.coins[coinType] += quantity
	} else {
		p.coins[coinType] = quantity
	}
}

func (p Pouch) Remove(coinType CoinType, quantity uint64) bool {
	if !isCoinTypeAllowed(coinType) {
		return false
	}

	if _, ok := p.coins[coinType]; ok {
		current := p.coins[coinType]
		if current >= quantity {
			p.coins[coinType] -= quantity
			return true
		}
	}

	return false
}

func (p Pouch) ContentValue() uint64 {
	var total uint64

	for value, qty := range p.coins {
		total += uint64(value) * qty
	}

	return total
}

func (p Pouch) String() string {
	var s string

	for t, qty := range p.coins {
		s += stringifiedCoinTypes[t] + ": " + strconv.FormatUint(qty, 10) + ", "
	}

	if len(s) == 0 {
		return "This pouch is empty"
	}

	return s[:len(s)-2]
}
