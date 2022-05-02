const assert = require('assert');
const Money = require('./money')
const Portfolio = require('./portfolio')


let tenEuros = new Money(10, "EUR");
let twentyEuros = new Money(20, "EUR");
assert.deepStrictEqual(tenEuros.times(2), twentyEuros)

let money = new Money(4002, "KRW");
assert.deepStrictEqual(money.divide(4), new Money(1000.5, "KRW"))

let fiveDollars = new Money(5, "USD");
let tenDollars = new Money(10, "USD");
let portfolio = new Portfolio();
portfolio.add(fiveDollars, tenDollars)
assert.deepStrictEqual(portfolio.evaluate("USD"), new Money(15, "USD"))
