package comandos

import (
	"fmt"
	"rerepolez/cola"
	"rerepolez/errores"
	"rerepolez/funciones"
	"rerepolez/funciones2"
	"rerepolez/votos"
	"strconv"
)

func IngresarVotante(lista_entrada []string, padrones []votos.Votante, cola_voto cola.Cola[int]) {
	DNI := funciones.DniValido(lista_entrada[funciones.DNI])
	dni, _ := strconv.Atoi(lista_entrada[funciones.DNI])

	votante := funciones2.BusquedaBinaria(padrones, DNI, 0, len(padrones)-1)
	if dni <= 0 {
		fmt.Println(errores.DNIError{})
	} else if votante == funciones2.NO_ENCONTRADO {
		fmt.Println(errores.DNIFueraPadron{})
	} else {
		cola_voto.Encolar(votante)
		fmt.Println("OK")
	}
}

func Votar(lista_entrada []string, padrones []votos.Votante, partidos []string, partidosCreados []votos.Partido, cola_voto cola.Cola[int]) {
	tipo_voto := funciones.TipoVoto(lista_entrada[funciones.TIPO_VOTO])
	alternativa, err := strconv.Atoi(lista_entrada[funciones.ALTERNATIVA])
	partidoEsta := funciones.PartidoValido(partidos, alternativa)

	if cola_voto.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else if funciones.VOTO_INVALIDO == tipo_voto {
		fmt.Println(errores.ErrorTipoVoto{})
	} else if err != nil || alternativa != funciones.VOTO_IMPUGNADO && !partidoEsta {
		fmt.Println(errores.ErrorAlternativaInvalida{})
	} else {
		votar := (padrones[cola_voto.VerPrimero()]).Votar(tipo_voto, alternativa)
		if votar == nil {
			fmt.Println("OK")
		} else {
			fmt.Println(votar)
			cola_voto.Desencolar()
		}
	}
}

func Deshacer(cola_voto cola.Cola[int], padrones []votos.Votante) {
	if cola_voto.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		test := padrones[cola_voto.VerPrimero()].Deshacer()
		if test == nil {
			fmt.Println("OK")
		} else if (test == errores.ErrorNoHayVotosAnteriores{}) {
			fmt.Println(test)
		} else {
			fmt.Println(test)
			cola_voto.Desencolar()
		}
	}
}

func FinVotar(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante) {
	if cola_voto.EstaVacia() {
		fmt.Println(errores.FilaVacia{})
	} else {
		votante := padrones[cola_voto.Desencolar()]
		datos, err := votante.FinVoto()
		if err != nil {
			fmt.Println(err)
		} else {
			if datos.Impugnado {
				partidosCreados[0].VotadoPara(votos.TipoVoto(votos.POS_LISTA_IMPUGNA))
				fmt.Println("OK")
			} else {
				fmt.Println("OK")
				for i := 0; i < 3; i++ {
					partidosCreados[datos.VotoPorTipo[i]].VotadoPara(votos.TipoVoto(i))
				}
			}
		}
	}
}

func Fin(partidosCreados []votos.Partido, cola_voto cola.Cola[int], padrones []votos.Votante) {
	if !cola_voto.EstaVacia() {
		fmt.Println(errores.ErrorCiudadanosSinVotar{})
	}
	for i := 0; i < 3; i++ {
		tipo_voto := votos.TipoVoto(i)
		fmt.Println(funciones.VotoTipo(tipo_voto) + ":")
		for j := 0; j < len(partidosCreados); j++ {
			fmt.Println(partidosCreados[j].ObtenerResultado(tipo_voto))
		}
		fmt.Println()
	}
	fmt.Println("Votos Impugnados:", partidosCreados[0].ObtenerResultado(votos.TipoVoto(votos.POS_LISTA_IMPUGNA)))
}
