package gomath

import (
	"fmt"
	"math"
)

// Returns the roots of a quadratic equation. Receives A, B, and C as float64 values.
func Bhaskara(A, B, C float64) (float64, float64, error) {
	var X1, X2 float64

	// Calculate the discriminant
	var Discriminant float64
	Discriminant = B*B - 4*A*C

	if Discriminant < 0 {
		return 0, 0, fmt.Errorf("Negative discriminant, %v! There are no real roots.", Discriminant)
	}

	// Calculate the roots
	X1 = (float64(-B) + math.Sqrt(float64(Discriminant))) / float64(A*2)

	X2 = (float64(-B) - math.Sqrt(float64(Discriminant))) / float64(A*2)

	// Resultado
	return X2, X1, nil

}
