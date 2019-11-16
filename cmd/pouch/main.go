package main

import (
	"fmt"
	"github.com/gandalfmagic/go-test-circleci/internal/currency"
)

func main() {
	p := currency.NewPouch()
	p.Add(currency.Copper, 1)
	p.Add(currency.Copper, 2)
	p.Add(currency.Silver, 20)

	fmt.Println(p)
}
