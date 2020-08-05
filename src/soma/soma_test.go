package main

import "testing"

func TestSoma(t *testing.T) {
	total := Soma(5, 5)
	if total != 10 {
		t.Error("Soma incorreta")
	}
}
