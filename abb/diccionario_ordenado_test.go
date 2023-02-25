package diccionario_test

import (
	"diccionario"
	"fmt"
	"math/rand"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// funcion de comparacion de strings
func compararStrings(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

// funcion de comparacion de ints
func compararInts(a, b int) int { return a - b }

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

//{12500, 25000, 50000, 100000, 200000, 400000}

func TestDiccionarioVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := diccionario.CrearABB[string, int](compararStrings)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}

func TestGuardarYBorrar(t *testing.T) {
	t.Log("Guarda y borra elementos del diccionario, y comprueba que funciona acorde")
	claves := []string{"A", "B", "C", "D", "E", "F", "G", "H"}

	dic := diccionario.CrearABB[string, string](compararStrings)
	for _, clave := range claves {
		dic.Guardar(clave, clave)
		require.True(t, dic.Pertenece(clave))
	}

	for _, clave := range claves {
		require.EqualValues(t, clave, dic.Borrar(clave))
	}
}

func TestUnElement(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := diccionario.CrearABB[string, int](compararStrings)
	dic.Guardar("A", 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece("A"))
	require.False(t, dic.Pertenece("B"))
	require.EqualValues(t, 10, dic.Obtener("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("B") })
}

func TestDiccionarioGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := diccionario.CrearABB[string, string](compararStrings)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDato(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := "Gato"
	clave2 := "Perro"

	dic := diccionario.CrearABB[string, string](compararStrings)
	dic.Guardar(clave, "miau")
	dic.Guardar(clave2, "guau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, "miau", dic.Obtener(clave))
	require.EqualValues(t, "guau", dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, "miu")
	dic.Guardar(clave2, "baubau")
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, "miu", dic.Obtener(clave))
	require.EqualValues(t, "baubau", dic.Obtener(clave2))
}

func TestDiccionarioBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	valor1 := "miau"
	valor2 := "guau"
	valor3 := "moo"
	claves := []string{clave1, clave2, clave3}
	valores := []string{valor1, valor2, valor3}

	dic := diccionario.CrearABB[string, string](compararStrings)

	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: revisa, para el caso que fuere un HashCerrado, que no haya problema " +
		"reinsertando un elemento borrado")
	dic := diccionario.CrearABB[string, string](compararStrings)
	clave := "hola"
	dic.Guardar(clave, "mundo!")
	dic.Borrar(clave)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(clave))
	dic.Guardar(clave, "mundooo!")
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, "mundooo!", dic.Obtener(clave))
}

func TestConClavesNumericas(t *testing.T) {
	t.Log("Valida que no solo funcione con strings")
	dic := diccionario.CrearABB[int, string](compararInts)
	clave := 10
	valor := "Gatito"

	dic.Guardar(clave, valor)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, valor, dic.Obtener(clave))
	require.EqualValues(t, valor, dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

func TestClaveVacia(t *testing.T) {
	t.Log("Guardamos una clave vacía (i.e. \"\") y deberia funcionar sin problemas")
	dic := diccionario.CrearABB[string, string](compararStrings)
	clave := ""
	dic.Guardar(clave, clave)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, clave, dic.Obtener(clave))
}

func TestCadenaLargaParticular(t *testing.T) {
	t.Log("Se han visto casos problematicos al utilizar la funcion de hashing de K&R, por lo que " +
		"se agrega una prueba con dicha funcion de hashing y una cadena muy larga")
	// El caracter '~' es el de mayor valor en ASCII (126).
	claves := make([]string, 10)
	cadena := "%d~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~" +
		"~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~"
	dic := diccionario.CrearABB[string, string](compararStrings)
	valores := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J"}
	for i := 0; i < 10; i++ {
		claves[i] = fmt.Sprintf(cadena, i)
		dic.Guardar(claves[i], valores[i])
	}
	require.EqualValues(t, 10, dic.Cantidad())

	ok := true
	for i := 0; i < 10 && ok; i++ {
		ok = dic.Obtener(claves[i]) == valores[i]
	}

	require.True(t, ok, "Obtener clave larga funciona")
}

func TestValorNulo(t *testing.T) {
	t.Log("Probamos que el valor puede ser nil sin problemas")
	dic := diccionario.CrearABB[string, *int](compararStrings)
	clave := "Pez"
	dic.Guardar(clave, nil)
	require.True(t, dic.Pertenece(clave))
	require.EqualValues(t, 1, dic.Cantidad())
	require.EqualValues(t, (*int)(nil), dic.Obtener(clave))
	require.EqualValues(t, (*int)(nil), dic.Borrar(clave))
	require.False(t, dic.Pertenece(clave))
}

//  TEST ITERADORES

func TestIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	claves := []string{clave1, clave2, clave3}
	dic := diccionario.CrearABB[string, *int](compararStrings)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []string{"", "", ""}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave string, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
}

func TestIteradorInternoValores(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"
	clave4 := "Burrito"
	clave5 := "Hamster"

	dic := diccionario.CrearABB[string, int](compararStrings)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ string, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

// prueba de volumen
func ejecutarPruebaVolumen(b *testing.B, n int) {
	dic := diccionario.CrearABB[string, int](compararStrings)
	// generar arrelgo de claves del 1 al n
	claves := make([]string, n)
	valores := make([]int, n)

	// agregamos elementos a los arreglos de claves y valores
	for i := 0; i < n; i++ {
		claves[i] = strconv.Itoa(i)
		valores[i] = i
	}

	// desordenamos los arreglos
	for i := 0; i < n; i++ {
		j := rand.Intn(n)
		claves[i], claves[j] = claves[j], claves[i]
	}

	/* Inserta 'n' parejas en el arbol */
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionario(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumen(b, n)
			}
		})
	}
}

func TestIterarDiccionarioVacio(t *testing.T) {
	t.Log("Iterar sobre diccionario vacio es simplemente tenerlo al final")
	dic := diccionario.CrearABB[string, int](compararStrings)
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestPruebaIterarTrasBorrados(t *testing.T) {
	t.Log("Prueba de caja blanca: Esta prueba intenta verificar el comportamiento del hash abierto cuando " +
		"queda con listas vacías en su tabla. El iterador debería ignorar las listas vacías, avanzando hasta " +
		"encontrar un elemento real.")

	clave1 := "Gato"
	clave2 := "Perro"
	clave3 := "Vaca"

	dic := diccionario.CrearABB[string, string](compararStrings)
	dic.Guardar(clave1, "")
	dic.Guardar(clave2, "")
	dic.Guardar(clave3, "")
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, "A")
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, "A", v1)
	require.EqualValues(t, clave1, iter.Siguiente())
	require.False(t, iter.HaySiguiente())
}

func ejecutarPruebasVolumenIterador(b *testing.B, n int) {
	dic := diccionario.CrearABB[string, *int](compararStrings)

	claves := make([]string, n)
	valores := make([]int, n)
	rand.Seed(time.Now().UnixNano())

	/* Inserta 'n' parejas en el hash */
	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", rand.Uint64())
		valores[i] = i
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		if clave == "" {
			ok = false
			break
		}
		valor = v1
		if valor == nil {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIterador(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIterador(b, n)
			}
		})
	}
}

func TestIteradorInterno(t *testing.T) {
	t.Log("Probamos el iterador interno con y sin condicion de corte")
	dic := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(15, 15)
	dic.Guardar(6, 6)
	dic.Guardar(5, 5)
	dic.Guardar(13, 13)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(8, 8)

	var suma int
	dic.Iterar(func(clave int, valor int) bool {
		suma += valor
		return true
	})
	require.EqualValues(t, 63, suma)

	suma = 0
	dic.Iterar(func(clave int, valor int) bool {
		suma += valor
		return false
	})
	require.EqualValues(t, 5, suma)
}

func TestIteradorInternoCorte(t *testing.T) {
	t.Log("Probamos el iterador interno con una condicion de corte especifica")
	dic := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })

	arr := []int{50, 20, 45, 15, 10, 4, 8, 7, 9, 30}
	for _, v := range arr {
		dic.Guardar(v, v)
	}

	dic.Iterar(func(clave int, valor int) bool {
		if clave == 15 {
			return false
		}
		return true
	})
}

func TestIteradorInternoInorder(t *testing.T) {
	t.Log("Probamos el iterador interno con recorrido inorder")
	dic := diccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, v := range arr {
		dic.Guardar(i, v)
	}

	var i int
	dic.Iterar(func(clave int, valor string) bool {
		require.EqualValues(t, arr[i], valor)
		i++
		return true
	})
}

func TestIteradorInternoRango(t *testing.T) {
	t.Log("Probamos el iterador interno con rango con y sin condicion de corte")
	dic := diccionario.CrearABB[int, int](func(a, b int) int { return a - b })
	dic.Guardar(15, 15)
	dic.Guardar(6, 6)
	dic.Guardar(5, 5)
	dic.Guardar(13, 13)
	dic.Guardar(9, 9)
	dic.Guardar(7, 7)
	dic.Guardar(8, 8)

	desde := 6
	hasta := 13

	var suma int
	dic.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		suma += valor
		return true
	})
	require.EqualValues(t, 43, suma)

	suma = 0
	dic.IterarRango(&desde, &hasta, func(clave int, valor int) bool {
		suma += valor
		return false
	})
	require.EqualValues(t, 6, suma)
}

func TestIteradorInternoRangoInorder(t *testing.T) {
	t.Log("Probamos el iterador interno con rango con recorrido inorder")
	dic := diccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, v := range arr {
		dic.Guardar(i, v)
	}

	desde := 0
	hasta := 10

	var i int
	dic.IterarRango(&desde, &hasta, func(clave int, valor string) bool {
		require.EqualValues(t, arr[clave], valor)
		i++
		return true
	})
}

func TestIteradorExterno(t *testing.T) {
	t.Log("Probamos el iterador externo con una suma de elementos")
	dic := diccionario.CrearABB[int, int](compararInts)
	dic.Guardar(14, 14)
	dic.Guardar(4, 4)
	dic.Guardar(3, 3)
	dic.Guardar(9, 9)
	dic.Guardar(22, 22)
	dic.Guardar(16, 16)
	dic.Guardar(24, 24)

	iter := dic.Iterador()
	var suma int
	for iter.HaySiguiente() {
		suma += iter.Siguiente()
	}
	require.EqualValues(t, 92, suma)
}

// test iterador externo recorrido inorder
func TestIteradorExternoInorder(t *testing.T) {
	t.Log("Probamos el iterador externo con recorrido inorder")
	dic := diccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, v := range arr {
		dic.Guardar(i, v)
	}
	dic.Iterador()
	var i int
	iter := dic.Iterador()
	for iter.HaySiguiente() {
		require.EqualValues(t, arr[i], dic.Obtener(iter.Siguiente()))
		i++
	}
}

func TestIteradorExternoRango(t *testing.T) {
	t.Log("Probamos el iterador externo con rango con una suma de elementos")
	dic := diccionario.CrearABB[int, int](compararInts)
	dic.Guardar(14, 14)
	dic.Guardar(4, 4)
	dic.Guardar(3, 3)
	dic.Guardar(9, 9)
	dic.Guardar(22, 22)
	dic.Guardar(16, 16)
	dic.Guardar(24, 24)

	desde := 4
	hasta := 22

	iter := dic.IteradorRango(&desde, &hasta)
	var suma int
	for iter.HaySiguiente() {
		suma += iter.Siguiente()
	}
	require.EqualValues(t, 65, suma)
}

// test iterador externo con rango recorrido inorder
func TestIteradorExternoRangoInorder(t *testing.T) {
	t.Log("Probamos el iterador externo con rango con recorrido inorder")
	dic := diccionario.CrearABB[int, string](func(a, b int) int { return a - b })
	arr := []string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	for i, v := range arr {
		dic.Guardar(i, v)
	}

	desde := 0
	hasta := 20

	var i int
	iter := dic.IteradorRango(&desde, &hasta)
	for iter.HaySiguiente() {
		require.EqualValues(t, arr[i], dic.Obtener(iter.Siguiente()))
		i++
	}
}

func TestDeLaDiscordia(t *testing.T) {
	t.Log("Probamos el iterador interno con condicion de corte, cortesia de la cátedra")
	dic := diccionario.CrearABB[string, int](compararStrings)
	dic.Guardar("d", 0)
	dic.Guardar("c", 1)
	dic.Guardar("b", 2)
	dic.Guardar("a", 3)
	dic.Guardar("e", 4)

	arr := [3]string{}
	dic.Iterar(func(clave string, valor int) bool {
		if clave == "d" {
			return false
		}
		arr[valor-1] = clave
		return true
	})
	require.EqualValues(t, "c", arr[0])
	require.EqualValues(t, "b", arr[1])
	require.EqualValues(t, "a", arr[2])
}
