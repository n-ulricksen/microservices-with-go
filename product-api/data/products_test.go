package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "nicks",
		Price: 1.00,
		SKU:   "abs-asdkfljs-jfe",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
