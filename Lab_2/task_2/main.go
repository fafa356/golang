package main

import "fmt"

type Currency interface {
	Name() string
	USDRate() float64
	IsCrypto() bool
}

type USD struct {
	name     string
	UsdRate  float64
	isCrypto bool
}
type Bitcoin struct {
	name     string
	UsdRate  float64
	isCrypto bool
}

func (u USD) Name() string {
	return u.name
}
func (u USD) USDRate() float64 {
	return u.UsdRate
}

func (u USD) IsCrypto() bool {
	return u.isCrypto
}

func (b Bitcoin) Name() string {
	return b.name
}
func (b Bitcoin) USDRate() float64 {
	return b.UsdRate
}

func (b Bitcoin) IsCrypto() bool {
	return b.isCrypto
}

func main() {
	bitcoin := Bitcoin{"Bitcoin", 67720.21, true}
	usd := USD{"USD", 1.0, false}

	currencies := []Currency{bitcoin, usd}
	for _, currency := range currencies {
		fmt.Println(currency.Name(), currency.USDRate())
	}
}
