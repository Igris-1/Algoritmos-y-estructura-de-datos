package pila_test

import (
	TDAPila "go/pila/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

// Pruebas de pila vacia
func TestPilaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con pila vacia")
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
}

// Pruebas sin redimensionar con capacidad 10
func TestPilaInicial(t *testing.T) {
	t.Log("Hacemos pruebas sin redimensionar")
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	pila.Apilar(2)
	require.EqualValues(t, 2, pila.VerTope())
	pila.Apilar(3)
	require.EqualValues(t, 3, pila.VerTope())
	pila.Apilar(4)
	require.EqualValues(t, 4, pila.VerTope())
	pila.Apilar(5)
	require.EqualValues(t, 5, pila.VerTope())
	require.EqualValues(t, 5, pila.Desapilar())
	require.EqualValues(t, 4, pila.Desapilar())
	require.EqualValues(t, 3, pila.Desapilar())
	require.EqualValues(t, 2, pila.Desapilar())
	require.EqualValues(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

// Prueba de volumen:
// redimensiona la pila apilando y desapilando 1000 elementos de tipo int
func TestPilaVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen")
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 1000; i++ {
		pila.Apilar(i)
		require.EqualValues(t, i, pila.VerTope())
	}
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, i, pila.Desapilar())
	}
	require.True(t, pila.EstaVacia())
}

// Pruebas con strings
func TestPilaStrings(t *testing.T) {
	t.Log("Hacemos pruebas con strings")
	pila := TDAPila.CrearPilaDinamica[string]()
	pila.Apilar("A")
	pila.Apilar("B")
	require.EqualValues(t, "B", pila.VerTope())
	require.EqualValues(t, "B", pila.Desapilar())
	require.EqualValues(t, "A", pila.Desapilar())
	require.True(t, pila.EstaVacia())
}


// Pruebas de parcial

// ejemplos
// balanceado("[{()}]") -> true
// balanceado("[{()}") -> false
// balanceado("[{()}]{}") -> true
// balanceado("[{()}]{}{") -> false

func Testbalanceado(texto string) bool {
	pila := TDAPila.CrearPilaDinamica[string]()
	for _, c := range texto {
		if c == '{' || c == '[' || c == '(' {
			pila.Apilar(c)
		}
		if c == '}' || c == ']' || c == ')' {
			if pila.EstaVacia() {
				return false
			}
			if c == '}' && pila.VerTope() != '{' {
				return false
			}
			if c == ']' && pila.VerTope() != '[' {
				return False
			}
			if  c == ')' && pila.VerTope() != '(' {
				return false
			}
			pila.Desapilar()
		}
	}
	return pila.EstaVacia()
}

