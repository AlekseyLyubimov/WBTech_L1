package main

import (
	"math/big"
)

func main() {
	first := big.NewFloat(33554432)
	second := big.NewFloat(67108864)

	println("multiplication")
	println(big.NewFloat(0).Mul(first, second).String())

	println("division")
	println(big.NewFloat(0).Quo(second, first).String())

	println("addition")
	println(big.NewFloat(0).Add(first, second).String())

	println("subtraction")
	println(big.NewFloat(0).Sub(second, first).String())
}
