package data

import "testing"

// The funcname should start with Test and the letter after this word should not be lower case
func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Sunil",
		Price: 5.0,
		SKU:   "abs-basd-aas",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}

}
