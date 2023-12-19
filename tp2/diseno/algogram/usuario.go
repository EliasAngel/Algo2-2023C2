package algogram

type Usuario interface {
	//Obtiene el nombre del usuario
	ObtenerNombre() string

	//Obtiene la afinidad del usuario en base a su posicion en el archivo de usuarios
	ObtenerAfinidad() int

	//Muestra el proximo post del feed del usuario
	VerSiguientePost() *Post

	//Likea un post
	Likear(post *Post)

	ActualizarFeed(feed *Feed)
}
