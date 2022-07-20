package main

import (
	stocks "tdd/stocks"
	"testing"
)

func TestMultiplication(t *testing.T) {
	actualMoney := stocks.NewMoney(10, "EUR").Times(2)
	expectedMoney := stocks.NewMoney(20, "EUR")

	assertEqual(t, expectedMoney, actualMoney)
}

func TestDivision(t *testing.T) {
	actualMoney := stocks.NewMoney(4002, "KRW").Divide(4)
	expectedMoney := stocks.NewMoney(1000.5, "KRW")

	assertEqual(t, expectedMoney, actualMoney)
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

	assertEqual(t, portfolioInDollars, fifteenDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)
	actualMoney, _ := portfolio.Evaluate("USD")

	expectedMoney := stocks.NewMoney(17, "USD")

	assertEqual(t, expectedMoney, actualMoney)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundredWon := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)
	actualMoney, _ := portfolio.Evaluate("KRW")

	expectedMoney := stocks.NewMoney(2200, "KRW")

	assertEqual(t, expectedMoney, actualMoney)
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

	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}
