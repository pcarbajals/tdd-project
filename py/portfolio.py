from typing import Dict

from money import Money


class Portfolio:
    def __init__(self) -> None:
        self.monies = []
        self.exchange_rates: Dict[str, float] = {
            "EUR->USD": 1.2,
            "USD->KRW": 1100,
        }

    def add(self, *monies) -> None:
        self.monies.extend(monies)

    def evaluate(self, currency: float) -> Money:
        total = sum(self.convert(money, currency) for money in self.monies)
        return Money(total, currency)

    def convert(self, money: Money, currency: str) -> float:
        rate = f"{money.currency}->{currency}"

        if money.currency == currency:
            return money.amount

        return money.amount * self.exchange_rates[rate]
