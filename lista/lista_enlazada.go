package lista

// struct Nodo
type nodo[T any] struct {
	dato T
	sig  *nodo[T]
}

// funcion auxiliar crearNodo
func crearNodo[T any](dato T) *nodo[T] {
	return &nodo[T]{dato: dato}
}

/* implementacion listaEnlazada */

// struct listaEnlazada
type listaEnlazada[T any] struct {
	prim  *nodo[T]
	ult   *nodo[T]
	largo int
}

// devuelve una listaEnlazada
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{}
}

func (l listaEnlazada[T]) EstaVacia() bool {
	return l.prim == nil && l.ult == nil && l.largo == 0
}

func (l *listaEnlazada[T]) InsertarPrimero(dato T) {
	nuevo := crearNodo(dato)
	if l.EstaVacia() {
		l.ult = nuevo
	} else {
		nuevo.sig = l.prim
	}
	l.prim = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) InsertarUltimo(dato T) {
	nuevo := crearNodo(dato)
	if l.EstaVacia() {
		l.prim = nuevo
	} else {
		l.ult.sig = nuevo
	}
	l.ult = nuevo
	l.largo++
}

func (l *listaEnlazada[T]) BorrarPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	elemento_borrar := l.prim.dato
	l.prim = l.prim.sig
	if l.prim == nil {
		l.ult = nil
	}
	l.largo--
	return elemento_borrar
}

func (l listaEnlazada[T]) VerPrimero() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.prim.dato
}

func (l listaEnlazada[T]) VerUltimo() T {
	if l.EstaVacia() {
		panic("La lista esta vacia")
	}
	return l.ult.dato
}

func (l listaEnlazada[T]) Largo() int {
	return l.largo
}

/* Implementacion iterListaEnlazada */

func (l *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	for act := l.prim; act != nil; act = act.sig {
		if !visitar(act.dato) {
			break
		}
	}
}

// iterador externo
// struct iterListaEnlazada
type iterListaEnlazada[T any] struct {
	lista *listaEnlazada[T]
	ant   *nodo[T]
	act   *nodo[T]
}

func (l *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{lista: l, act: l.prim}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.act.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.act != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	valor_actual := iterador.act.dato
	iterador.ant = iterador.act
	iterador.act = iterador.act.sig
	return valor_actual
}

func (iterador *iterListaEnlazada[T]) Insertar(dato T) {
	nuevo := crearNodo(dato)
	if iterador.lista.EstaVacia() {
		iterador.lista.prim = nuevo
		iterador.lista.ult = nuevo
	} else if iterador.act == iterador.lista.prim {
		nuevo.sig = iterador.lista.prim
		iterador.lista.prim = nuevo
	} else {
		iterador.ant.sig = nuevo
		nuevo.sig = iterador.act
		if nuevo.sig == nil {
			iterador.lista.ult = nuevo
		}
	}
	iterador.act = nuevo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento_borrar := iterador.act.dato
	if iterador.act == iterador.lista.prim {
		iterador.lista.prim = iterador.lista.prim.sig
		iterador.act = iterador.lista.prim
		iterador.ant = iterador.lista.prim
		if iterador.lista.prim == nil {
			iterador.lista.ult = nil
		}
	} else if iterador.act == iterador.lista.ult {
		iterador.ant.sig = nil
		iterador.lista.ult = iterador.ant
		iterador.act = nil
	} else {
		iterador.act = iterador.act.sig
		iterador.ant.sig = iterador.act
	}
	iterador.lista.largo--
	return elemento_borrar
}
