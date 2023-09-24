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
	RUBLES_TO_DOLLAR  = 1.1
	RUBLES_TO_EUROS   = 1.15
	EUROS_TO_DOLLARS  = 0.9
)

func (balance Balance) getBalance() map[string]float64 {

	var balances map[string]float64

	balances["rubles"] = balance.Rubles
	balances["dollars"] = balance.Dollars
	balances["euros"] = balance.Euros

	return balances
}

func (balance Balance) convert(sum float64, currency string) float64 {
	var res float64
	switch currency {
	case "DOLTOEU":
		res = sum * DOLLARS_TO_EUROS
	case "DOLTORUB":
		res = sum * DOLLARS_TO_RUBLES
	case "EUTORU":
		res = sum * EUROS_TO_RUBLES
	case "EUTODOL":
		res = sum * EUROS_TO_DOLLARS
	case "RUTOEU":
		res = sum * RUBLES_TO_EUROS
	case "RUTODOL":
		res = sum * RUBLES_TO_DOLLAR

	}

	return res
}
