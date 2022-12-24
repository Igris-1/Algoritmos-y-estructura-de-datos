package errores

import "fmt"

/* ERROR LECTURA DE ARCHIVOS */
type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

/* ERROR PARAMETROS */
type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

/* ERROR DNI */
type DNIError struct{}

func (e DNIError) Error() string {
	return "ERROR: DNI incorrecto"
}

/* ERROR DNI VOTANTE */
type DNIFueraPadron struct{}

func (e DNIFueraPadron) Error() string {
	return "ERROR: DNI fuera del padrón"
}

/* ERROR FILA VACIA */
type FilaVacia struct{}

func (e FilaVacia) Error() string {
	return "ERROR: Fila vacía"
}

/* ERROR VOTANTE FRAUDULENTO */
type ErrorVotanteFraudulento struct {
	Dni string
}

func (e ErrorVotanteFraudulento) Error() string {
	return fmt.Sprintf("ERROR: Votante FRAUDULENTO: %v", e.Dni)
}

/* ERROR TIPO VOTO */
type ErrorTipoVoto struct{}

func (e ErrorTipoVoto) Error() string {
	return "ERROR: Tipo de voto inválido"
}

/* ERROR ALTERNATIVA INVALIDA */
type ErrorAlternativaInvalida struct{}

func (e ErrorAlternativaInvalida) Error() string {
	return "ERROR: Alternativa inválida"
}

/* ERROR NO HAY VOTOS ANTERIORES */
type ErrorNoHayVotosAnteriores struct{}

func (e ErrorNoHayVotosAnteriores) Error() string {
	return "ERROR: Sin voto a deshacer"
}

/* ERROR CIUDADANOS SIN VOTAR */
type ErrorCiudadanosSinVotar struct{}

func (e ErrorCiudadanosSinVotar) Error() string {
	return "ERROR: Ciudadanos sin terminar de votar"
}
