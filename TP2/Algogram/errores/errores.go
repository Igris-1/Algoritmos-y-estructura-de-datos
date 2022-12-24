package errores

/* ERROR USUARIO */

// En caso que ya hubiera un usuario loggeado
type UsuarioYaLoggeado struct{}

func (e UsuarioYaLoggeado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

// En caso que el usuario indicado no exista
type UsuarioNoExiste struct{}

func (e UsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

/* ERROR SIN USUARIO LOGGEADO */

// logout y publicar post en caso que no hubiera un usuario loggeado
type UsuarioNoLoggeado struct{}

func (e UsuarioNoLoggeado) Error() string {
	return "Error: no habia usuario loggeado"
}

/* ERROR VER PROXIMO POST SIN USUARIO LOGGEADO O SIN POSTS */

// En caso que un usuario no tenga más posts para ver, o bien que no haya usuario loggeado
type NoHayMasPost struct{}

func (e NoHayMasPost) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type SinPostsOsinLoggeado struct{}

func (e SinPostsOsinLoggeado) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

/* ERROR LIKEAR UN POST*/

// En caso que no haya un usuario loggeado o el post en cuestión no exista
type PostNoExiste struct{}

func (e PostNoExiste) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

/* ERROR MOSTRAR LIKES */

// En caso que un post no tenga likes, o bien no exista el id
type PostInexistenteOSinLikes struct{}

func (e PostInexistenteOSinLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}

/* ERROR DE LECTURA DE ARCHIVO */

type ErrorLecturaArchivo struct{}

func (e ErrorLecturaArchivo) Error() string {
	return "Error: No se pudo leer el archivo"
}
