package currency

import (
	"testing"
)

func TestPouch_Add(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		p := NewPouch()
		p.Add(Copper, 1)
		p.Add(Copper, 2)
		p.Add(Silver, 20)

		if got := p.coins[Copper]; got != 3 {
			t.Errorf("Copper Coins() = %v, want %v", got, 3)
		}
		if got := p.coins[Silver]; got != 20 {
			t.Errorf("Silver Coins() = %v, want %v", got, 20)
		}

		if got := p.ContentValue(); got != 203 {
			t.Errorf("ContentValue() = %v, want %v", got, 203)
		}

		p.Remove(Copper, 1)

		if got := p.coins[Copper]; got != 2 {
			t.Errorf("Copper Coins() = %v, want %v", got, 3)
		}
		if got := p.coins[Silver]; got != 20 {
			t.Errorf("Silver Coins() = %v, want %v", got, 20)
		}

		if got := p.ContentValue(); got != 202 {
			t.Errorf("ContentValue() = %v, want %v", got, 203)
		}

		p.Remove(Silver, 5)

		if got := p.coins[Copper]; got != 2 {
			t.Errorf("Copper Coins() = %v, want %v", got, 3)
		}
		if got := p.coins[Silver]; got != 15 {
			t.Errorf("Silver Coins() = %v, want %v", got, 15)
		}

		if got := p.ContentValue(); got != 152 {
			t.Errorf("ContentValue() = %v, want %v", got, 153)
		}

	})
}
