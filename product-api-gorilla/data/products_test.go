package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test",
		Price: 20,
		SKU:   "abc",
	}

	err := p.Validate()
	if err != nil {
		t.Fatal(err)
	}
}
