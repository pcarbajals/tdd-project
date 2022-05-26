from functools import reduce
from lib2to3.pytree import convert
from operator import add

from money import Money


class Portfolio:
    def __init__(self) -> None:
        self.monies = []
        self.eur_to_usd_rate: float = 1.2

    def add(self, *monies) -> None:
        self.monies.extend(monies)

    def evaluate(self, currency: float) -> Money:
        total = sum(self.convert(money, currency) for money in self.monies)
        return Money(total, currency) 

    def convert(self, money: Money, currency: str) -> float:
        return money.amount if money.currency == currency else money.amount * self.eur_to_usd_rate
