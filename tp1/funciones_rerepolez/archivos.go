package funciones_rerepolez

import (
	"bufio"
	"os"
	"rerepolez/diseno_alumnos/errores"
	eleccion "rerepolez/diseno_alumnos/votos"
	"strconv"
	"strings"
)

const (
	PARTIDO_IMPUGNADO = 1
	TAMANIO_BUCKET    = 10
)

// LeerArchivoListas lee el archivo de listas y devuelve una lista de partidos o un error de tipo ErrorLeerArchivo en caso que no se pueda leer el archivo
func LeerArchivoListas(ruta string) ([]eleccion.Partido, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, errores.ErrorLeerArchivo{}
	}
	defer archivo.Close()
	listas := make([]eleccion.Partido, PARTIDO_IMPUGNADO)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		datos := strings.Split(linea, ",")
		nombrePartido := datos[0]
		candidatos := [eleccion.CANT_VOTACION]string{datos[1], datos[2], datos[3]}
		listas = append(listas, eleccion.CrearPartido(nombrePartido, candidatos))
	}
	err = s.Err()
	if err != nil {
		return nil, error(errores.ErrorLeerArchivo{})
	}
	return listas, nil
}

// LeerArchivoPadron lee el archivo de padron y devuelve una lista de dnis ordenada o un error de tipo ErrorLeerArchivo en caso que no se pueda leer el archivo
func LeerArchivoPadron(ruta string) ([]eleccion.Votante, error) {
	archivo, err := os.Open(ruta)
	if err != nil {
		return nil, error(errores.ErrorLeerArchivo{})
	}
	defer archivo.Close()
	padron := make([]eleccion.Votante, 0)
	s := bufio.NewScanner(archivo)
	for s.Scan() {
		linea := s.Text()
		dni, _ := strconv.Atoi(linea)
		padron = append(padron, eleccion.CrearVotante(dni))
	}
	err = s.Err()
	if err != nil {
		return nil, error(errores.ErrorLeerArchivo{})
	}
	radixSort(padron)
	return padron, nil
}

// VerificarEntrada verifica que la entrada del programa sea correcta para poder ejecutar el programa, y devuelve los partidos y el padron, caso contrario
// devuelve un error de tipo ErrorParametros
func VerificarEntrada(entrada []string) ([]eleccion.Partido, []eleccion.Votante, error) {
	if len(entrada) != 3 {
		return nil, nil, error(errores.ErrorParametros{})
	}
	partidos, err := LeerArchivoListas(entrada[1])
	if err != nil {
		return nil, nil, err
	}
	padron, err := LeerArchivoPadron(entrada[2])
	if err != nil {
		return nil, nil, err
	}
	return partidos, padron, nil
}

// Función de conteo para contar la frecuencia de cada dígito en la lista de dnis
func contarFrecuencia(dni []eleccion.Votante, potencia int) []int {
	frecuencias := make([]int, TAMANIO_BUCKET)
	for _, votante := range dni {
		numero := votante.LeerDNI()
		digito := (numero / potencia) % TAMANIO_BUCKET
		frecuencias[digito]++
	}
	return frecuencias
}

// Función de conteo para calcular las posiciones finales de cada dígito en la lista de dnis
func calcularSumAcum(frecuencias []int) []int {
	posiciones := make([]int, TAMANIO_BUCKET)
	posiciones[0] = frecuencias[0]
	for i := 1; i < TAMANIO_BUCKET; i++ {
		posiciones[i] = posiciones[i-1] + frecuencias[i]
	}
	return posiciones
}

// Función de ordenación para realizar el ordenamiento basado en el dígito actual de la lista de dnis
func ordenarPorDigito(dni []eleccion.Votante, potencia int) {
	frecuencias := contarFrecuencia(dni, potencia)
	sumAcumulativa := calcularSumAcum(frecuencias)
	resultado := make([]eleccion.Votante, len(dni))

	for i := len(dni) - 1; i >= 0; i-- {
		digito := (dni[i].LeerDNI() / potencia) % TAMANIO_BUCKET
		pos := sumAcumulativa[digito] - 1
		resultado[pos] = dni[i]
		sumAcumulativa[digito]--
	}

	copy(dni, resultado)
}

// Función de ordenación de Radix Sort para ordenar una lista de dnis
func radixSort(dni []eleccion.Votante) {
	maximo := maximoNumero(dni)

	for potencia := 1; maximo/potencia > 0; potencia *= 10 {
		ordenarPorDigito(dni, potencia)
	}
}

// Función para encontrar el maximo en una lista de dnis
func maximoNumero(padron []eleccion.Votante) int {
	maximo := padron[0].LeerDNI()
	for _, votante := range padron {
		numero := votante.LeerDNI()
		if numero > maximo {
			maximo = numero
		}
	}
	return maximo
}

// Realizo la busqueda binaria para ver si un dni pertenece al padron
func BusquedaBinariaEstaEnPadron(dni int, padron []eleccion.Votante) eleccion.Votante {
	return busqueda(dni, padron, 0, len(padron))
}

// busqueda binaria recursiva para ver si un dni se encuentra en el padron, devuelve nil si el dni no esta en el padron
func busqueda(dni int, padron []eleccion.Votante, inicio int, fin int) eleccion.Votante {
	if inicio >= fin {
		return nil
	}
	medio := (inicio + fin) / 2
	if padron[medio].LeerDNI() == dni {
		return padron[medio]
	} else if padron[medio].LeerDNI() > dni {
		return busqueda(dni, padron, inicio, medio)
	} else {
		return busqueda(dni, padron, medio+1, fin)
	}
}
