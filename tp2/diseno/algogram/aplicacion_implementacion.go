package algogram

import (
	TDAHash "tdas/diccionario"
)

type aplicacionImplementacion struct {
	usuarioActual       Usuario
	usuariosRegistrados TDAHash.Diccionario[string, *Usuario]
	post                TDAHash.Diccionario[int, *Post]
}

// IniciarAplicacion() Inicia la aplicación con los usuarios registrados pasados por parámetro. 0(1)
func IniciarAplicacion(usuarios TDAHash.Diccionario[string, *Usuario]) Aplicacion {
	return &aplicacionImplementacion{
		usuarioActual:       nil,
		usuariosRegistrados: usuarios,
		post:                TDAHash.CrearHash[int, *Post](),
	}
}

// IniciarSesion() inicia sesion con el usuario pasado por parametro y lo agrega al atributo de UsuarioActual O(1)
func (app *aplicacionImplementacion) IniciarSesion(usuario *Usuario) {
	app.usuarioActual = *usuario
}

// CerrarSesion() cierra la sesion del usuario actual O(1)
func (app *aplicacionImplementacion) CerrarSesion() {
	app.usuarioActual = nil
}

// ObtenerUsuarioLogueado() Obtiene el usuario actualmente logueado en la aplicación. O(1)
func (app *aplicacionImplementacion) ObtenerUsuarioLogueado() *Usuario {
	if app.usuarioActual == nil {
		return nil
	}
	return &app.usuarioActual
}

// UsuarioEstaRegistrado() Verifica si un usuario está registrado en la aplicación por su nombre. O(1)
func (app *aplicacionImplementacion) UsuarioEstaRegistrado(nombre string) bool {
	return app.usuariosRegistrados.Pertenece(nombre)
}

// ObtenerUsuario() Obtiene un usuario registrado en la aplicación por su nombre. O(1)
func (app *aplicacionImplementacion) ObtenerUsuario(nombre string) *Usuario {
	return app.usuariosRegistrados.Obtener(nombre)
}

// SubirPost() Publica un nuevo post y lo agrega al feed de todos los usuarios registrados en la aplicación. O(U * log p)
func (app *aplicacionImplementacion) SubirPost(textoPost string) {
	post := CrearPost(textoPost, &app.usuarioActual)
	iter := app.usuariosRegistrados.Iterador()
	for iter.HaySiguiente() {
		_, user := iter.VerActual()
		if post.ObtenerAutor() != (*user).ObtenerNombre() {
			feed := &Feed{auxPost: &post, afinidad: (*user).ObtenerAfinidad()}
			(*user).ActualizarFeed(feed)
		}
		iter.Siguiente()
	}
	app.post.Guardar(post.ObtenerID(), &post)
}

// PostExiste() Verifica si existe un post con el ID especificado en la aplicación. O(1)
func (app *aplicacionImplementacion) PostExiste(id int) bool {
	return app.post.Pertenece(id)
}

// ObtenerPost() Obtiene el post con el ID especificado en la aplicación. O(1)
func (app *aplicacionImplementacion) ObtenerPost(id int) *Post {
	if app.PostExiste(id) {
		return app.post.Obtener(id)
	}
	return nil
}
