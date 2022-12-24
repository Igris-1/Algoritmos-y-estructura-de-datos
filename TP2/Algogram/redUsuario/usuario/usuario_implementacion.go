package redUsuario

import (
	post "Algogram/redUsuario/post"
	hash "Algogram/tdas/hash"
	heap "Algogram/tdas/heap"
)

// funcion auxiliar para calcular el modulo
func modulo(a int) int {
	if a < 0 {
		return -a
	} else {
		return a
	}
}

func compareInts(af1, af2, idA, idB int) int {
	if af1 < af2 {
		return 1
	}
	if af1 > af2 {
		return -1
	}
	if idA < idB {
		return 1
	}
	if idA > idB {
		return -1
	}
	return 0
}

// funcion auxiliar de calculo de afnidades respecto a un determinado usuario
// si las afinidades son iguales, se compara el id de los posts
func Cmp(u usuario) func(a, b *post.Post) int {
	return func(a, b *post.Post) int {
		registro := u.Registro()
		postA := *a
		postB := *b

		usuarioA := registro.Obtener(postA.Publicador())
		usuarioB := registro.Obtener(postB.Publicador())

		afinidadA := modulo(u.PosicionUsuario() - usuarioA.PosicionUsuario())
		afinidadB := modulo(u.PosicionUsuario() - usuarioB.PosicionUsuario())

		idA := postA.PostID()
		idB := postB.PostID()
		return compareInts(afinidadA, afinidadB, idA, idB)
	}
}

type usuario struct {
	nombre   string
	posicion int
	cmp      func(u usuario) func(a, b *post.Post) int
	feed     heap.ColaPrioridad[*post.Post]
	reg      *hash.Diccionario[string, User]
}

// CrearUsuario: crea un usuario con el nombre, posicion, registro de la red y una funcion de comparacion
func CrearUsuario(nombre string, pos int, reg *hash.Diccionario[string, User], cmp func(u usuario) func(a *post.Post, b *post.Post) int) User {
	var usuario usuario
	usuario.nombre = nombre
	usuario.posicion = pos
	usuario.cmp = cmp
	usuario.reg = reg
	usuario.feed = heap.CrearHeap(cmp(usuario))
	return usuario
}

// VerProximoPost: devuelve el post con mayor afinidad respecto al usuario
func (u usuario) VerProximoPost() *post.Post { return u.feed.Desencolar() }

// Feed: devuelve la cola de prioridad de posts del usuario
func (u usuario) Feed() heap.ColaPrioridad[*post.Post] { return u.feed }

// NombreUsuario: devuelve el nombre del usuario
func (u usuario) NombreUsuario() string { return u.nombre }

// PosicionUsuario: devuelve la posicion del usuario en el registro de la red
func (u usuario) PosicionUsuario() int { return u.posicion }

// Registro: devuelve el diccionario de usuarios registrados en la red
func (u usuario) Registro() hash.Diccionario[string, User] { return *u.reg }
