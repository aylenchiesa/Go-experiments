package main

type Persona struct {
	Nombre string
	Edad   int
}

func (p Persona) Saludar() string {
	return "Hola, soy " + p.Nombre
}

func main() {
	p = Persona{Nombre: "Aylén", Edad: 20}
}
