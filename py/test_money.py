import unittest

from money import Money
from portfolio import Portfolio


class TestMoney(unittest.TestCase):
    def test_multiplication(self):
        actual_money = Money(10, "EUR").times(2)
        expected_money = Money(20, "EUR")

        self.assertEqual(actual_money, expected_money)

    def test_division(self):
        actual_money = Money(4002, "KRW").divide(4)
        expected_money = Money(1000.5, "KRW")

        self.assertEqual(actual_money, expected_money)


class TestPortfolio(unittest.TestCase):
    def test_addition_in_usd(self):
        five_dollars = Money(5, "USD")
        ten_dollars = Money(10, "USD")
        portfolio = Portfolio()
        portfolio.add(five_dollars, ten_dollars)

        actual_money = portfolio.evaluate("USD")
        expected_money = Money(15, "USD")

        self.assertEqual(actual_money, expected_money)

    def test_addition_of_usd_and_eur(self):
        five_dollars = Money(5, "USD")
        ten_euros = Money(10, "EUR")
        portfolio = Portfolio()
        portfolio.add(five_dollars, ten_euros)

        actual_money = portfolio.evaluate("USD")
        expected_money = Money(17, "USD")

        self.assertEqual(actual_money, expected_money)

    def test_addition_of_usd_and_won(self):
        one_dollar = Money(1, "USD")
        eleven_hundred_wons = Money(1100, "KRW")
        portfolio = Portfolio()
        portfolio.add(one_dollar, eleven_hundred_wons)

        actual_money = portfolio.evaluate("KRW")
        expected_money = Money(2200, "KRW")

        self.assertEqual(actual_money, expected_money)


if __name__ == '__main__':
    unittest.main()
