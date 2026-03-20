package src

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() string {
	return "Hola, soy " + p.Nombre
}

func NewPersona(Nombre string, Edad int) *Persona {
	return &Persona{Nombre, Edad}
}

func (p Persona) GetNombre() string {
	return p.Nombre
}

func (p Persona) GetEdad() int {
	return p.Edad
}
