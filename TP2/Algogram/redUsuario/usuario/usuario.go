package redUsuario

import (
	post "Algogram/redUsuario/post"
	hash "Algogram/tdas/hash"
	heap "Algogram/tdas/heap"
)

type User interface {
	// VerProximoPost: devuelve el proximo post del feed del usuario
	VerProximoPost() *post.Post

	// PosicionUsuario: devuelve la posicion del usuario en la red
	PosicionUsuario() int

	// Feed: devuelve la cola de prioridad de posts del usuario
	Feed() heap.ColaPrioridad[*post.Post]

	// NombreUsuario: devuelve el nombre del usuario
	NombreUsuario() string

	// Registro: devuelve el diccionario de usuarios registrados en la red
	Registro() hash.Diccionario[string, User]
}
