package main

import (
	stocks "tdd/stocks"
	"testing"
)

func TestMultiplication(t *testing.T) {
	actualMoney := stocks.NewMoney(10, "EUR").Times(2)
	expectedMoney := stocks.NewMoney(20, "EUR")

	assertEqual(actualMoney, expectedMoney, t)
}

func TestDivision(t *testing.T) {
	actualMoney := stocks.NewMoney(4002, "KRW").Divide(4)
	expectedMoney := stocks.NewMoney(1000.5, "KRW")

	assertEqual(actualMoney, expectedMoney, t)
}

func TestAddition(t *testing.T) {
	var portfolio stocks.Portfolio
	var portfolioInDollars stocks.Money

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fifteenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, _ = portfolio.Evaluate("USD")

	assertEqual(fifteenDollars, portfolioInDollars, t)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)
	actualMoney, _ := portfolio.Evaluate("USD")

	expectMoney := stocks.NewMoney(17, "USD")

	assertEqual(actualMoney, expectMoney, t)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundredWon := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)
	actualMoney, _ := portfolio.Evaluate("KRW")

	expectedMoney := stocks.NewMoney(2200, "KRW")

	assertEqual(actualMoney, expectedMoney, t)
}

func TestAddtionWithMultipleMissingExchangeRates(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	oneEuro := stocks.NewMoney(1, "EUR")
	oneWon := stocks.NewMoney(1, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(oneEuro)
	portfolio = portfolio.Add(oneWon)

	expectedErrorMessage :=
		"Missing exchange rate(s):[USD->Kalganid,EUR->Kalganid,KRW->Kalganid,]"
	_, actualError := portfolio.Evaluate("Kalganid")

	if expectedErrorMessage != actualError.Error() {
		t.Errorf("Expected %s Got %s",
			expectedErrorMessage, actualError.Error())
	}
}

func assertEqual(actualMoney stocks.Money, expectedMoney stocks.Money, t *testing.T) {
	if actualMoney != expectedMoney {
		t.Errorf("Expected %+v, got %+v", expectedMoney, actualMoney)
	}
}
