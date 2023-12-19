package main

import (
	TDAs "algogram/diseno/algogram"
	Comando "algogram/diseno/comandos"
	Funciones "algogram/diseno/funciones"
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	entradaPrincipal := os.Args
	usuarios, err := Funciones.LeerArchivoUsuario(entradaPrincipal[1])
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	app := TDAs.IniciarAplicacion(usuarios)
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		entradaUsuario := strings.Split(entrada.Text(), " ")
		Comando.ProcesarEntrada(entradaUsuario, &app)
	}
}
