package comandos

import (
	errores "Algogram/errores"
	post "Algogram/redUsuario/post"
	red "Algogram/redUsuario/red"
	"fmt"
	"strconv"
)

// LoggIn: loggea al usuario, si no exste devuelve un error
func LoggIn(red red.Red, nombre string) {
	err := red.LoggIn(nombre)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Hola", nombre)
	}
}

// LoggOut: desloggea al usuario, si no habia nadie loggeado devuelve un error
func LoggOut(red red.Red) {
	err := red.LoggOut()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Adios")
	}
}

// Publicar: publica un post en la red, s no haba nadie loggeado devuelve un error
func Publicar(red red.Red, contenido string) {
	loggeado, err := red.Loggeado()
	if err != nil {
		fmt.Println(err)
	} else {
		cantPost := red.CantidadPost()
		post := post.CrearPost(cantPost, loggeado.NombreUsuario(), contenido)
		red.PublicarPost(post)
		fmt.Println("Post publicado")
	}
}

// VerSiguienteFeed: muestra el siguiente post del feed del usuario loggeado, si no hay nadie loggeado devuelve un error
func VerSiguienteFeed(red red.Red) {
	usuario, err := red.Loggeado()
	if err != nil || usuario.Feed().EstaVacia() {
		fmt.Println(errores.SinPostsOsinLoggeado{})
	} else {
		post := *usuario.Feed().Desencolar()
		fmt.Println("Post ID", post.PostID())
		fmt.Println(post.Publicador(), "dijo:", post.Contenido())
		fmt.Println("Likes:", post.PostLikes().Cantidad())
	}
}

// Likear: Likea un post de la red, si no hay nadie loggeado o el post no existe devuelve un error
func Likear(red red.Red, id string) {
	idInt, _ := strconv.Atoi(id)
	err := red.Likear(idInt)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Post likeado")
	}
}

// MostrarLikes: muestr el el posteador, el id del post y la cantidad de personas que likearon el post
// si el post no existe o la cantidad de likes es 0 devuelve un error
func MostrarLikes(red red.Red, id string) {
	idInt, _ := strconv.Atoi(id)
	arbol, err := red.MostrarLikes(idInt)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("El post tiene", arbol.Cantidad(), "likes:")
	arbol.Iterar(func(key string, value int) bool {
		fmt.Println("	" + key)
		return true
	})
}
