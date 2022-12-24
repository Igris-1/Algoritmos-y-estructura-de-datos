package funciones2

import (
	"rerepolez/votos"
	"strconv"
)

const (
	_DESDE        = 0
	_HASTA        = 9
	NO_ENCONTRADO = -1
)

func RadixSort(padrones []votos.Votante) []votos.Votante {
	for i := 7; i >= 0; i-- {
		padrones = countingSort(padrones, _DESDE, _HASTA, i, i+1)
	}
	return padrones
}

// Counting sort
func countingSort(padrones []votos.Votante, inicio, fin, indiceI, indiceF int) []votos.Votante {
	// suma de frecuencias
	frecuencias := make([]int, fin-inicio+1)

	// sumas acumuladas
	sumasAcumuladas := make([]int, fin-inicio+1)

	// arreglo final
	arr := make([]votos.Votante, len(padrones))

	// suma de frecuencias
	for _, votante := range padrones {
		elementoInt, _ := strconv.Atoi(votante.LeerDNI()[indiceI:indiceF])
		frecuencias[elementoInt-inicio]++
	}

	// sumas acumuladas
	for i := 1; i < len(frecuencias); i++ {
		sumasAcumuladas[i] = sumasAcumuladas[i-1] + frecuencias[i-1]
	}

	// arreglo final
	for _, votante := range padrones {
		elementoInt, _ := strconv.Atoi(votante.LeerDNI()[indiceI:indiceF])
		arr[sumasAcumuladas[elementoInt-inicio]] = votante
		sumasAcumuladas[elementoInt-inicio]++
	}

	// devolveer el arreglo final
	return arr
}

// busqueda binaria
func BusquedaBinaria(padrones []votos.Votante, dni string, inicio, fin int) int {
	if inicio > fin {
		return -1
	}

	medio := (inicio + fin) / 2

	if padrones[medio].LeerDNI() == dni {
		return medio
	}

	if padrones[medio].LeerDNI() > dni {
		return BusquedaBinaria(padrones, dni, inicio, medio-1)
	} else {
		return BusquedaBinaria(padrones, dni, medio+1, fin)
	}
}
