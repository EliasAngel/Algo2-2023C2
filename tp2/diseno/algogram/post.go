package algogram

type Post interface {
	// Obtiene el id de un  post
	ObtenerID() int

	//Agrega un like a un post
	AgregarLike(nombre string)

	//Obtiene el texto del pie de post
	VerPieDePost() string

	//Devuelve la afinidad del autor del post
	ObtenerAutor() string

	ObtenerAfinidad() int

	//Muestra el post
	MostrarPost()

	//Muestra los likes del post
	MostrarLikes()

	//Devuelve la cantidad de likes que tiene el post
	CantidadLikes() int
}
