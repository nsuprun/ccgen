# ccgen
Minimalistic Go package for random credit card numbers generation

## Usage

```go
import (
	"github.com/nsuprun/ccgen"
)
```
At the moment, the following card schemes are supported:
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
To generate card number of random valid length

```go
ccgen.AmericanExpress.Generate()
// 346770125855275

ccgen.Mastercard.Generate()
// 2720243672467447
```

To generate card number of selected valid length:
```go
ccgen.Solo.GenerateOfLength(19)
// 6767137537983398740
```