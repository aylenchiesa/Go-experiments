package src

type Biblioteca struct {
	Items []IPrestable
}

func (b *Biblioteca) PrestarLibro(u *Usuario, titulo string) bool {

	for _, item := range b.Items {

		switch i := item.(type) {

		case *Book:
			if i.Titulo == titulo && i.Disponible && u.PuedePedirLibro() {
				i.Prestar()
				u.PedirLibro()
				return true
			}

		case *Tesis:
			if i.Titulo == titulo && i.Disponible && u.PuedePedirLibro() {
				i.Prestar()
				u.PedirLibro()
				return true
			}
		}
	}

	return false
}