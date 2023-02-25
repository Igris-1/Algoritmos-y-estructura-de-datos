package admmemoria_test

import (
	Vector "administracionmemoria"
	"administracionmemoria/administrador"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestVectorNulo(t *testing.T) {
	t.Log("Hacemos pruebas con vector de tamaño 0")
	vec := Vector.CrearVector(0)
	require.NotNil(t, vec)
	require.Zero(t, vec.Largo())
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(0, 5) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(1, 10) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(15, 0) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Obtener(0) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Obtener(1) })
	vec.Destruir()
	administrador.Finalizar()
}

func TestVectorAlgunosElementos(t *testing.T) {
	t.Log("Hacemos pruebas con algunos elementos, algunos en posiciones invalidas")
	vec := Vector.CrearVector(5)
	require.EqualValues(t, 5, vec.Largo())
	vec.Guardar(0, 20)
	vec.Guardar(1, 30)
	vec.Guardar(0, 15)
	vec.Guardar(4, 7)
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(5, 35) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(6, 40) })
	require.EqualValues(t, 15, vec.Obtener(0))
	require.EqualValues(t, 30, vec.Obtener(1))
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Obtener(5) })
	vec.Destruir()
	administrador.Finalizar()
}

func TestRedimension(t *testing.T) {
	vec := Vector.CrearVector(3)
	vec.Guardar(0, 5)
	vec.Guardar(1, 10)

	// Agrandamos
	vec.Redimensionar(10)
	require.EqualValues(t, 10, vec.Largo())
	require.EqualValues(t, 5, vec.Obtener(0))
	require.EqualValues(t, 10, vec.Obtener(1))
	vec.Guardar(5, 50)
	require.EqualValues(t, 50, vec.Obtener(5))
	// Ahora achicamos
	vec.Redimensionar(2)
	require.EqualValues(t, 2, vec.Largo())
	require.EqualValues(t, 5, vec.Obtener(0))
	require.EqualValues(t, 10, vec.Obtener(1))
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Obtener(2) })
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(5, 0) })
	vec.Destruir()
	administrador.Finalizar()
}

func TestVolumen(t *testing.T) {
	tam := 10000
	vec := Vector.CrearVector(tam)
	require.EqualValues(t, tam, vec.Largo())
	for i := 0; i < tam; i++ {
		vec.Guardar(i, i)
	}
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Guardar(tam, tam) })

	for i := 0; i < tam; i++ {
		require.EqualValues(t, i, vec.Obtener(i))
	}
	require.PanicsWithValue(t, "Fuera de rango", func() { vec.Obtener(tam) })

	vec.Destruir()
	administrador.Finalizar()
}
