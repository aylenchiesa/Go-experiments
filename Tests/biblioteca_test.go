package tests

import (
	"Go-experiments/src"
	"testing"
)

func TestPrestarLibro(t *testing.T) {

	book := src.Book{
		Item: src.Item{
			ID:     1,
			Titulo: "Go Básico",
		},
		Autor:      "Autor",
		Disponible: true,
	}

	ok := book.Prestar()

	if !ok {
		t.Errorf("El libro debería poder prestarse")
	}

	if book.Disponible {
		t.Errorf("El libro debería quedar no disponible")
	}
	book.Devolver()

}

func TestDevolverTesis(t *testing.T) {
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

func TestNoPrestarLibroPorLimiteUsuario(t *testing.T) {

	b := src.Biblioteca{ // creo biblioteca
		Items: []src.IPrestable{
			&src.Book{
				Item: src.Item{
					ID:     1,
					Titulo: "El Diario de Anne Frank",
				},
				Autor:      "Autor",
				Disponible: true,
			},
		},
	}

	u := src.Usuario{ // usuario ya en el límite
		Nombre:              "Aylén",
		LibrosPrestados:     2,
		MaxLibrosPermitidos: 2,
	}

	ok := b.PrestarLibro(&u, "El Diario de Anne Frank") // intento prestar

	if ok {
		t.Errorf("No debería poder prestar más libros porque alcanzó el límite")
	}

	if u.LibrosPrestados != 2 {
		t.Errorf("La cantidad de libros prestados no debería cambiar")
	}
}