from audioop import mul
import unittest


class Money:
    def __init__(self, amount, currency) -> None:
        self.amount = amount
        self.currency = currency

    def times(self, multiplier):
        return Money(self.amount * multiplier, self.currency)

    def divide(self, divisor):
        return Money(self.amount / divisor, self.currency)


class TestMoney(unittest.TestCase):
    def test_multiplication_in_usd(self):
        actual_money = Money(5, "USD").times(2)
        expected_money = Money(10, "USD")

        self.assertEqual(actual_money.amount, expected_money.amount)
        self.assertEqual(actual_money.currency, expected_money.currency)

    def test_multiplication_in_euros(self):
        actual_money = Money(10, "EUR").times(2)
        expected_money = Money(20, "EUR")

        self.assertEqual(actual_money.amount, expected_money.amount)
        self.assertEqual(actual_money.currency, expected_money.currency)

    def test_division(self):
        actual_money = Money(4002, "KRW").divide(4)
        expected_money = Money(1000.5, "KRW")

        self.assertEqual(actual_money.amount, expected_money.amount)
        self.assertEqual(actual_money.currency, expected_money.currency)


if __name__ == '__main__':
    unittest.main()
