package lista

/* type interface listaEnlazada */

type Lista[T any] interface {

	// EstaVacia devuelve verdadero si la cola no tiene elementos encolados, false en caso contrario.
	EstaVacia() bool

	// InsertarPrimero inserta el dato al inicio de la lista.
	InsertarPrimero(T)

	// InsertarUltimo inserta el dato al final de la lista.
	InsertarUltimo(T)

	// BorrarPrimero borra el dato al inicio de la lista. Si la lista se encontraba vacía, entra el pánico con el
	// mensaje 'La lista esta vacia'.
	BorrarPrimero() T

	// VerPrimero devuelve el elemento al inicio de la lista (el primero). Si la lista se encontraba vacía, entra en
	// pánico con el mensaje 'La lista esta vacia'.
	VerPrimero() T

	// VerUltimo devuelve el elemento al final de la lista (el ultimo). Si la lista se encontraba vacía, entra em
	// pánico con el mensaje 'La lista esta vacia'.
	VerUltimo() T

	// Largo devuelve la cantidad de elementos de la lista
	Largo() int

	// Iterar aplica la funcion pasada por parametro a todos los elementos de la lista, hasta que no hayan más
	// elementos, o la función en cuestión devualva false.
	Iterar(func(T) bool)

	// Iterador devuelve un IteradorLista para esta Lista.
	Iterador() IteradorLista[T]
}

/* type interface iterador de listaEnlazada */

type IteradorLista[T any] interface {

	// VerActual devuelve el elemento actual en el que está situado el iterador, sin avanzarlo. Si no hay siguiente
	// para ver, entra en pánico con el mensaje 'El iterador termino de iterar'.
	VerActual() T

	// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
	// el iterador hay un elemento.
	HaySiguiente() bool

	// Siguiente devuelve el siguiente a ver en la iteracion, y avanza el iterador a la siguiente posicion. Si no
	// HaySiguiente para ver, entra en pánico con el mensaje 'El iterador termino de iterar'.
	Siguiente() T

	// Insertar inserta el elemento entre el elemento anterior (si es que hay) y el actual (si es que hay). El
	// iterador queda apuntando al nuevo elemento insertado.
	Insertar(T)

	// Borrar borra el elemento actual de la lista, y lo devuelve. El elemento anterior queda apuntando al que era
	// el próximo. El iterador queda apuntando al que era el próximo. Si no hay siguiente para ver, entra en pánico
	// con el mensaje 'El iterador termino de iterar'.
	Borrar() T
}
