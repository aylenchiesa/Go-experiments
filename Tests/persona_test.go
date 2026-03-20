package tests

import (
	src "Go-experiments/src"
	"testing"
)

func TestSaludar(t *testing.T) {
	p := src.Persona{Nombre: "Aylén", Edad: 20}
	resultado := p.Saludar()

	esperado := "Hola, soy Aylén"

	if resultado != esperado {
		t.Errorf("esperado %s pero obtuve %s", esperado, resultado)
	}
}
