package diccionario

import (
	TDAPila "Algogram/tdas/pila"
)

type hijo string

const (
	_ESTANDAR         = 0
	_COMPARACION      = 0
	_HIJO_IZQ    hijo = "izq"
	_HIJO_DER    hijo = "der"
	_SIN_HIJOS   hijo = "sin"
)

type nodoAbb[K comparable, V any] struct {
	clave K
	dato  V
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{nil, _ESTANDAR, funcion_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{clave, dato, nil, nil}
}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	act, ant := a.buscar(a.raiz, clave, a.raiz)
	if act != nil {
		act.dato = dato
	} else {
		nuevo := crearNodo(clave, dato)
		if a.raiz == nil {
			a.raiz = nuevo
		} else {
			valor := a.cmp(clave, ant.clave)
			if valor < _COMPARACION {
				ant.izq = nuevo
			} else {
				ant.der = nuevo
			}
		}
		a.cantidad++
	}
}

func (a *abb[K, V]) Pertenece(clave K) bool {
	act, _ := a.buscar(a.raiz, clave, a.raiz)
	return act != nil
}

func (a *abb[K, V]) Borrar(clave K) V {
	act, ant := a.buscar(a.raiz, clave, a.raiz)
	if act == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := act.dato
	valor := a.cmp(act.clave, ant.clave)
	if act.izq == nil && act.der == nil {
		borrarHijos(a, act, ant, valor, _SIN_HIJOS)
	} else if act.izq != nil && act.der == nil {
		borrarHijos(a, act, ant, valor, _HIJO_IZQ)
	} else if act.izq == nil && act.der != nil {
		borrarHijos(a, act, ant, valor, _HIJO_DER)
	} else if act.izq != nil && act.der != nil {
		borrarDosHijos(a, act, valor)
	}
	a.cantidad--
	return dato
}

func borrarHijos[K comparable, V any](a *abb[K, V], act *nodoAbb[K, V], ant *nodoAbb[K, V], valor2 int, borrado hijo) {
	switch borrado {
	case _HIJO_IZQ:
		if act == a.raiz {
			a.raiz = act.izq
		} else {
			if valor2 < _COMPARACION {
				ant.izq = act.izq
			} else {
				ant.der = act.izq
			}
		}
	case _HIJO_DER:
		if act == a.raiz {
			a.raiz = act.der
		} else {
			if valor2 < _COMPARACION {
				ant.izq = act.der
			} else {
				ant.der = act.der
			}
		}
	case _SIN_HIJOS:
		if act == a.raiz {
			a.raiz = nil
		} else {
			if valor2 < _COMPARACION {
				ant.izq = nil
			} else {
				ant.der = nil
			}
		}
	}
}

func borrarDosHijos[K comparable, V any](a *abb[K, V], act *nodoAbb[K, V], valor2 int) {
	nodo, ant := reemplazo(a, act)
	valor3 := a.cmp(nodo.clave, ant.clave)
	if nodo.izq == nil && nodo.der == nil {
		borrarHijos(a, nodo, ant, valor3, _SIN_HIJOS)
	} else if nodo.der != nil {
		borrarHijos(a, nodo, ant, valor3, _HIJO_DER)
	}
	act.clave = nodo.clave
	act.dato = nodo.dato
}

func reemplazo[K comparable, V any](a *abb[K, V], act *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	ant := act
	nodo := act.der
	for nodo.izq != nil {
		ant = nodo
		nodo = nodo.izq
	}
	return nodo, ant
}

func (a *abb[K, V]) Obtener(clave K) V {
	act, _ := a.buscar(a.raiz, clave, a.raiz)
	if act == nil {
		panic("La clave no pertenece al diccionario")
	}
	return act.dato
}

func (a *abb[K, V]) buscar(act *nodoAbb[K, V], clave K, ant *nodoAbb[K, V]) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if act == nil {
		return act, ant
	}
	valor := a.cmp(clave, act.clave)
	if valor == _COMPARACION {
		return act, ant
	} else if valor < _COMPARACION {
		return a.buscar(act.izq, clave, act)
	} else {
		return a.buscar(act.der, clave, act)
	}
}

func (a *abb[K, V]) Cantidad() int { return a.cantidad }

func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) { a.IterarRango(nil, nil, visitar) }

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	iterarRango(a.raiz, desde, hasta, visitar, a)
}

func iterarRango[K comparable, V any](n *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, a *abb[K, V]) bool {
	var brake bool
	if n == nil {
		return true
	}
	if desde != nil && a.cmp(n.clave, *desde) < _COMPARACION {
		return iterarRango(n.der, desde, hasta, visitar, a)
	}
	if hasta != nil && a.cmp(n.clave, *hasta) > _COMPARACION {
		return iterarRango(n.izq, desde, hasta, visitar, a)
	}
	brake = iterarRango(n.izq, desde, hasta, visitar, a)
	if !brake {
		return false
	}
	if !visitar(n.clave, n.dato) {
		return false
	}
	return iterarRango(n.der, desde, hasta, visitar, a)
}

func (a *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return a.IteradorRango(nil, nil)
}

type iteradorRangoAbb[K comparable, V any] struct {
	a     *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iteradorRangoAbb := iteradorRangoAbb[K, V]{a, pila, desde, hasta}
	apilarIzqRango(&iteradorRangoAbb, a.raiz)
	return &iteradorRangoAbb
}

func (i *iteradorRangoAbb[K, V]) HaySiguiente() bool { return !i.pila.EstaVacia() }

func (i *iteradorRangoAbb[K, V]) VerActual() (K, V) {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := i.pila.VerTope()
	return nodo.clave, nodo.dato
}

func (i *iteradorRangoAbb[K, V]) Siguiente() K {
	if !i.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := i.pila.Desapilar()
	if nodo.der != nil {
		apilarIzqRango(i, nodo.der)
	}
	return nodo.clave
}

func apilarIzqRango[K comparable, V any](i *iteradorRangoAbb[K, V], n *nodoAbb[K, V]) {
	if n == nil {
		return
	}
	if i.desde != nil && i.a.cmp(n.clave, *i.desde) < _COMPARACION {
		apilarIzqRango(i, n.der)
		return
	}
	if i.hasta != nil && i.a.cmp(n.clave, *i.hasta) > _COMPARACION {
		apilarIzqRango(i, n.izq)
		return
	}
	i.pila.Apilar(n)
	apilarIzqRango(i, n.izq)
}
