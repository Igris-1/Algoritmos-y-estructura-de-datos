package funciones

import (
	"bufio"
	"os"
	"rerepolez/votos"
	"strings"
)

// ARCHIVO_LISTA:
// partido, nombre del partido, nombre presidente, nombre gobernador, nombre intendente
// lectura de archivo de partidos
const (
	COMANDO        = 0
	DNI            = 1
	TIPO_VOTO      = 1
	ALTERNATIVA    = 2
	VOTO_INVALIDO  = 3
	VOTO_IMPUGNADO = 0
	ESTANDAR       = 8
)

// me fijo si el partido esta
func PartidoValido(partidos []string, numeroLista int) bool {
	return numeroLista <= len(partidos)
}

func CrearPartidos(partidos []string) []votos.Partido {
	var partidosCreados []votos.Partido
	partido_Blanco := votos.CrearVotosEnBlanco()
	partidosCreados = append(partidosCreados, partido_Blanco)
	for _, partido := range partidos {
		partido_lista := strings.Split(partido, ",")
		integrantes := [3]string{partido_lista[1], partido_lista[2], partido_lista[3]}
		partidosCreados = append(partidosCreados, votos.CrearPartido(partido_lista[0], integrantes))
	}
	return partidosCreados
}

func LecturaPartidos(archivo string) []string {
	var partidos []string

	archivoPartidos, err := os.Open(archivo)
	if err != nil {
		return nil
	}
	defer archivoPartidos.Close()

	partido := bufio.NewScanner(archivoPartidos)
	for partido.Scan() {
		partidos = append(partidos, partido.Text())
	}
	return partidos
}

// lectura de archivo de padrones
func LecturaPadrones(archivo string) []votos.Votante {
	var padrones []votos.Votante

	archivoPadrones, err := os.Open(archivo)
	if err != nil {
		return nil
	}
	defer archivoPadrones.Close()

	dni := bufio.NewScanner(archivoPadrones)
	for dni.Scan() {
		var votante votos.Votante
		if len(dni.Text()) < ESTANDAR {
			votante = votos.CrearVotante("0" + dni.Text())
		} else {
			votante = votos.CrearVotante(dni.Text())
		}
		padrones = append(padrones, votante)
	}
	return padrones
}

func DniValido(dni string) string {
	if len(dni) < ESTANDAR {
		dni = "0" + dni
	}
	return dni
}

func TipoVoto(tipo_voto string) votos.TipoVoto {
	switch tipo_voto {
	case "Presidente":
		return votos.PRESIDENTE
	case "Gobernador":
		return votos.GOBERNADOR
	case "Intendente":
		return votos.INTENDENTE
	default:
		return VOTO_INVALIDO
	}
}

func VotoTipo(tipo_voto votos.TipoVoto) string {
	switch tipo_voto {
	case votos.PRESIDENTE:
		return "Presidente"
	case votos.GOBERNADOR:
		return "Gobernador"
	case votos.INTENDENTE:
		return "Intendente"
	default:
		return "Relleno"
	}
}
