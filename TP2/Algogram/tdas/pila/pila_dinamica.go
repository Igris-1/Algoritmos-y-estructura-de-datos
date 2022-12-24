package pila

const (
	_CAPACIDAD_INICIAL  = 10
	_AUMENTAR_CAPACIDAD = 2.0
	_REDUCIR_CAPACIDAD  = 0.5
	_MULTIPLO           = 4
	_CANTIDAD_INICIAL   = 0
)

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func crearArreglo[T any](capacidad int) []T { return make([]T, capacidad) }

func CrearPilaDinamica[T any]() *pilaDinamica[T] {
	return &pilaDinamica[T]{datos: crearArreglo[T](_CAPACIDAD_INICIAL)}
}

func (p *pilaDinamica[T]) Apilar(dato T) {
	if p.cantidad == cap(p.datos) {
		p.redimensionar(_AUMENTAR_CAPACIDAD)
	}
	p.datos[p.cantidad] = dato
	p.cantidad++
}

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

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) EstaVacia() bool { return p.cantidad == _CANTIDAD_INICIAL }

func (p *pilaDinamica[T]) redimensionar(valor float64) {
	nuevosDatos := crearArreglo[T](int(float64(cap(p.datos)) * valor))
	copy(nuevosDatos, p.datos)
	p.datos = nuevosDatos
}
