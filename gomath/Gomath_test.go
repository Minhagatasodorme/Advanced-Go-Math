package gomath

import "testing"

func TestBhaskara(t *testing.T) {
	X1, X2, err := Bhaskara(2, -10, 12)

	if X1 != 3 && X2 != 2 && err != nil {
		t.Errorf("Esperava se 3 e 2, recebido %v e %v\nERROR: %v", X1, X2, err)
	}
}
