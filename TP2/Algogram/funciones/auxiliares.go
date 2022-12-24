package auxiliares

import (
	red "Algogram/redUsuario/red"
	usuario "Algogram/redUsuario/usuario"
	"bufio"
	"os"
	"strings"
)

// AgregarUsuarios: Lee el archivo de usuarios, crea una red, un registro y agrega los usuarios a la red
func AgregarUsuarios(ruta string) (red.Red, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, err
	}
	defer archivo.Close()

	red := red.CrearRed()
	registro := red.Registrados()

	usuarios := bufio.NewScanner(archivo)
	posicion := 0

	for usuarios.Scan() {
		usuario := usuario.CrearUsuario(usuarios.Text(), posicion, &registro, usuario.Cmp)
		registro.Guardar(usuarios.Text(), usuario)
		posicion++
	}
	return red, nil
}

// LeerComand: toma el input del usuario y lo separa en comando y argumento
func LeerComando(entrada string) (string, string) {
	dato := strings.Split(entrada, " ")
	comando := dato[0]
	return comando, strings.Join(dato[1:], " ")
}
