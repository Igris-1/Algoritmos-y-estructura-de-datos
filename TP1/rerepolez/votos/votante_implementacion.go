package votos

import (
	"rerepolez/errores"
	"rerepolez/pila"
)

type votanteImplementacion struct {
	dni            string
	voto           Voto
	yaVoto         bool
	historial_voto pila.Pila[Voto]
}

// Se está apilando el voto el blanco, toda persona que entre al cuarto oscuro y
// se le aplique fin-votar va a votar en blanco, osea es el voto base que tendría
// una persona si se le aplica el comando deshacer hasta que no haya más votos
func CrearVotante(dni string) Votante {
	var votante votanteImplementacion
	votante.dni = dni
	votante.historial_voto = pila.CrearPilaDinamica[Voto]()
	votante.historial_voto.Apilar(votante.voto)
	return &votante
}

func (votante votanteImplementacion) LeerDNI() string {
	return votante.dni
}

func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.yaVoto {
		if votante.LeerDNI()[0:1] == "0" {
			return errores.ErrorVotanteFraudulento{Dni: votante.dni[1:]}
		} else {
			return errores.ErrorVotanteFraudulento{Dni: votante.dni}
		}
	} else if alternativa == VOTO_IMPUGNADO {
		votante.voto.Impugnado = true
	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}
	votante.historial_voto.Apilar(votante.voto)
	return nil
}

func (votante *votanteImplementacion) Deshacer() error {
	if votante.yaVoto {
		if votante.LeerDNI()[0:1] == "0" {
			return errores.ErrorVotanteFraudulento{Dni: votante.dni[1:]}
		} else {
			return errores.ErrorVotanteFraudulento{Dni: votante.dni}
		}
	}

	votante.historial_voto.Desapilar()
	if votante.historial_voto.EstaVacia() {
		votante.historial_voto.Apilar(votante.voto)
		return errores.ErrorNoHayVotosAnteriores{}
	} else {
		votante.voto = votante.historial_voto.VerTope()
	}
	return nil
}

func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	if votante.yaVoto {
		if votante.LeerDNI()[0:1] == "0" {
			return votante.voto, errores.ErrorVotanteFraudulento{Dni: votante.dni[1:]}
		} else {
			return votante.voto, errores.ErrorVotanteFraudulento{Dni: votante.dni}
		}
	}
	votante.yaVoto = true
	return votante.voto, nil
}
