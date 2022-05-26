from functools import reduce
from lib2to3.pytree import convert
from operator import add

from money import Money


class Portfolio:
    def __init__(self) -> None:
        self.monies = []

    def add(self, *monies) -> None:
        self.monies.extend(monies)

    def evaluate(self, currency: float) -> Money:
        total = reduce(add, [self.convert(money, currency) for money in self.monies], 0)
        return Money(total, currency) 

    def convert(self, money: Money, currency: str) -> float:
        return money.amount if money.currency == currency else money.amount * 1.2
