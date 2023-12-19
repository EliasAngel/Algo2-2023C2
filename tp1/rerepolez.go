package main

//Commit de Prueba
import (
	//Importo el TDA Cola
	"bufio"
	"fmt"
	"os"
	Elector "rerepolez/diseno_alumnos/votos"
	Comando "rerepolez/funciones_rerepolez"
	"strings"
	TdaCola "tdas/cola"
)

func main() {
	entradaPrincipal := os.Args
	partidos, padron, err := Comando.VerificarEntrada(entradaPrincipal)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	votoEnBlanco := Elector.CrearVotosEnBlanco()
	fila := TdaCola.CrearColaEnlazada[Elector.Votante]()
	impugnados := 0
	entrada := bufio.NewScanner(os.Stdin)
	for entrada.Scan() {
		entradaUsuario := strings.Split(entrada.Text(), " ")
		Comando.ProcesarEntrada(entradaUsuario, padron, partidos, fila, votoEnBlanco, &impugnados)
	}
	Comando.MostrarResultados(fila, partidos, impugnados, votoEnBlanco)
}
