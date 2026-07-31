package gomath

import (
	"fmt"
	"math"
)

// Retorna as Raízes de uma equação do Segundo Grau
func Bhaskara() {
	var A, B, C, X1, X2 float64
	// Pergunta o valor de A, B e C
	fmt.Print("Qual o Valor de A: ")
	fmt.Scanln(&A)

	fmt.Print("Qual o valor de B: ")
	fmt.Scanln(&B)

	fmt.Print("Qual o valor de C: ")
	fmt.Scanln(&C)

	// Calcula o Delta
	var Delta float64
	Delta = B*B - 4*A*C

	if Delta < 0 {
		fmt.Printf("Delta negativo, %v! Não existem raízes reais.", Delta)
		return
	}

	// Calcula as raízes
	X1 = (float64(-B) + math.Sqrt(float64(Delta))) / float64(A*2)

	X2 = (float64(-B) - math.Sqrt(float64(Delta))) / float64(A*2)

	// Resultado
	if Delta == 0 {
		fmt.Printf("A raiz da equação é: %.1f\n", X1)
	} else {
		fmt.Printf("As raízes da equação são: %.1f e %.1f\n", X1, X2)
	}

	// Encerrar
	var encerrar int
	fmt.Println("Digite qualquer coisa e pressione Enter para encerrar o programa!")
	fmt.Scan(&encerrar)
}
