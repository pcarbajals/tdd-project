const Money = require('./money')

class Portfolio {
    constructor() {
        this.monies = new Array();
    }

    add(...monies) {
        this.monies = this.monies.concat(monies);
    }

    evaluate(currency) {
        let total = this.monies.reduce((sum, money) => {
            return sum + this.convert(money, currency)
        }, 0);
        return new Money(total, currency);
    }

    convert(money, currency) {
        if (money.currency === currency) {
            return money.amount;
        }

        return money.amount * 1.2;
    }
}

module.exports = Portfolio
