package main

import (
	"reflect"
	stocks "tdd/stocks"
	"testing"
)

var bank stocks.Bank

func init() {
	bank = stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	bank.AddExchangeRate("USD", "KRW", 1100)
}

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

	fiveDollars := stocks.NewMoney(5, "USD")
	tenDollars := stocks.NewMoney(10, "USD")
	fifteenDollars := stocks.NewMoney(15, "USD")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenDollars)
	portfolioInDollars, err := portfolio.Evaluate(bank, "USD")

	assertNil(t, err)
	assertEqual(t, fifteenDollars, *portfolioInDollars)
}

func TestAdditionOfDollarsAndEuros(t *testing.T) {
	var portfolio stocks.Portfolio

	fiveDollars := stocks.NewMoney(5, "USD")
	tenEuros := stocks.NewMoney(10, "EUR")

	portfolio = portfolio.Add(fiveDollars)
	portfolio = portfolio.Add(tenEuros)
	actualMoney, err := portfolio.Evaluate(bank, "USD")

	expectedMoney := stocks.NewMoney(17, "USD")

	assertNil(t, err)
	assertEqual(t, expectedMoney, *actualMoney)
}

func TestAdditionOfDollarsAndWons(t *testing.T) {
	var portfolio stocks.Portfolio

	oneDollar := stocks.NewMoney(1, "USD")
	elevenHundredWon := stocks.NewMoney(1100, "KRW")

	portfolio = portfolio.Add(oneDollar)
	portfolio = portfolio.Add(elevenHundredWon)
	actualMoney, err := portfolio.Evaluate(bank, "KRW")

	expectedMoney := stocks.NewMoney(2200, "KRW")

	assertNil(t, err)
	assertEqual(t, expectedMoney, *actualMoney)
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
	value, actualError := portfolio.Evaluate(bank, "Kalganid")

	assertNil(t, value)
	assertEqual(t, expectedErrorMessage, actualError.Error())
}

func TestConversion(t *testing.T) {
	bank := stocks.NewBank()
	bank.AddExchangeRate("EUR", "USD", 1.2)
	tenEuros := stocks.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "USD")
	assertNil(t, err)
	assertEqual(t, stocks.NewMoney(12, "USD"), *actualConvertedMoney)
}

func TestConversionWithMissingExchangeRate(t *testing.T) {
	bank := stocks.NewBank()
	tenEuros := stocks.NewMoney(10, "EUR")
	actualConvertedMoney, err := bank.Convert(tenEuros, "Kalganid")
	assertNil(t, actualConvertedMoney)
	assertEqual(t, "EUR->Kalganid", err.Error())
}

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if actual != expected {
		t.Errorf("Expected %+v, got %+v", expected, actual)
	}
}

func assertNil(t *testing.T, actual interface{}) {
	if actual != nil && !reflect.ValueOf(actual).IsNil() {
		t.Errorf("Expected to be nil, found: [%+v]", actual)
	}
}
