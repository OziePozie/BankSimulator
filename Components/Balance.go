package Components

type Balance struct {
	Dollars float64
	Euros   float64
	Rubles  float64
}

const (
	DOLLARS_TO_RUBLES = 0.1
	EUROS_TO_RUBLES   = 0.15
	DOLLARS_TO_EUROS  = 1.1
)

func (balance Balance) getBalance() map[string]float64 {

	var balances map[string]float64

	balances["rubles"] = balance.Rubles
	balances["dollars"] = balance.Dollars
	balances["euros"] = balance.Euros

	return balances
}

func (balance Balance) convert(sum float64, switcher int) float64 {
	var res float64
	switch switcher {
	case 1:
		res = sum * DOLLARS_TO_EUROS
	case 2:
		res = sum * DOLLARS_TO_RUBLES
	case 3:
		res = sum * EUROS_TO_RUBLES

	}

	return res
}
