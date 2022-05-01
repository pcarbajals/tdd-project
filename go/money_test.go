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

type Money struct {
	amount   float64
	currency string
}

func (m Money) Times(multiplier int) Money {
	return Money{amount: m.amount * float64(multiplier), currency: m.currency}
}

func (m Money) Divide(divisor int) Money {
	return Money{amount: m.amount / float64(divisor), currency: m.currency}
}

type Portfolio []Money

func (p Portfolio) Add(money Money) Portfolio {
	return append(p, money)
}

func (p Portfolio) Evaluate(currency string) Money {
	total := 0.0
	for _, money := range p {
		total = total + money.amount
	}
	return Money{total, currency}
}
