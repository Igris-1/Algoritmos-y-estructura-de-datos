package redUsuario

import (
	abb "Algogram/tdas/abb"
)

type Post interface {
	// PostID: devuelve el id del post
	PostID() int

	// Publicar: devuelve el nombre del publicador del post
	Publicador() string

	// Contenido: devuelve el mensaje del post
	Contenido() string

	// PostLikes: devuelve el diccionario de personas que le dieron like al post
	PostLikes() abb.DiccionarioOrdenado[string, int]
}
