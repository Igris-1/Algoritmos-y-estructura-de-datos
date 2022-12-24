package redUsuario

import (
	post "Algogram/redUsuario/post"
	user "Algogram/redUsuario/usuario"
	abb "Algogram/tdas/abb"
	hash "Algogram/tdas/hash"
)

type Red interface {
	// LoggIn: loggea al usuario en la red, si ya se hallaba un usuario loggeado o el usuario en cuestio
	// no existe, devuelve error
	LoggIn(nombre string) error

	// LoggOut: desloggea al usuario de la red, si no se hallaba un usuario loggeado, devuelve error
	LoggOut() error

	// Loggeado: devuelve el usuario loggeado, si no se hallaba un usuario loggeado, devuelve error
	Loggeado() (user.User, error)

	// MostrarLikes: devuelve el diccionario de personas que likearon el correspondiente post
	// si el post no existe, devuelve error
	MostrarLikes(id int) (abb.DiccionarioOrdenado[string, int], error)

	// PublicarPost: publica un post en la red, si no se halla un usuario loggeado, devuelve error
	PublicarPost(post post.Post) error

	// Likear: permte al usuario loggeado likear un determinado post por su id, si no se halla un usuario
	// loggeado o el post no existe, devuelve error
	Likear(id int) error

	// Regstrados: devuelve el diccionario de usuarios registrados en la red
	Registrados() hash.Diccionario[string, user.User]

	// CantidadPost: devuelve la cantidad de posts publicados en la red
	CantidadPost() int
}
