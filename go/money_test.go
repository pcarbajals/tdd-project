package main

import (
	"testing"
)

func TestMultiplicationInUSD(t *testing.T) {
	fiver := Money{amount: 5, currency: "USD"}
	actualMoney := fiver.Times(2)
	expectedMoney := Money{amount: 10, currency: "USD"}

	if actualMoney != expectedMoney {
		t.Errorf("Expected %+v, got %+v", expectedMoney, actualMoney)
	}
}

func TestMultiplicationInEuros(t *testing.T) {
	tenEuros := Money{amount: 10, currency: "EUR"}
	actualMoney := tenEuros.Times(2)
	expectedMoney := Money{amount: 20, currency: "EUR"}

	if actualMoney != expectedMoney {
		t.Errorf("Expected %+v, got %+v", expectedMoney, actualMoney)
	}
}

func TestDivision(t *testing.T) {
	money := Money{amount: 4002, currency: "KRW"}
	actualMoneyAfterDivision := money.Divide(4)
	expectedMoneyAfterDivision := Money{amount: 1000.5, currency: "KRW"}

	if actualMoneyAfterDivision != expectedMoneyAfterDivision {
		t.Errorf("Expected %+v, got %+v", expectedMoneyAfterDivision, actualMoneyAfterDivision)
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
