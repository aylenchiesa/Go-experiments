package tests

import (
	"Go-experiments/src"
	"testing"
)

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

}
