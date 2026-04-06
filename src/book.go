package src

type Book struct {
	Item       //composicionnnnn
	Autor      string
	Disponible bool
}

func (l *Book) Prestar() bool {
	ok := l.Disponible
	l.Disponible = false
	return ok
}

func (l *Book) Devolver() {
	l.Disponible = true
}

func (t *Book) EstaDisponible() bool {
	return t.Disponible
}