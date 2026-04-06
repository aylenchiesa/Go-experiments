package src

type IPrestable interface {
	Prestar() bool
	Devolver()
	EstaDisponible() bool
}
