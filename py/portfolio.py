from functools import reduce
from operator import add

from money import Money


class Portfolio:
    def __init__(self) -> None:
        self.monies = []

    def add(self, *monies) -> None:
        self.monies.extend(monies)

    def evaluate(self, currency) -> Money:
        total = reduce(add, [m.amount for m in self.monies], 0)
        return Money(total, currency) 
