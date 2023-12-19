package algogram

import (
	TDAHeap "tdas/cola_prioridad"
)

type UsuarioImplementacion struct {
	nombre       string
	afinidad     int // afinidad representa la posicion del usuario en el archivo de usuarios
	postsParaVer TDAHeap.ColaPrioridad[*Feed]
}

type Feed struct {
	auxPost  *Post
	afinidad int
}

// calcularDistancia() calcula la distancia entre dos numeros O(1)
func calcularDistancia(a, b int) int {
	if a-b > 0 {
		return a - b
	}
	return -(a - b)
}

// compararID() compara dos id de post y devuelve 1 si a es mayor que b, -1 si a es menor que b y 0 si son iguales O(1)
func compararID(a, b int) int {
	if a > b {
		return -1
	}
	if a < b {
		return 1
	}
	return 0
}

// compararPost() compara dos post y en base a la afinidad del usuario y la afinidad de los post O(1)
func compararPost(a, b *Feed) int {
	postA := *(*a).auxPost
	postB := *(*b).auxPost
	afinidadDelUsuario := (*a).afinidad
	afinidadDePostA := postA.ObtenerAfinidad()
	afinidadDePostB := postB.ObtenerAfinidad()
	if calcularDistancia(afinidadDelUsuario, afinidadDePostA) > calcularDistancia(afinidadDelUsuario, afinidadDePostB) {
		return -1
	}
	if calcularDistancia(afinidadDelUsuario, afinidadDePostA) < calcularDistancia(afinidadDelUsuario, afinidadDePostB) {
		return 1
	} else {
		return compararID(postA.ObtenerID(), postB.ObtenerID())
	}

}

// CrearUsuario() crea un usuario con el nombre y afinidad pasados por parametro O(1)
func CrearUsuario(nombre string, afinidad int) Usuario {
	return &UsuarioImplementacion{
		nombre:       nombre,
		afinidad:     afinidad,
		postsParaVer: TDAHeap.CrearHeap(compararPost),
	}
}

// ObtenerNombre() Obtiene el nombre del usuario O(1)
func (user *UsuarioImplementacion) ObtenerNombre() string {
	return user.nombre
}

// VerSiguientePost() Muestra el proximo post del feed del usuario O(log n)
func (user *UsuarioImplementacion) VerSiguientePost() *Post {
	if !user.postsParaVer.EstaVacia() {
		feed := user.postsParaVer.Desencolar()
		return feed.auxPost
	}
	return nil
}

// ObtenerAfinidad() Obtiene la afinidad del usuario en base a su posicion en el archivo de usuarios O(1)
func (user *UsuarioImplementacion) ObtenerAfinidad() int {
	return user.afinidad
}

// Likear() Likea un post  O(log like)
func (user *UsuarioImplementacion) Likear(post *Post) {
	(*post).AgregarLike(user.nombre)
}

// ActualizarFeed() Actualiza el feed del usuario con un nuevo post O(log n)
func (user *UsuarioImplementacion) ActualizarFeed(feed *Feed) {
	user.postsParaVer.Encolar(feed)
}
