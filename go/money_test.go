package main

import (
	"testing"
)

func TestMultiplicationInUSD(t *testing.T) {
	actualMoney := Money{5, "USD"}.Times(2)
	expectedMoney := Money{10, "USD"}

	assertEqual(actualMoney, expectedMoney, t)
}

func TestMultiplicationInEuros(t *testing.T) {
	actualMoney := Money{10, "EUR"}.Times(2)
	expectedMoney := Money{20, "EUR"}

	assertEqual(actualMoney, expectedMoney, t)
}

func TestDivision(t *testing.T) {
	actualMoney := Money{4002, "KRW"}.Divide(4)
	expectedMoney := Money{1000.5, "KRW"}

	assertEqual(actualMoney, expectedMoney, t)
}

func TestAddition(t *testing.T) {
	var portfolio Portfolio
	var portfolioInDollars Money

	fiveDollars := Money{5, "USD"}
	tenDollars := Money{10, "USD"}
	fifteenDollars := Money{15, "USD"}

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars = portfolio.Evaluate("USD")

	assertEqual(fifteenDollars, portfolioInDollars, t)
}

func assertEqual(actualMoney Money, expectedMoney Money, t *testing.T) {
	if actualMoney != expectedMoney {
		t.Errorf("Expected %+v, got %+v", expectedMoney, actualMoney)
	}
}
