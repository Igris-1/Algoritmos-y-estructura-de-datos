package redUsuario

import (
	abb "Algogram/tdas/abb"
)

type post struct {
	id         int
	publicador string
	contenido  string
	personas   abb.DiccionarioOrdenado[string, int]
}

func compareString(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	} else {
		return 0
	}
}

// CrearPost: crea un post con el id, el publicador y el contenido
func CrearPost(id int, publicador string, contenido string) Post {
	return &post{id, publicador, contenido, abb.CrearABB[string, int](compareString)}
}

// PostID: devuelve el id del post
func (p post) PostID() int { return p.id }

// PostLikes: devuelve el diccionario de personas que le dieron like al post
func (p post) PostLikes() abb.DiccionarioOrdenado[string, int] { return p.personas }

// Publicar: devuelve el nombre del publicador del post
func (p post) Publicador() string { return p.publicador }

// Contenido: devuelve el mensaje del post
func (p post) Contenido() string { return p.contenido }
