// run: go run index.go
package main

import (
	"fmt"
	"math"
)

// import method from hello.go

func main() {
	//printHelloWorld()
	calculate()
}

func calculate() {
	stockPrice := 95.0
	strikePrice := 100.0
	riskFreeRate := 0.05
	timeToExpiration := 1.0
	volatility := 0.2

	contractMaxPrice := 10.0

	yesPrice := blackScholesOption(stockPrice, strikePrice, riskFreeRate, timeToExpiration, volatility, "call")
	noPrice := blackScholesOption(stockPrice, strikePrice, riskFreeRate, timeToExpiration, volatility, "put")

	fmt.Printf("The value of the call option is: %.2f\n", yesPrice)
	fmt.Printf("The value of the put option is: %.2f\n", noPrice)

	if yesPrice > noPrice {
		x := noPrice * 100 / yesPrice
		y := 100 - x
		putPrice := (x * contractMaxPrice) / 100
		callPrice := (y * contractMaxPrice) / 100
		fmt.Printf("yesFavourable: x = %.2f%% y = %.2f%% \n", x, y)
		fmt.Printf("yesPrice = %.2f noPrice = %.2f\n", callPrice, putPrice)
	} else {
		x := yesPrice * 100 / noPrice
		y := 100 - x
		putPrice := (y * contractMaxPrice) / 100
		callPrice := (x * contractMaxPrice) / 100
		fmt.Printf("noFavourable: x = %.2f%% y = %.2f%% \n", x, y)
		fmt.Printf("yesPrice = %.2f noPrice = %.2f\n", callPrice, putPrice)
	}

}

// blackScholesOption calculates the value of a European call or put option using the Black-Scholes formula.
// S:      Current stock price
// K:      Strike price of the option
// r:      Risk-free interest rate
// T:      Time to option expiration in years
// sigma:  Stock's volatility (standard deviation)
// optionType: Type of option - "call" or "put"
// Returns the value of the option
func blackScholesOption(S, K, r, T, sigma float64, optionType string) float64 {
	d1 := (math.Log(S/K) + (r+0.5*math.Pow(sigma, 2))*T) / (sigma * math.Sqrt(T))
	d2 := d1 - sigma*math.Sqrt(T)

	var Nd1, Nd2 float64
	if optionType == "call" {
		Nd1 = 0.5 * (1 + math.Erf(d1/math.Sqrt(2)))
		Nd2 = 0.5 * (1 + math.Erf(d2/math.Sqrt(2)))
	} else if optionType == "put" {
		Nd1 = 0.5 * (1 - math.Erf(d1/math.Sqrt(2)))
		Nd2 = 0.5 * (1 - math.Erf(d2/math.Sqrt(2)))
	} else {
		panic("Invalid option type. Must be either 'call' or 'put'.")
	}

	var optionValue float64
	if optionType == "call" {
		optionValue = S*Nd1 - K*math.Exp(-r*T)*Nd2
	} else {
		optionValue = K*math.Exp(-r*T)*Nd2 - S*Nd1
	}

	return optionValue
}
