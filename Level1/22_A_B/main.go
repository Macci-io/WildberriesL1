package main

import (
	"fmt"
	"math/big"
)

type Ab struct {
	a, b *big.Int
}

func (a *Ab) Add() *big.Int {
	return new(big.Int).Add(a.a, a.b)
}

func (a *Ab) Multiple() *big.Int {
	return new(big.Int).Mul(a.a, a.b)
}

func (a *Ab) Sub() *big.Int {
	return new(big.Int).Sub(a.a, a.b)
}

func (a *Ab) Div() *big.Int {
	return new(big.Int).Div(a.a, a.b)
}

func main() {
	pool := &Ab{
		a: new(big.Int),
		b: new(big.Int),
	}
	pool.a = pool.a.SetInt64(1048576)
	pool.b = pool.b.SetInt64(1048576)
	fmt.Println("Плюс:", pool.Add())
	fmt.Println("Минус:", pool.Sub())
	fmt.Println("Усножение:", pool.Multiple())
	fmt.Println("Деление:", pool.Div())

}
