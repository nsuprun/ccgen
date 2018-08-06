package ccgen

import (
	"math/rand"
	"strconv"
	"time"
)

// CardType represents card number type to be generated
type CardType int

// Generate generates card number of a random card type
func Generate() string {
	return getRandomCardType().Generate()
}

// Returns random card type
// @todo: consider other way to limit highest type
func getRandomCardType() CardType {
	rand.Seed(time.Now().Unix())
	randType := rand.Intn(int(Mir))

	if randType == 0 {
		randType++
	}

	return CardType(randType)
}

// Generate generates random card number of CardType
func (t CardType) Generate() string {
	length := t.getRandomValidLength()
	binNumber := t.getRandomCardPrefix()

	return completeDigits(binNumber, length)
}

// GenerateOfLength generates card number of specified length
// If length is not valid, random among valid ones will be used
func (t CardType) GenerateOfLength(l int) string {
	if !t.ValidLength(l) {
		l = t.getRandomValidLength()
	}

	binNumber := t.getRandomCardPrefix()

	return completeDigits(binNumber, l)
}

// ValidLength validates if specified card length is valid for the card type
func (t CardType) ValidLength(l int) bool {
	for _, n := range cardLength[t] {
		if n == l {
			return true
		}
	}
	return false
}

// Returns random card prefix among valid for the card type
func (t CardType) getRandomCardPrefix() string {
	rand.Seed(time.Now().Unix())

	return cardPrefix[t][rand.Intn(len(cardPrefix[t]))]
}

// Returns random card length among valid for the card type
func (t CardType) getRandomValidLength() int {
	rand.Seed(time.Now().Unix())
	return cardLength[t][rand.Intn(len(cardLength[t]))]
}

// ValidNumber validates card number for a prefix and Luhn algo
// @todo: add card prefix check
func (t CardType) ValidNumber(cnumber string) bool {
	return checkLuhn(cnumber)
}

// Checks if the card number has valid check digit
func checkLuhn(cnumber string) bool {
	if len(cnumber) < 12 || len(cnumber) > 19 {
		return false
	}
	checkDigit := cnumber[len(cnumber)-1:]
	return checkDigit == generateCheckDigit(cnumber[:len(cnumber)-1])
}

// generates the rest of the digits after the card prefix
// leaving the last position for the check digit
func completeDigits(bin string, l int) string {
	rand.Seed(time.Now().Unix())

	randomNumberLength := l - (len(bin) + 1)

	for i := 0; i < randomNumberLength; i++ {
		rand := rand.Intn(10)
		bin += strconv.Itoa(rand)
	}

	checkDigit := generateCheckDigit(bin)
	return bin + checkDigit
}

// Generates last digit according to Luhn algo
func generateCheckDigit(cnumber string) string {
	sum := 0
	for i := 0; i < len(cnumber); i++ {
		digit, _ := strconv.Atoi((string(cnumber[i])))

		if i%2 == 0 {
			digit = digit * 2

			if digit > 9 {
				digit = (digit / 10) + (digit % 10)
			}

		}

		sum += digit

	}

	mod := sum % 10

	if mod == 0 {
		return strconv.Itoa(0)
	}

	return strconv.Itoa(10 - mod)
}
