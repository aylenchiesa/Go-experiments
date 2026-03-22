package src

type Biblioteca struct {
	Libros []Libro
}

func (b *Biblioteca) BuscarLibro(titulo string) *Libro {
	for i := range b.Libros {
		switch b.Libros[i].Titulo == titulo {
		case true:
			return &b.Libros[i]
		}
	}
	return nil
}

func (b *Biblioteca) PrestarLibro(u *Usuario, titulo string) bool {

	libro := b.BuscarLibro(titulo)

	switch {
	case libro == nil:
		return false
	case !libro.Disponible:
		return false
	case !u.PuedePedirLibro():
		return false
	}

	libro.Prestar()
	u.PedirLibro()

	return true
}
