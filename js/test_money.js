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


let fiver = new Money(5, "USD");
assert.deepStrictEqual(fiver.times(2), new Money(10, "USD"))

let tenEuros = new Money(10, "EUR");
assert.deepStrictEqual(tenEuros.times(2), new Money(20, "EUR"))

let money = new Money(4002, "KRW");
assert.deepStrictEqual(money.divide(4), new Money(1000.5, "KRW"))
