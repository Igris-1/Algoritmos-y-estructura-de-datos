package cola_prioridad_test

import (
	heap "cola_prioridad"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

func compararStrings(a, b string) int {
	if a == b {
		return 0
	}
	if a < b {
		return -1
	}
	return 1
}

func compararInts(a, b int) int { return a - b }

func TestCrearHeap(t *testing.T) {
	t.Log("Test Crear Heap")
	heap := heap.CrearHeap(compararStrings)
	require.NotNil(t, heap)
}

func TestHeapVacio(t *testing.T) {
	t.Log("Al crear un Heap debe estar vacio")
	heap := heap.CrearHeap(compararStrings)
	require.True(t, heap.EstaVacia())
}

func TestPanicHeapVacio(t *testing.T) {
	t.Log("Ver el maximo/desencolar un Heap vacio debe entrar en pánico")
	heap := heap.CrearHeap(compararStrings)
	require.Panics(t, func() { heap.VerMax() })
	require.Panics(t, func() { heap.Desencolar() })
}

func TestHeapNoVacio(t *testing.T) {
	t.Log("Al agregar elementos al Heap no debe estar vacio")
	heap := heap.CrearHeap(compararStrings)
	heap.Encolar("a")
	require.False(t, heap.EstaVacia())
	heap.Encolar("b")
	require.False(t, heap.EstaVacia())
	heap.Encolar("c")
	require.False(t, heap.EstaVacia())

	require.Equal(t, "c", heap.VerMax())
	require.Equal(t, "c", heap.Desencolar())
	require.Equal(t, "b", heap.Desencolar())
	require.Equal(t, "a", heap.Desencolar())
}

func TestHeapUnElemento(t *testing.T) {
	t.Log("Al agregar un elemento al Heap debe ser el maximo")
	heap := heap.CrearHeap(compararStrings)
	heap.Encolar("a")
	require.Equal(t, "a", heap.VerMax())
}

func TestHeapVariosElementos(t *testing.T) {
	t.Log("Al agregar varios elementos al Heap deben desencolarse de mayor a menor")
	heap := heap.CrearHeap(compararStrings)
	arr := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, v := range arr {
		heap.Encolar(v)
	}

	for i := 0; i < len(arr); i++ {
		require.False(t, heap.EstaVacia())
		require.Equal(t, arr[len(arr)-i-1], heap.Desencolar())
		require.Equal(t, len(arr)-i-1, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapInts(t *testing.T) {
	t.Log("Al agregar varios elementos al Heap deben desencolarse de mayor a menor")
	heap := heap.CrearHeap(compararInts)
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, v := range arr {
		heap.Encolar(v)
	}

	for i := 0; i < len(arr); i++ {
		require.False(t, heap.EstaVacia())
		require.Equal(t, arr[len(arr)-i-1], heap.Desencolar())
		require.Equal(t, len(arr)-i-1, heap.Cantidad())
	}
	require.True(t, heap.EstaVacia())
}

func TestEncolarStringsRandoms(t *testing.T) {
	t.Log("Test Encolar Strings Randoms")
	heap := heap.CrearHeap(compararStrings)
	arr := []string{"a", "b", "c", "h", "c", "f", "g", "h", "k", "j", "k", "l", "m", "z", "o", "l", "y", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	for _, v := range arr {
		heap.Encolar(v)
	}
	require.Equal(t, "z", heap.VerMax())
	require.False(t, heap.EstaVacia())
	require.Equal(t, "z", heap.Desencolar())
	require.Equal(t, "z", heap.Desencolar())
	require.Equal(t, "y", heap.Desencolar())
	require.Equal(t, "y", heap.Desencolar())
	require.Equal(t, "x", heap.Desencolar())

	for i := 0; i < len(arr)-5; i++ {
		require.False(t, heap.EstaVacia())
		heap.Desencolar()
	}
	require.True(t, heap.EstaVacia())
}

func TestEncolarRandoms(t *testing.T) {
	t.Log("Agregramos elementos randoms al heap y verificamos que se desencolan en orden de maximos")
	heap := heap.CrearHeap(compararInts)
	arr := []int{40, 20, 10, 30, 50, 60, 70, 100, 90, 100}
	for _, v := range arr {
		heap.Encolar(v)
		require.False(t, heap.EstaVacia())
	}
	require.EqualValues(t, 10, heap.Cantidad())
	require.EqualValues(t, 100, heap.VerMax())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 100, heap.Desencolar())
	require.EqualValues(t, 90, heap.Desencolar())
	require.EqualValues(t, 70, heap.Desencolar())
	require.EqualValues(t, 60, heap.Desencolar())
	require.EqualValues(t, 50, heap.Desencolar())
	require.EqualValues(t, 40, heap.Desencolar())
	require.EqualValues(t, 30, heap.Desencolar())
	require.EqualValues(t, 20, heap.Desencolar())
	require.EqualValues(t, 10, heap.Desencolar())
	require.True(t, heap.EstaVacia())
}

func TestHeapVolumen(t *testing.T) {
	t.Log("Test Volumen")
	heap := heap.CrearHeap(compararInts)
	for i := 0; i < 100000; i++ {
		heap.Encolar(i)
	}
	for i := 0; i < 100000; i++ {
		require.EqualValues(t, 100000-i, heap.Cantidad())
		require.EqualValues(t, 99999-i, heap.Desencolar())
	}
	require.True(t, heap.EstaVacia())
}

func TestHeapSort(t *testing.T) {
	t.Log("Test Heap Sort")
	arr := []int{40, 20, 10, 30, 50, 60, 70, 100, 90, 100}
	heap.HeapSort(arr, compararInts)
	require.EqualValues(t, 10, len(arr))
	require.EqualValues(t, 10, arr[0])
	require.EqualValues(t, 20, arr[1])
	require.EqualValues(t, 30, arr[2])
	require.EqualValues(t, 40, arr[3])
	require.EqualValues(t, 50, arr[4])
	require.EqualValues(t, 60, arr[5])
	require.EqualValues(t, 70, arr[6])
	require.EqualValues(t, 90, arr[7])
	require.EqualValues(t, 100, arr[8])
	require.EqualValues(t, 100, arr[9])
}

func TestPanicHeapSort(t *testing.T) {
	t.Log("Test Heap Sort")
	arr := []int{40, 20, 10, 30, 50, 60, 70, 100, 90, 100}
	require.Panics(t, func() { heap.HeapSort(arr, nil) })
}

func TestHeapSortStrings(t *testing.T) {
	t.Log("Test HeapSort Strings")

	arr1 := []string{"α", "β", "γ", "δ", "ε", "ζ", "η", "θ", "ι", "κ", "λ", "μ", "ν", "ξ", "ο", "π", "ρ", "σ", "τ", "υ", "φ", "χ", "ψ", "ω"}
	arr2 := []string{"α", "β", "γ", "δ", "ε", "ζ", "η", "θ", "ι", "κ", "λ", "μ", "ν", "ξ", "ο", "π", "ρ", "σ", "τ", "υ", "φ", "χ", "ψ", "ω"}

	for i := 0; i < len(arr2)/2; i++ {
		arr2[i], arr2[len(arr2)-i-1] = arr2[len(arr2)-i-1], arr2[i]
	}

	heap.HeapSort(arr2, compararStrings)
	require.EqualValues(t, 24, len(arr2))

	for i := 0; i < len(arr1); i++ {
		require.EqualValues(t, arr1[i], arr2[i])
	}
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Test Heap Sort Volumen")
	arr := []int{}
	for i := 0; i < 1000000; i++ {
		arr = append(arr, rand.Intn(1000000))
	}
	heap.HeapSort(arr, compararInts)
	for i := 0; i < len(arr)-1; i++ {
		require.True(t, arr[i] <= arr[i+1])
	}
}
