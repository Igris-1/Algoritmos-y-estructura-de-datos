package pila

const (
	_CAPACIDAD_INICIAL  = 10
	_AUMENTAR_CAPACIDAD = 2.0
	_REDUCIR_CAPACIDAD  = 0.5
	_MULTIPLO           = 4
)

/* Definición del struct pila proporcionado por la cátedra. */
type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

/* Implementación pila */
func CrearPilaDinamica[T any]() *pilaDinamica[T] {
	return &pilaDinamica[T]{datos: make([]T, _CAPACIDAD_INICIAL)}
}

// Apilar agrega un elemento a la pila.
func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(_AUMENTAR_CAPACIDAD)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

// Desapilar devuelve el elemento que está en el tope de la pila.
func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if p.cantidad*_MULTIPLO <= cap(p.datos) && cap(p.datos) > _CAPACIDAD_INICIAL {
		p.redimensionar(_REDUCIR_CAPACIDAD)
	}
	p.cantidad--
	return p.datos[p.cantidad]
}

// ver el tope de la pila
func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

// devuelve true or false si la pila esta vacia
func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

// duplicar / reducir 1/2 la capacidad de la pila
func (p *pilaDinamica[T]) redimensionar(valor float64) {
	nuevosDatos := make([]T, int(float64(cap(p.datos))*valor))
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
