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
            return sum + money.amount
        }, 0);
        return new Money(total, currency);
    }
}

module.exports = Portfolio
