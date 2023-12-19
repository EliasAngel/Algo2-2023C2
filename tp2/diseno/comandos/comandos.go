package comandos

import (
	TDAs "algogram/diseno/algogram"
	"algogram/diseno/errores"
	"fmt"
	"strconv"
	"strings"
)

// login() loguea a un usuario
func login(entrada string, app *TDAs.Aplicacion) error {
	if !(*app).UsuarioEstaRegistrado(entrada) {
		return &errores.ErrorUsuarioNoExiste{}
	}
	if (*app).ObtenerUsuarioLogueado() != nil {
		return &errores.ErrorUsuarioYaLogueado{}
	}
	user := (*app).ObtenerUsuario(entrada)
	(*app).IniciarSesion(user)
	fmt.Println("Hola " + (*user).ObtenerNombre())
	return nil
}

// logout() desloguea a un usuario
func logout(entrada []string, app *TDAs.Aplicacion) error {
	if len(entrada) != 1 {
		return &errores.ErrorParametros{}
	}
	if (*app).ObtenerUsuarioLogueado() == nil {
		return &errores.ErrorUsuarioNoLogueado{}
	}
	(*app).CerrarSesion()
	fmt.Println("Adios")
	return nil
}

// publicar() publica un post
func publicar(entrada []string, app *TDAs.Aplicacion) error {
	if (*app).ObtenerUsuarioLogueado() == nil {
		return &errores.ErrorUsuarioNoLogueado{}
	}
	textoPost := strings.Join(entrada[1:], " ")
	(*app).SubirPost(textoPost)
	return nil
}

// verSiguienteFeed() muestra el siguiente post del feed
func verSiguienteFeed(entrada []string, app *TDAs.Aplicacion) error {
	if len(entrada) != 1 {
		return &errores.ErrorParametros{}
	}
	user := (*app).ObtenerUsuarioLogueado()
	if user == nil {
		return &errores.ErroMostrarFeed{}
	}
	postAMostrar := (*user).VerSiguientePost()
	if postAMostrar != nil {
		postID := (*postAMostrar).ObtenerID()
		post := (*app).ObtenerPost(postID)
		(*post).MostrarPost()
	} else {
		return errores.ErroMostrarFeed{}
	}
	return nil
}

// likearPost() likea un post
func likearPost(entrada []string, app *TDAs.Aplicacion) error {
	if len(entrada) != 2 {
		return &errores.ErrorParametros{}
	}
	if (*app).ObtenerUsuarioLogueado() == nil {
		return &errores.ErrorUsuarioNoLogueadoOpostNoExiste{}
	}
	idPost := entrada[1]
	idPostInt, err := strconv.Atoi(idPost)
	if err != nil {
		return &errores.ErrorAlConvertirAEntero{}
	}
	user := (*app).ObtenerUsuarioLogueado()
	if !(*app).PostExiste(idPostInt) {
		return &errores.ErrorUsuarioNoLogueadoOpostNoExiste{}
	}
	post := (*app).ObtenerPost(idPostInt)
	(*user).Likear(post)
	return nil
}

// mostrarLikes() muestra los likes de un post
func mostrarLikes(entrada []string, app *TDAs.Aplicacion) error {
	if len(entrada) != 2 {
		return &errores.ErrorParametros{}
	}
	id, err := strconv.Atoi(entrada[1])
	if err != nil {
		return &errores.ErrorAlConvertirAEntero{}
	}
	post := (*app).ObtenerPost(id)
	if post == nil {
		return &errores.ErrorPostNoExiste{}
	} else if (*post).CantidadLikes() == 0 {
		return &errores.ErrorMostrarLikes{}
	}
	fmt.Printf("El post tiene %d likes:\n", (*post).CantidadLikes())
	(*post).MostrarLikes()
	return nil
}

// comandoLogin() ejecuta el comando login
func comandoLogin(entradaUsuario []string, app *TDAs.Aplicacion) {
	entrada := strings.Join(entradaUsuario[1:], " ")
	error := login(entrada, app)
	if error != nil {
		fmt.Println(error.Error())
	}
}

// comandoLogout() ejecuta el comando logout
func comandoLogout(entradaUsuario []string, app *TDAs.Aplicacion) {
	error := logout(entradaUsuario, app)
	if error != nil {
		fmt.Println(error.Error())
	}
}

// comandoPublicar() ejecuta el comando publicar
func comandoPublicar(entradaUsuario []string, app *TDAs.Aplicacion) {
	error := publicar(entradaUsuario, app)
	if error != nil {
		fmt.Println(error.Error())
	} else {
		fmt.Println("Post publicado")
	}
}

// comandoVerSiguienteFeed() ejecuta el comando ver_siguiente_feed
func comandoVerSiguienteFeed(entradaUsuario []string, app *TDAs.Aplicacion) {
	error := verSiguienteFeed(entradaUsuario, app)
	if error != nil {
		fmt.Println(error.Error())
	}
}

// comandoLikearPost() ejecuta el comando likear_post
func comandoLikearPost(entradaUsuario []string, app *TDAs.Aplicacion) {
	error := likearPost(entradaUsuario, app)
	if error != nil {
		fmt.Println(error.Error())
	}
}

// comandoMostrarLikes() ejecuta el comando mostrar_likes
func comandoMostrarLikes(entradaUsuario []string, app *TDAs.Aplicacion) {
	error := mostrarLikes(entradaUsuario, app)
	if error != nil {
		fmt.Println(error.Error())
	}
}

// ProcesarEntrada() procesa la entrada del usuario
func ProcesarEntrada(entradaUsuario []string, app *TDAs.Aplicacion) {
	switch entradaUsuario[0] {
	case "login":
		comandoLogin(entradaUsuario, app)
	case "logout":
		comandoLogout(entradaUsuario, app)
	case "publicar":
		comandoPublicar(entradaUsuario, app)
	case "ver_siguiente_feed":
		comandoVerSiguienteFeed(entradaUsuario, app)
	case "likear_post":
		comandoLikearPost(entradaUsuario, app)
	case "mostrar_likes":
		comandoMostrarLikes(entradaUsuario, app)
	default:
		fmt.Println(&errores.ErrorComandoInvalido{})
	}
}
