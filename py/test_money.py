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

    def __eq__(self, other: object) -> bool:
        return self.amount == other.amount and self.currency == other.currency


class TestMoney(unittest.TestCase):
    def test_multiplication_in_usd(self):
        actual_money = Money(5, "USD").times(2)
        expected_money = Money(10, "USD")

        self.assertEqual(actual_money, expected_money)

    def test_multiplication_in_euros(self):
        actual_money = Money(10, "EUR").times(2)
        expected_money = Money(20, "EUR")

        self.assertEqual(actual_money, expected_money)

    def test_division(self):
        actual_money = Money(4002, "KRW").divide(4)
        expected_money = Money(1000.5, "KRW")

        self.assertEqual(actual_money, expected_money)


if __name__ == '__main__':
    unittest.main()
