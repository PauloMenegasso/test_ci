package main

import "testing"

func TestSoma(t *testing.T) {
	firstArgument := 15
	secondArgument := 15

	total := soma(firstArgument, secondArgument)
	expectedValue := firstArgument + secondArgument

	if total != expectedValue {
		t.Errorf("Valor esperado e valor encontrado não são iguais. Valor encontrado: %d, valor esperado: %d", total, expectedValue)
	}
}
