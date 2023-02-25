package lista_test

import (
	TDALista "lista"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// Test lista vacia
func TestListaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con lista vacia")
	l := TDALista.CrearListaEnlazada[int]()
	require.True(t, l.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerUltimo() })
	require.Equal(t, 0, l.Largo())
}

// Test lista con dos elemento
func TestListaInicial(t *testing.T) {
	t.Log("Hacemos pruebas con dos elemento")
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarPrimero(1)
	require.False(t, l.EstaVacia())
	require.Equal(t, 1, l.VerPrimero())
	require.Equal(t, 1, l.VerUltimo())
	require.Equal(t, 1, l.Largo())
	l.BorrarPrimero()
	require.True(t, l.EstaVacia())
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.BorrarPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerPrimero() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.VerUltimo() })
	require.Equal(t, 0, l.Largo())
	l.InsertarUltimo(2)
	require.EqualValues(t, 2, l.VerPrimero())
	require.EqualValues(t, 2, l.VerUltimo())
	l.InsertarPrimero(1)
	require.EqualValues(t, 1, l.VerPrimero())
	require.EqualValues(t, 2, l.VerUltimo())
	require.Equal(t, 2, l.Largo())
	l.BorrarPrimero()
	l.BorrarPrimero()
	require.True(t, l.EstaVacia())
}

// Test de volumen con elementos de tipo string
func TestListavolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen con elementos de tipo string")
	l := TDALista.CrearListaEnlazada[string]()
	for i := 0; i < 1000; i++ {
		l.InsertarPrimero(strconv.Itoa(i))
		require.EqualValues(t, strconv.Itoa(i), l.VerPrimero())
	}
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, strconv.Itoa(i), l.VerPrimero())
		require.EqualValues(t, l.VerPrimero(), l.BorrarPrimero())
	}

	for i := 0; i < 1000; i++ {
		l.InsertarUltimo(strconv.Itoa(i))
		require.EqualValues(t, strconv.Itoa(i), l.VerUltimo())
	}
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, strconv.Itoa(i), l.VerPrimero())
		require.EqualValues(t, l.VerPrimero(), l.BorrarPrimero())
	}
	require.True(t, l.EstaVacia())
}

/* Test iterador externo */

// Test insertar posicion inicio, medio y fin
func TestIteradorExterno(t *testing.T) {
	t.Log("Hacemos pruebas de iterador externo: insertar inicio, medio, fin")
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarUltimo(2)
	l.InsertarUltimo(4)
	l.InsertarUltimo(5)

	iter := l.Iterador()
	require.True(t, iter.HaySiguiente())
	t.Log("Insertar en la posición en la que se crea el iterador")
	iter.Insertar(1)
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, 1, l.VerPrimero())
	require.EqualValues(t, 1, iter.Siguiente())
	require.EqualValues(t, 2, iter.Siguiente())
	require.EqualValues(t, 4, iter.VerActual())

	t.Log("Insertar en la posición en el medio")
	iter.Insertar(3)
	require.EqualValues(t, 3, iter.VerActual())
	require.EqualValues(t, 3, iter.Siguiente())
	require.EqualValues(t, 4, iter.Siguiente())
	require.EqualValues(t, 5, iter.VerActual())
	require.EqualValues(t, 5, iter.Siguiente())

	t.Log("Insertar en la posición final")
	iter.Insertar(6)
	require.EqualValues(t, 6, iter.VerActual())
	require.EqualValues(t, 6, iter.Siguiente())
	require.False(t, iter.HaySiguiente())

	iter2 := l.Iterador()
	for i := 1; i <= 6; i++ {
		require.True(t, iter2.HaySiguiente())
		require.EqualValues(t, i, iter2.VerActual())
		iter2.Siguiente()
	}
}

// Test borrar posicion inicio, medio y fin
func TestIterBorrar(t *testing.T) {
	t.Log("Hacemos pruebas de iterador externo: borrar inicio, medio, fin")
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	l.InsertarUltimo(4)
	l.InsertarUltimo(5)
	require.False(t, l.EstaVacia())

	iter := l.Iterador()
	require.True(t, iter.HaySiguiente())

	t.Log("Borrar en la posición en la que se crea el iterador")
	require.EqualValues(t, 1, iter.Borrar())
	require.EqualValues(t, 2, iter.VerActual())
	require.EqualValues(t, 2, iter.Siguiente())
	require.EqualValues(t, 2, l.VerPrimero())

	t.Log("Borrar en el medio")
	require.EqualValues(t, 3, iter.Borrar())
	require.EqualValues(t, 4, iter.VerActual())
	require.EqualValues(t, 4, iter.Siguiente())

	t.Log("Borrar en la ultima posicion")
	require.EqualValues(t, 5, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.False(t, l.EstaVacia())
	require.EqualValues(t, 4, l.VerUltimo())
}

// Test pruebas de volumen con datos de tipo string
func TestVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen con datos de tipo string")
	l := TDALista.CrearListaEnlazada[string]()
	for i := 0; i < 1000; i++ {
		l.InsertarUltimo(strconv.Itoa(i))
	}
	iter := l.Iterador()
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, strconv.Itoa(i), iter.VerActual())
		require.EqualValues(t, strconv.Itoa(i), iter.Siguiente())
	}
	require.False(t, iter.HaySiguiente())
}

// Test pruebas de panics de iterador externo
func TestPanics(t *testing.T) {
	t.Log("Hacemos pruebas de panics de iterador externo")
	l := TDALista.CrearListaEnlazada[int]()

	iter := l.Iterador()
	//Prueba de panics al crear un iterador en una lista vacia o estar al final de la lista
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}

// Test invariante
func TestInvarianteLista(t *testing.T) {
	t.Log("Hacemos pruebas de invariante 'largo' de listaEnlazada")
	l := TDALista.CrearListaEnlazada[int]()

	// primitivas de lista
	l.InsertarUltimo(1)
	require.EqualValues(t, 1, l.Largo())
	l.InsertarPrimero(0)
	require.EqualValues(t, 2, l.Largo())

	// primitivas de iterador externo
	iter := l.Iterador()
	iter.Insertar(2)
	require.EqualValues(t, 3, l.Largo())
	iter.Insertar(3)
	require.EqualValues(t, 4, l.Largo())
	iter.Borrar()
	require.EqualValues(t, 3, l.Largo())
	iter.Borrar()
	iter.Borrar()
	require.EqualValues(t, 1, l.Largo())
	iter.Borrar()
	require.EqualValues(t, 0, l.Largo())
}

/* Test Iterador Interno */

// Test suma de elementos de una lista
func TestSuma(t *testing.T) {
	t.Log("Hacemos pruebas de iterador interno: suma de elementos de una lista")
	l := TDALista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	l.InsertarUltimo(4)
	l.InsertarUltimo(5)

	//suma de elementos sin condicion de corte
	suma_uno := 0
	suma_e := &suma_uno
	l.Iterar(func(e int) bool {
		*suma_e += e
		return true
	})
	require.EqualValues(t, 15, suma_uno)

	//suma de elementos con condicion de corte
	suma_dos := 0
	suma_e = &suma_dos
	l.Iterar(func(e int) bool {
		*suma_e += e
		return e != 3
	})
	require.EqualValues(t, 6, suma_dos)
}

// Test lista vacia con iterador externo
func TestListaVaciaDos(t *testing.T) {
	t.Log("Hacemos pruebas de iterador externo: lista vacia")
	l := TDALista.CrearListaEnlazada[int]()

	iter := l.Iterador()
	iter.Insertar(1)
	iter.Insertar(2)
	require.EqualValues(t, 2, iter.VerActual())
	require.False(t, l.EstaVacia())

	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, 1, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.True(t, l.EstaVacia())

	iter.Insertar(1)
	iter.Insertar(2)
	require.EqualValues(t, 2, iter.VerActual())
	require.False(t, l.EstaVacia())

	require.EqualValues(t, 2, iter.Borrar())
	require.EqualValues(t, 1, iter.VerActual())
	require.EqualValues(t, 1, iter.Borrar())
	require.False(t, iter.HaySiguiente())
	require.True(t, l.EstaVacia())
}

func TestCatedraErrores(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()
	iter := l.Iterador()
	iter.Insertar(1)
	iter.Siguiente()
	l.InsertarUltimo(5)
	iter2 := l.Iterador()
	require.EqualValues(t, 1, iter2.Siguiente())
	require.EqualValues(t, 5, iter2.Siguiente())
	require.EqualValues(t, 1, l.BorrarPrimero())
	require.EqualValues(t, 5, l.BorrarPrimero())

}

func TestCatedraErrores2(t *testing.T) {
	l := TDALista.CrearListaEnlazada[int]()

	l.InsertarUltimo(2)
	l.InsertarUltimo(4)
	l.InsertarUltimo(5)
	l.InsertarUltimo(6)
	l.InsertarUltimo(8)
	l.InsertarUltimo(9)
	iter2 := l.Iterador()
	require.EqualValues(t, 2, iter2.Siguiente())
	require.EqualValues(t, 4, iter2.Siguiente())
	require.EqualValues(t, 5, iter2.Siguiente())
	require.EqualValues(t, 6, iter2.Siguiente())
	iter2.Insertar(7)
	require.EqualValues(t, 7, iter2.Siguiente())
	require.EqualValues(t, 8, iter2.Siguiente())
	require.EqualValues(t, 9, iter2.Siguiente())
	iter3 := l.Iterador()
	require.EqualValues(t, 2, iter3.Siguiente())
	require.EqualValues(t, 4, iter3.Siguiente())
	require.EqualValues(t, 5, iter3.Siguiente())
	require.EqualValues(t, 6, iter3.Siguiente())
	require.EqualValues(t, 7, iter3.Siguiente())
	require.EqualValues(t, 8, iter3.Siguiente())
	require.EqualValues(t, 9, iter3.Siguiente())
	require.EqualValues(t, 2, l.BorrarPrimero())
	require.EqualValues(t, 4, l.BorrarPrimero())
	require.EqualValues(t, 5, l.BorrarPrimero())
	require.EqualValues(t, 6, l.BorrarPrimero())
	require.EqualValues(t, 7, l.BorrarPrimero())
	require.EqualValues(t, 8, l.BorrarPrimero())
	require.EqualValues(t, 9, l.BorrarPrimero())

}
