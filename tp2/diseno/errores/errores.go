package errores

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

type ErrorComandoInvalido struct{}

func (e ErrorComandoInvalido) Error() string {
	return "ERROR: Comando inválido"
}

type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros, o los parámetros son inválidos"
}

type ErrorUsuarioNoExiste struct{}

func (e ErrorUsuarioNoExiste) Error() string {
	return "Error: usuario no existente"
}

type ErrorUsuarioYaLogueado struct{}

func (e ErrorUsuarioYaLogueado) Error() string {
	return "Error: Ya habia un usuario loggeado"
}

type ErrorUsuarioNoLogueado struct{}

func (e ErrorUsuarioNoLogueado) Error() string {
	return "Error: no habia usuario loggeado"
}

type ErrorUsuarioNoLogueadoOpostNoExiste struct{}

func (e ErrorUsuarioNoLogueadoOpostNoExiste) Error() string {
	return "Error: Usuario no loggeado o Post inexistente"
}

type ErrorAlConvertirAEntero struct{}

func (e ErrorAlConvertirAEntero) Error() string {
	return "ERROR: no se pudo convertir a entero"
}

type ErrorUsuarioNoTienePost struct{}

func (e ErrorUsuarioNoTienePost) Error() string {
	return "ERROR: Post inexistente"
}

type ErrorUsuarioYaLikeoPost struct{}

func (e ErrorUsuarioYaLikeoPost) Error() string {
	return "ERROR: El usuario ya likeo el post"
}

type ErrorPostNoExiste struct{}

func (e ErrorPostNoExiste) Error() string {
	return "Error: Post inexistente o sin likes"
}

type ErroMostrarFeed struct{}

func (e ErroMostrarFeed) Error() string {
	return "Usuario no loggeado o no hay mas posts para ver"
}

type ErrorMostrarLikes struct{}

func (e ErrorMostrarLikes) Error() string {
	return "Error: Post inexistente o sin likes"
}
