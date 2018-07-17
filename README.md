# ccgen
Minimalistic Go package for random credit card numbers generation

## Usage
Intall the package with:

```
go get github.com/nsuprun/ccgen
```
Import
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

```go
func main() {
	// To generate card number of random valid length
	fmt.Printf("American Express: %s\n", ccgen.AmericanExpress.Generate())
	fmt.Printf("Diners Club: %s\n", ccgen.DinersClub.Generate())

	// To generate card number of selected valid length:
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