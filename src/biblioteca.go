package src

type Biblioteca struct {
	Items []IPrestable
}

func (b *Biblioteca) Prestar(u *Usuario, titulo string) bool {

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

func (b *Biblioteca) Devolver(u *Usuario, titulo string) bool {
    for _, item := range b.Items {
        
        if item.EstaDisponible() == false {  // lógica de comparación de títulos
            item.Devolver() // 
            u.LibrosPrestados-- // decrementa el contador de libros prestados del usuario
            return true
        }
    }
    return false
}