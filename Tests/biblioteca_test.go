package tests

import (
	"Go-experiments/src"
	"testing"
)

/*
func TestPrestarLibro(t *testing.T) {
	b := src.Biblioteca{ //creo objeto biblioteca
		Libros: []src.Book{
			{Titulo: "El Diario de Anne Frank", Autor: "Autor", Disponible: true},
		},
	}

	u := src.Usuario{ //creo objeto usuario
		Nombre:              "Aylén",
		MaxLibrosPermitidos: 2,
	}

	ok := b.PrestarLibro(&u, "El Diario de Anne Frank") //llamo al metodo prestarlibro

	if !ok {
		t.Errorf("Deberia poder prestar el libro")
	}

	if u.LibrosPrestados != 1 {
		t.Errorf("El usuario debería tener 1 libro prestado")
	}
}*/

func TestPrestarYDevolverTesis(t *testing.T) {
	// creamos una tesis disponible
	tesis := src.Tesis{
		Item: src.Item{
			ID:     1,
			Titulo: "Inteligencia Artificial",
		},
		Autor:      "Aylén",
		Area:       "Informática",
		Disponible: true,
	}

	// prestamos la tesis
	ok := tesis.Prestar()

	if !ok {
		t.Errorf("La tesis debería poder prestarse")
	}

	if tesis.Disponible {
		t.Errorf("La tesis debería quedar no disponible después de prestarla")
	}

	// devolvemos la tesis
	tesis.Devolver()

	if !tesis.Disponible {
		t.Errorf("La tesis debería estar disponible después de devolverla")
	}
}
