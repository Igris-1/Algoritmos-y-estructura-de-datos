package admmemoria_test

import (
	Persona "administracionmemoria"
	"administracionmemoria/administrador"
	"testing"
)

func TestPersonaUnica(t *testing.T) {
	per := Persona.CrearPersona("Jorge", nil)
	per.Destruir()
	administrador.Finalizar()
}

func TestPersonaConUnHijoYUnNieto(t *testing.T) {
	per := Persona.CrearPersona("Abuelo", nil)
	Persona.CrearPersona("Nieto", Persona.CrearPersona("hijo", per))
	per.Destruir()
	administrador.Finalizar()
}

func TestConVariosHijosYNietos(t *testing.T) {
	estanislao := Persona.CrearPersona("Estanislao", nil)
	jolanta := Persona.CrearPersona("Jolanta", estanislao)
	Persona.CrearPersona("Maximiliano", jolanta)
	Persona.CrearPersona("Francisco", jolanta)
	Persona.CrearPersona("Marcelina", Persona.CrearPersona("Igor", Persona.CrearPersona("Cristofer", estanislao)))
	estanislao.Destruir()
	administrador.Finalizar()
}
