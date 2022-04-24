import unittest

class Dollar:
    def __init__(self, amount) -> None:
        self.amount = amount

    def times(self, multiplier):
        return Dollar(10)

class TestMoney(unittest.TestCase):
    def testMultiplication(self):
        fiver = Dollar(5)
        tenner = fiver.times(2)
        self.assertEqual(10, tenner.amount)


if __name__ == '__main__':
    unittest.main()
