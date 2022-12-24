package main

import (
	"bufio"
	"fmt"
	"os"
	"rerepolez/cola"
	"rerepolez/comandos"
	"rerepolez/errores"
	"rerepolez/funciones"
	"rerepolez/funciones2"
	"strings"
)

// ARCHIVO_LISTA:
// partido, nombre del partido, nombre presidente, nombre gobernador, nombre intendente

// ARCHIVO_PADRON: contiene un DNI por linea

func main() {
	//lo primero en probar son los archivos, si rompen cortamos toda la ejecucion
	argumentos := os.Args[1:]

	if len(argumentos) != 2 {
		fmt.Println(errores.ErrorParametros{})
		return
	}

	// la url de los archivos archivos es con / no con \
	partidos := funciones.LecturaPartidos(argumentos[0])
	padrones := funciones.LecturaPadrones(argumentos[1])
	if partidos == nil || padrones == nil {
		fmt.Println(errores.ErrorLeerArchivo{})
		return
	}

	// arreglo de TDAs de partidos
	partidosCreados := funciones.CrearPartidos(partidos)

	// padrones ordenados
	padrones = funciones2.RadixSort(padrones)

	// cola de votantes(orden en el arreglo)
	cola_voto := cola.CrearColaEnlazada[int]()

	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		lista_entrada := strings.Split(entrada.Text(), " ")

		switch lista_entrada[funciones.COMANDO] {
		case "ingresar":
			// ingreso de votantes
			comandos.IngresarVotante(lista_entrada, padrones, cola_voto)

		case "votar":
			// votacion
			comandos.Votar(lista_entrada, padrones, partidos, partidosCreados, cola_voto)

		case "deshacer":
			// deshacer votacion
			comandos.Deshacer(cola_voto, padrones)

		case "fin-votar":
			// fin de votacion
			comandos.FinVotar(partidosCreados, cola_voto, padrones)
		}
	}

	// fin de ejecucion
	comandos.Fin(partidosCreados, cola_voto, padrones)
}
