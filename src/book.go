package src

type Book struct {
	ID         int
	Titulo     string
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
