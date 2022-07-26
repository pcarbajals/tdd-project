const Money = require('./money')

class Portfolio {
    constructor() {
        this.monies = new Array();
    }

    add(...monies) {
        this.monies = this.monies.concat(monies);
    }

    evaluate(bank, currency) {
        let failures = [];
        let total = this.monies.reduce((sum, money) => {
            try {
                let convertedMoney = bank.convert(money, currency);
                return sum + convertedMoney.amount;

            } catch (error) {
                failures.push(error.message);
                return sum;
            }
        }, 0);
        if (!failures.length) {
            return new Money(total, currency);
        }
        throw new Error("Missing exchange rate(s):[" + failures.join() + "]");
    }
}

module.exports = Portfolio
