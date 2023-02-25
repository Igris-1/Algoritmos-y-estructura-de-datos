package administrador

/**
El objetivo de este módulo es la de SIMULAR el pedido de memoria explícito del heap al Sistema Operativo.
Esto es para tener una idea inicial de cómo sucede esto por debajo en el lenguaje, y aprender cómo es
la operación en otros lenguajes de bajo nivel (como C).
Esta librería ÚNICAMENTE se utilizará para esta entrega particular. Para las siguientes entregas trabajaremos
directamente con el recolector de basura (Garbage Collector) del lenguaje, usando las funciones make, new, etc...
Sin mayores problemas. Esta entrega está pensada únicamente para practicar un poco el tema de memoria dinámica
(hasta tanto salga el plan nuevo) y empezar a ver un poco de manejo de estructuras en Go.

Para quien le interese, dejamos una equivalencia de estas funciones a las de C:
- PedirMemoria y PedirArreglo: malloc
- RedimensionarMemoria: realloc
- LiberarMemoria y LiberarArreglo: free
- Finalizar: no hay (debemos usar otros programas adicionales que hacen algo similar a lo que planteamos en
esta librería, como es Valgrind).
*/

import (
	"fmt"
	"runtime/debug"
	"unsafe"
)

type info struct {
	previo *info
	stack  []byte
	len    int
}

var _pedidos = make(map[any]*info)

// PedirMemoria nos simula el pedido de memoria del heap que se le hace al sistema operativo para almacenar
// un tipo de dato dado. Lo que se pida con esta función debe liberarse con LiberarMemoria. Si no se libera,
// fallará luego al invocarse Finalizar
func PedirMemoria[T any]() *T {
	mem := new(T)
	_pedidos[mem] = &info{stack: debug.Stack(), len: int(unsafe.Sizeof(mem)), previo: nil}
	return mem
}

// PedirArreglo es símil a PedirMemoria pero en vez de pedir memoria del heap para un dato, lo hace para un arreglo
// del tipo indicado. Esta memoria puede luego redimensionarse con RedimensionarMemoria, y eventualmente debe ser
// liberada con LiberarArreglo. Si no se libera, fallará luego al invocarse Finalizar.
func PedirArreglo[T any](n int) *[]T {
	mem := make([]T, n)
	_pedidos[&mem] = &info{stack: debug.Stack(), len: n * int(unsafe.Sizeof(new(T))), previo: nil}
	return &mem
}

// LiberarMemoria libera la memoria del heap pedida
func LiberarMemoria[T any](dato *T) {
	_, ok := _pedidos[dato]
	if !ok {
		panic("Liberando memoria no pedida o ya liberada!")
	}
	delete(_pedidos, dato)
}

// LiberarArreglo (como LiberarMemoria) libera la memoria del heap pedida
func LiberarArreglo[T any](datos *[]T) {
	_, ok := _pedidos[datos]
	if !ok {
		panic("Liberando memoria no pedida o ya liberada!")
	}
	delete(_pedidos, datos)
}

// RedimensionarMemoria permite cambiar el tamaño de la memoria pedida anteriormente. Puede ser más grande o más chica.
// Se copian todos los valores que entren de los datos anteriores al arreglo actual.
func RedimensionarMemoria[T any](datos *[]T, nuevoTam int) *[]T {
	ptr := PedirArreglo[T](nuevoTam)
	copy(*ptr, *datos)
	_pedidos[ptr].previo = _pedidos[&datos]
	LiberarArreglo(datos)
	return ptr
}

func (i info) imprimirInfo(org bool) {
	if org {
		fmt.Printf("Se pidio memoria para %d bytes en: \n", i.len)
	}
	fmt.Print(string(i.stack))
	if i.previo != nil {
		fmt.Println("Habiéndose redimensionado del pedido original:")
		i.previo.imprimirInfo(false)
	}
}

// Finalizar chequea que no haya quedado ninguna porción de memoria pedida sin haberse liberado. Si hay memoria
// sin liberar, mostrará un mensaje con el stacktrace del camino donde se pidió exactamente esa memoria, a fin
// de debuggear y poder encontrar una solución
func Finalizar() {
	if len(_pedidos) == 0 {
		return
	}

	fmt.Println("Error: Memoria no liberada")
	for _, info := range _pedidos {
		info.imprimirInfo(true)
		fmt.Println("")
	}
	// Limpiamos para no afectar a otros tests
	_pedidos = make(map[any]*info)
	panic("No se libero toda la memoria")
}
