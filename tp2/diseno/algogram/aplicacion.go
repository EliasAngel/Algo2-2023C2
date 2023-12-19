package algogram

type Aplicacion interface {

	// Inicia la aplicacion con un usuario logueado
	IniciarSesion(usuario *Usuario)

	// Cierra la sesion del usuario actual
	CerrarSesion()

	// Devuelve el usuario logueado
	ObtenerUsuarioLogueado() *Usuario

	// Devuelve true si el usuario esta registrado en el hash de usuarios registrados
	UsuarioEstaRegistrado(nombre string) bool

	// Devuelve el usuario registrado
	ObtenerUsuario(nombre string) *Usuario

	// Devuelve el post con el id pasado por parametro
	ObtenerPost(id int) *Post

	// Devuelve true si el post existe
	PostExiste(id int) bool

	SubirPost(textoPost string)
}
