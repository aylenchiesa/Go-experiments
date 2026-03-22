package src

type Libro struct {
	ID         int
	Titulo     string
	Autor      string
	Disponible bool
}

func (l *Libro) Prestar() bool {
	ok := l.Disponible
	l.Disponible = false
	return ok
}

func (l *Libro) Devolver() {
	l.Disponible = true
}
