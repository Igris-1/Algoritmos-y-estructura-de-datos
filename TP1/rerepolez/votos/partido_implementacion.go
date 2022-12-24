package votos

import (
	"strconv"
)

// creamos el tipo Partido, toma el nombre del partido, los candidatos y las votaciones por candidato
type partidoImplementacion struct {
	nombrePartido string
	candidatos    [CANT_VOTACION]string
	votaciones    [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosBlancos [CANT_VOTACION]int
	impugnados   int
}

// devuelve la implementacion de partido con el nombre del partido, los candidatos y las votaciones
func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	return &partidoImplementacion{nombre, candidatos, [CANT_VOTACION]int{}}
}

// devuelve la implementacion de votos en blanco
func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

// dependiendo del tipo de voto, suma un voto al candidato correspondiente
func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votaciones[tipo]++
}

// devuelve el nombre del partido, el nombre del candidato y la cantidad de votos
func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	if partido.votaciones[tipo] == 1 {
		return partido.nombrePartido + " - " + partido.candidatos[tipo] + ": " + strconv.Itoa(partido.votaciones[tipo]) + " voto"
	}
	return partido.nombrePartido + " - " + partido.candidatos[tipo] + ": " + strconv.Itoa(partido.votaciones[tipo]) + " votos"
}

// aca llamamos a la funcion VotadoPara(tipo TipoVoto) y dependiendo del tipo de voto,
// le va a sumar 1 al voto en blanco, ya sea presidente, gobernador o intendente
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	if tipo == TipoVoto(POS_LISTA_IMPUGNA) {
		blanco.impugnados++
	} else {
		blanco.votosBlancos[tipo]++
	}
}

// devuelve el total de votos en blanco
func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	if tipo == POS_LISTA_IMPUGNA {
		return blanco.ObtenerVotosImpugnados()
	}
	if blanco.votosBlancos[tipo] == 1 {
		return "Votos en Blanco: " + strconv.Itoa(blanco.votosBlancos[tipo]) + " voto"
	}
	return "Votos en Blanco: " + strconv.Itoa(blanco.votosBlancos[tipo]) + " votos"
}

// devuelve el total de votos impugnados
func (blanco partidoEnBlanco) ObtenerVotosImpugnados() string {
	if blanco.impugnados == 1 {
		return strconv.Itoa(blanco.impugnados) + " voto"
	}
	return strconv.Itoa(blanco.impugnados) + " votos"
}
