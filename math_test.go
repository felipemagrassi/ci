package main

import "testing"

func TestSoma(t *testing.T) {
	if Soma(1, 1) != 3 {
		t.Errorf("Resultado invalido para soma %d", Soma(1, 1))
	}
}
