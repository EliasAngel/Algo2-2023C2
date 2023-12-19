package algogram

import (
	"fmt"
	"strings"
	TDADiccionario "tdas/diccionario"
)

var (
	contador = -1 //Empiezo en -1 para que haya un post con ID 0
)

type postImplementacion struct {
	Likes         TDADiccionario.DiccionarioOrdenado[string, bool]
	ID            int
	pieDePost     string
	autor         string
	afinidadAutor int
}

// CrearPost() crea un post con el pie de post pasado por parametro y lo devuelve O(1)
func CrearPost(pieDePost string, autor *Usuario) Post {
	post := &postImplementacion{
		Likes:         TDADiccionario.CrearABB[string, bool](strings.Compare),
		ID:            contador + 1,
		pieDePost:     pieDePost,
		autor:         (*autor).ObtenerNombre(),
		afinidadAutor: (*autor).ObtenerAfinidad(),
	}
	contador++
	return post
}

// ObtenerID() Obtiene el id único del post. O(1)
func (post *postImplementacion) ObtenerID() int {
	return post.ID
}

// AgregarLike() Agrega un like al post por parte de un usuario con el nombre proporcionado.
// También registra si el usuario actual dio like a su propio post. O(log like)
func (post *postImplementacion) AgregarLike(nombre string) {
	if !post.Likes.Pertenece(nombre) {
		post.Likes.Guardar(nombre, true)
	}
	fmt.Println("Post likeado")
}

// VerPieDePost() Obtiene el contenido del pie de post del post. O(1)
func (post *postImplementacion) VerPieDePost() string {
	return post.pieDePost
}

// ObtenerAutor() Obtiene el autor del post. O(1)
func (post *postImplementacion) ObtenerAutor() string {
	return post.autor
}

func (post *postImplementacion) ObtenerAfinidad() int {
	return post.afinidadAutor
}

// MostrarPost() Muestra la información básica del post, incluyendo su identificador, el nombre del autor, el
// contenido del pie de post y la cantidad de likes que tiene. O(1)
func (post *postImplementacion) MostrarPost() {
	fmt.Printf("Post ID %d\n", post.ID)
	fmt.Printf("%s dijo: %s\n", post.autor, post.VerPieDePost())
	fmt.Printf("Likes: %d\n", post.Likes.Cantidad())
}

// MostrarLikes() Muestra la cantidad de likes del post y lista los nombres de los usuarios que dieron like. O(U)
func (post *postImplementacion) MostrarLikes() {
	post.Likes.Iterar(func(clave string, _ bool) bool {
		fmt.Println("	" + clave)
		return true
	})
}

// CantidadLikes() Devuelve la cantidad de likes que tiene el post. O(1)
func (post *postImplementacion) CantidadLikes() int {
	return post.Likes.Cantidad()
}
