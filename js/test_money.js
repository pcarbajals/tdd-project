const assert = require('assert');
const Money = require('./money')
const Portfolio = require('./portfolio')


class MoneyTest {
    testMultiplication() {
        let tenEuros = new Money(10, "EUR");
        let twentyEuros = new Money(20, "EUR");
        assert.deepStrictEqual(tenEuros.times(2), twentyEuros)
    }

    testDivision() {
        let money = new Money(4002, "KRW");
        assert.deepStrictEqual(money.divide(4), new Money(1000.5, "KRW"))
    }

    testAddition() {
        let fiveDollars = new Money(5, "USD");
        let tenDollars = new Money(10, "USD");
        let portfolio = new Portfolio();
        portfolio.add(fiveDollars, tenDollars)
        assert.deepStrictEqual(portfolio.evaluate("USD"), new Money(15, "USD"))
    }

    testAdditionOfDollarsAndEuros() {
        let fiveDollars = new Money(5, "USD");
        let tenEuros = new Money(10, "EUR");
        let portfolio = new Portfolio();
        portfolio.add(fiveDollars, tenEuros)
        assert.deepStrictEqual(portfolio.evaluate("USD"), new Money(17, "USD"))
    }

    testAdditionOfDollarsAndWons() {
        let oneDollar = new Money(1, "USD");
        let elevenHundredWons = new Money(1100, "KRW");
        let portfolio = new Portfolio();
        portfolio.add(oneDollar, elevenHundredWons)
        assert.deepStrictEqual(portfolio.evaluate("KRW"), new Money(2200, "KRW"))
    }

    getAllTestMethods() {
        let moneyPrototype = MoneyTest.prototype;
        let allProps = Object.getOwnPropertyNames(moneyPrototype);
        let testMethods = allProps.filter(p => {
            return typeof moneyPrototype[p] === 'function' && p.startsWith("test");
        });
        return testMethods;
    }

    runAllTest() {
        let counter = 0;
        let testMethods = this.getAllTestMethods();
        testMethods.forEach(m => {
            let method = Reflect.get(this, m);
            try {
                Reflect.apply(method, this, []);
                console.log("%s ... OK", m);

            } catch (error) {
                counter += 1;
                console.log("======================================================================");
                console.log("FAIL: %s", m);
                console.log("----------------------------------------------------------------------");
                if (error instanceof assert.AssertionError) {
                    console.log(error);
                } else {
                    throw error;
                }
                console.log("----------------------------------------------------------------------");
            }
        });

        console.log("\nRan %d tests", testMethods.length);
        if (counter > 0) {
            console.log("FAILED (failures=%d)", counter);
        } else {
            console.log("OK");
        }
    }
}

new MoneyTest().runAllTest();
