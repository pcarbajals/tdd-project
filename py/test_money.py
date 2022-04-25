from audioop import mul
import unittest


class Money:
    def __init__(self, amount, currency) -> None:
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)


class TestMoney(unittest.TestCase):
    def test_multiplication_in_usd(self):
        fiver = Money(5, "USD")
        tenner = fiver.times(2)
        self.assertEqual(tenner.amount, 10)
        self.assertEqual(tenner.currency, "USD")

    def test_multiplication_in_euros(self):
        tenEuros = Money(10, "EUR")
        twentyEuros = tenEuros.times(2)
        self.assertEqual(twentyEuros.amount, 20)
        self.assertEqual(twentyEuros.currency, "EUR")


if __name__ == '__main__':
    unittest.main()
