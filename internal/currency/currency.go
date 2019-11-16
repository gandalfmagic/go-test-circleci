package currency

type CoinType uint64

const (
	Copper   CoinType = 1
	Silver   CoinType = 10
	Electrum CoinType = 50
	Gold     CoinType = 100
	Platinum CoinType = 1000
)

var (
	allowedCoinTypes     = []CoinType{Copper, Silver, Electrum, Gold, Platinum}
	stringifiedCoinTypes = map[CoinType]string{
		Copper:   "Copper",
		Silver:   "Silver",
		Electrum: "Electrum",
		Gold:     "Gold",
		Platinum: "Platinum",
	}
)

func isCoinTypeAllowed(t CoinType) bool {
	for _, c := range allowedCoinTypes {
		if t == c {
			return true
		}
	}
	return false
}
