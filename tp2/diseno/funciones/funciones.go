package funciones

import (
	TDAs "algogram/diseno/algogram"
	"algogram/diseno/errores"
	"bufio"
	"os"
	TDAHash "tdas/diccionario"
)

// LeerArchivoUsuario lee un archivo y devuelve un diccionario donde la clave es el nombre del usuario y el valor es el usuario
func LeerArchivoUsuario(ruta string) (TDAHash.Diccionario[string, *TDAs.Usuario], error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, &errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	usuarios := TDAHash.CrearHash[string, *TDAs.Usuario]()
	s := bufio.NewScanner(archivo)
	contador := 0
	for s.Scan() {
		usuario := s.Text()
		user := TDAs.CrearUsuario(usuario, contador)
		usuarios.Guardar(usuario, &user)
		contador++
	}
	err = s.Err()
	if err != nil {
		return nil, &errores.ErrorLeerArchivo{}
	}
	return usuarios, nil
}
