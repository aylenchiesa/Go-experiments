package src

type Tesis struct {
	Item
	Autor      string
	Area       string
	Disponible bool
}

func (t *Tesis) Prestar() bool {
	ok := t.Disponible
	t.Disponible = false
	return ok
}

func (t *Tesis) Devolver() {
	t.Disponible = true
}
