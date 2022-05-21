package stocks

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total = total + convert(money, currency)
	}
	return Money{total, currency}
}

func convert(money Money, currency string) float64 {
	eurToUsd := 1.2

	if money.currency == currency {
		return money.amount
	}

	return money.amount * eurToUsd
}
