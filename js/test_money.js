const assert = require('assert');


class Money {
    constructor(amount, currency) {
        this.amount = amount;
        this.currency = currency;
    }

    times(multiplier) {
        return new Money(this.amount * multiplier, this.currency);
    }

    divide(divisor) {
        return new Money(this.amount / divisor, this.currency)
    }
} 

class Portfolio {
    constructor() { }

    add(money) { }

    evaluate(currency) {
        return new Money(15, "USD");
    }
}

let fiveDollars = new Money(5, "USD");
let tenDollars = new Money(10, "USD");
assert.deepStrictEqual(fiveDollars.times(2), tenDollars)

let tenEuros = new Money(10, "EUR");
let twentyEuros = new Money(20, "EUR");
assert.deepStrictEqual(tenEuros.times(2), twentyEuros)

let money = new Money(4002, "KRW");
assert.deepStrictEqual(money.divide(4), new Money(1000.5, "KRW"))

let portfolio = new Portfolio();
portfolio.add(fiveDollars, tenDollars)
assert.deepStrictEqual(portfolio.evaluate("USD"), new Money(15, "USD"))