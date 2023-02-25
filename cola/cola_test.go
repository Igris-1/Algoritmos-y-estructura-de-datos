package cola_test

import (
	TDACola "cola"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"
)

// test cola vacia
func TestColaVacia(t *testing.T) {
	t.Log("Hacemos pruebas con cola vacia")
	c := TDACola.CrearColaEnlazada[string]()
	require.True(t, c.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.VerPrimero() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.Desencolar() })
}

// test de primitivas con un elemento
func TestColaPrimitivas(t *testing.T) {
	t.Log("Hacemos pruebas de primitivas")
	c := TDACola.CrearColaEnlazada[string]()
	c.Encolar("a")
	require.EqualValues(t, "a", c.VerPrimero())
	require.False(t, c.EstaVacia())
	require.EqualValues(t, "a", c.Desencolar())
	require.True(t, c.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.VerPrimero() })
}

// test cola con elementos de tipo string
func TestColaInicial(t *testing.T) {
	t.Log("Hacemos pruebas con elementos de tipo string")
	c := TDACola.CrearColaEnlazada[string]()
	c.Encolar("a")
	require.EqualValues(t, "a", c.VerPrimero())
	c.Encolar("b")
	require.EqualValues(t, "a", c.VerPrimero())
	c.Encolar("c")
	require.EqualValues(t, "a", c.VerPrimero())
	require.EqualValues(t, "a", c.Desencolar())
	require.EqualValues(t, "b", c.Desencolar())
	require.EqualValues(t, "c", c.Desencolar())
	require.True(t, c.EstaVacia())
}

// test de volumen encolando / desencolando 1000 elementos de tipo string
func TestColaVolumen(t *testing.T) {
	t.Log("Hacemos pruebas de volumen con elementos de tipo string")
	c := TDACola.CrearColaEnlazada[string]()
	for i := 0; i < 1000; i++ {
		c.Encolar(strconv.Itoa(i))
	}
	for i := 0; i < 1000; i++ {
		require.EqualValues(t, strconv.Itoa(i), c.Desencolar())
	}
	require.True(t, c.EstaVacia())
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { c.VerPrimero() })
}

// test con elementos de tipo int
func TestColaInt(t *testing.T) {
	t.Log("Hacemos pruebas con elementos de tipo int")
	c := TDACola.CrearColaEnlazada[int]()
	c.Encolar(1)
	require.EqualValues(t, 1, c.VerPrimero())
	c.Encolar(2)
	require.EqualValues(t, 1, c.VerPrimero())
	c.Encolar(3)
	require.EqualValues(t, 1, c.VerPrimero())
	require.EqualValues(t, 1, c.Desencolar())
	require.EqualValues(t, 2, c.Desencolar())
	require.EqualValues(t, 3, c.Desencolar())
	require.True(t, c.EstaVacia())
}
