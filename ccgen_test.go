package ccgen_test

import (
	"testing"

	"github.com/nsuprun/ccgen"
)

var cardsToTests = []ccgen.CardType{
	ccgen.AmericanExpress,
	ccgen.DinersClub,
	ccgen.DinersClubUS,
	ccgen.Discover,
	ccgen.JCB,
	ccgen.Laser,
	ccgen.Maestro,
	ccgen.Mastercard,
	ccgen.Solo,
	ccgen.Unionpay,
	ccgen.Visa,
	ccgen.Mir,
}

func TestGenerate(t *testing.T) {

	for _, cardType := range cardsToTests {
		cnumber := cardType.Generate()
		Suite(t, cnumber, cardType)
	}
}

func TestValidLength(t *testing.T) {
	if ccgen.AmericanExpress.ValidLength(16) {
		t.Error("Length is incorrect")
	}

	if ccgen.Visa.ValidLength(1) {
		t.Error("Length is incorrect")
	}
}

func TestGenerateOfLength(t *testing.T) {
	// test with valid card length
	cnumber := ccgen.Visa.GenerateOfLength(13)
	Suite(t, cnumber, ccgen.Visa)

	// test with invalid card length
	cnumber = ccgen.Mastercard.GenerateOfLength(100)
	Suite(t, cnumber, ccgen.Mastercard)
}

func TestValidNumber(t *testing.T) {
	cnumber := ccgen.Visa.GenerateOfLength(13)

	// too short card length
	if ccgen.Visa.ValidNumber(cnumber[:len(cnumber)-1]) {
		t.Error("Card number is not valid")
	}

}

func Suite(t *testing.T, cnumber string, card ccgen.CardType) {
	// check length is valid
	if !card.ValidLength(len(cnumber)) {
		t.Error("Not valid card length", len(cnumber), card)
	}

	// Luhn check
	if !card.ValidNumber(cnumber) {
		t.Error("Luhn check failed: ", cnumber)
	}
}
