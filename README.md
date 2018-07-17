
# ccgen
Minimalistic Go package for a random Luhn-compliant credit card numbers generation

[![Go Report Card](https://goreportcard.com/badge/github.com/nsuprun/ccgen)](https://goreportcard.com/report/github.com/nsuprun/ccgen)  [![Build Status](https://travis-ci.org/nsuprun/ccgen.svg?branch=master)](https://travis-ci.org/nsuprun/ccgen)

## Introduction
The package allows generating random credit card numbers according to the [Luhn algorithm](https://en.wikipedia.org/wiki/Luhn_algorithm). Card prefix (first N digits of a card number) will correspond to the selected card scheme.

## Usage
Install the package with:

```
go get github.com/nsuprun/ccgen
```
Import:
```go
import (
	"github.com/nsuprun/ccgen"
)
```
At the moment the following card schemes are supported:
```go
const (
	AmericanExpress CardType = iota + 1
	DinersClub
	DinersClubUS
	Discover
	JCB
	Laser
	Maestro
	Mastercard
	Solo
	Unionpay
	Visa
	Mir
)
```

```go
func main() {
	// To generate a card number of random valid length
	fmt.Printf("American Express: %s\n", ccgen.AmericanExpress.Generate())
	fmt.Printf("Diners Club: %s\n", ccgen.DinersClub.Generate())

	// To generate a card number of selected valid length:
	// Solo 19 digits card
	fmt.Printf("Solo: %s\n", ccgen.Solo.GenerateOfLength(19))
}
```

Output:
```
American Express: 346905926744572
Diners Club: 36690592674457
Solo: 6334690592674457710
```