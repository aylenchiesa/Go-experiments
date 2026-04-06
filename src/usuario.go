package src

type Usuario struct {
	Nombre              string
	LibrosPrestados     int
	MaxLibrosPermitidos int
}

func (u *Usuario) PuedePedirLibro() bool {
	return u.LibrosPrestados < u.MaxLibrosPermitidos
}

func (u *Usuario) PedirLibro() {
	u.LibrosPrestados++
}

/*func (u *Usuario) Devolver() {
	if u.LibrosPrestados > 0 {
		u.LibrosPrestados--
	}
}*/