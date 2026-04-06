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
	
		b := src.Biblioteca{
			Items: []src.IPrestable{
				&src.Tesis{
			Item: src.Item{
				ID:     1,
				Titulo: "Inteligencia Artificial",
			},
			Autor:      "Aylén",
			Area:       "Informática",
			Disponible: true,
		},
	},
}

	u := src.Usuario{
			Nombre:              "Martu",
			MaxLibrosPermitidos: 2,
		}

	// prestamos la tesis
	ok := b.Prestar(&u, "Tesis IA")

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

func TestMaximoLibrosPedidos(t *testing.T) {

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

	if u.PuedePedirLibro() {
		t.Errorf("El usuario NO debería poder pedir más libros")
	}

	ok := b.Prestar(&u, "El Diario de Anne Frank") // intento prestar

	if ok {
		t.Errorf("No debería poder prestar más libros porque alcanzó el límite")
	}

	if u.LibrosPrestados != 2 {
		t.Errorf("La cantidad de libros prestados no debería cambiar")
	}
}

func TestDisponibilidad(t *testing.T) {
    // creamos la biblioteca directamente con sus items o sea libros cargados.
    biblioteca := src.Biblioteca{
        Items: []src.IPrestable{
            &src.Book{Item: src.Item{ID: 1, Titulo: "El Principito"}, 
														Autor: "Saint-Exupéry", 
														Disponible: true,
										 }, 
            &src.Book{Item: src.Item{ID: 2, Titulo: "Rayuela"}, Autor: "Julio Cortázar", Disponible: false},
										 
        },
			}

		usuario := &src.Usuario{
		Nombre: "Aylén", 
		LibrosPrestados: 0, 
		MaxLibrosPermitidos: 2,
    }

		libro := biblioteca.Items[0].(*src.Book)

		if !libro.Disponible {
			t.Errorf("El Principito debería estar disponible")
		}

		ok := biblioteca.Prestar(usuario, "El Principito") //  presto

		if !ok {
        t.Errorf("Debería haber podido prestar 'El Principito' porque estaba en la lista y disponible")
    }
}