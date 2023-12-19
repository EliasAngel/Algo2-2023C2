package funciones_rerepolez

import (
	"errors"
	"fmt"
	"rerepolez/diseno_alumnos/errores"
	Elector "rerepolez/diseno_alumnos/votos"
	"strconv"
	TdaCola "tdas/cola"
)

// ConvertirAEntero() convierte una cadena a entero y devuelve el entero y un error en caso de que no se pueda realizar la operacion
func convertirAEntero(cadena string) (int, error) {
	entero, err := strconv.Atoi(cadena)
	if err != nil {
		return 0, err
	}
	return entero, nil
}

// esEntero() verifica que la cadena sea un entero y devuelve true en caso de que lo sea, false en caso contrario
func esEntero(cadena string) bool {
	_, err := convertirAEntero(cadena)
	if err != nil {
		return false
	}
	return true
}

// esDniValido() verifica que el dni sea valido y devuelve true en caso de que lo sea, false en caso contrario
func esDniValido(dni string) bool {
	//Verifico que el dni sea un numero
	dniE, err := convertirAEntero(dni)
	if err != nil {
		return false
	}
	if !esEntero(dni) {
		return false
	} else if dniE <= 0 && dniE >= 99999999 { //En Argetina recien se empiezan a usar los dni empezados en 70M
		return false
	}
	return true
}

// ingresarVotante() ingresa un votante al padron y devuelve un error en caso de que no se pueda realizar la operacion
func ingresarVotante(dni []string, padron []Elector.Votante) (Elector.Votante, error) {
	//Verifico si la entrada es la correcta
	if len(dni) != 2 {
		return nil, &errores.ErrorParametros{}
	}
	//Verifico el dni
	if !esDniValido(dni[1]) {
		return nil, &errores.DNIError{}
	}
	//Verifico que el votante este en el padron
	dniE, err := convertirAEntero(dni[1])
	if err != nil {
		return nil, err
	}
	//Si esta en el padron, creo el votante
	votante := BusquedaBinariaEstaEnPadron(dniE, padron)
	if votante != nil {
		return votante, nil
	}
	//Si no esta en el padron, devuelvo error
	return nil, &errores.DNIFueraPadron{}
}

// filaEstaVacia() devuelve true si la fila esta vacia, false en caso contrario
func filaEstaVacia(fila TdaCola.Cola[Elector.Votante]) bool {
	return fila.EstaVacia()
}

// VerificarNumLista() verifica que el numero de lista sea valido y devuelve un error en caso contrario
func verificarNumLista(numero int, partido []Elector.Partido) error {
	//Verifico que el numero de lista sea valido fijandome en el largo del partido
	if numero < 0 || numero > len(partido)-1 {
		return &errores.ErrorAlternativaInvalida{}
	}
	return nil
}

// VerificarElegibilidadParaVotar verifica que el votante actual en la fila este habilitado para votar, y devuelve un error en caso contrario
func VerificarElegibilidadParaVotar(tipoVoto string, fila TdaCola.Cola[Elector.Votante], numeroLista string, partido []Elector.Partido) (error, int) {
	//Verifico el la conversion de numeroLista
	numero, err := convertirAEntero(numeroLista)
	//Compruebo que la fila no este vacia
	if filaEstaVacia(fila) {
		return &errores.FilaVacia{}, -1
	}
	if !esTipoVoto(tipoVoto) {
		return &errores.ErrorTipoVoto{}, -1
	}
	if err != nil {
		return &errores.ErrorAlternativaInvalida{}, -1
	}
	// Verifico que el numero de lista sea valido
	error1 := verificarNumLista(numero, partido)
	if error1 != nil {
		return error1, -1
	}
	return nil, numero
}

// obtenerTipoVoto devuelve el tipo de voto segun el string pasado por parametro
func obtenerTipoVoto(tipoVoto string) (Elector.TipoVoto, error) {
	switch tipoVoto {
	case "Presidente":
		return Elector.PRESIDENTE, nil
	case "Gobernador":
		return Elector.GOBERNADOR, nil
	case "Intendente":
		return Elector.INTENDENTE, nil
	default:
		return -1, &errores.ErrorTipoVoto{}
	}
}

// ingresarVoto ingresa el voto del votante actual en la fila, y devuelve un error en caso de que no se pueda realizar el voto
func ingresarVoto(entrada []string, partido []Elector.Partido, fila TdaCola.Cola[Elector.Votante]) error {
	// Verifico si la entrada es la correcta
	if len(entrada) != 3 {
		return &errores.ErrorParametros{}
	}
	// Verifico la elegibilidad para votar
	err, numero := VerificarElegibilidadParaVotar(entrada[1], fila, entrada[2], partido)
	if err != nil {
		return err
	}
	// Realizo el voto
	tipoVoto := entrada[1]
	votante := fila.VerPrimero()
	tipoElegido, err := obtenerTipoVoto(tipoVoto)
	if err != nil {
		fila.Desencolar()
		return err
	}
	err = votante.Votar(tipoElegido, numero)
	if err != nil {
		fila.Desencolar()
		return err
	}
	return nil
}

// deshacerVoto() deshace el ultimo voto emitido por el votante actual en la fila. Si no hay votos para deshacer, se debe devolver error.
func deshacerVoto(fila TdaCola.Cola[Elector.Votante]) error {
	//Compruebo que la fila no este vacia
	if filaEstaVacia(fila) {
		return &errores.FilaVacia{}
	}
	votante := fila.VerPrimero()
	//Si no voto, deshago el voto
	err2 := votante.Deshacer()
	if errors.Is(err2, &errores.ErrorNoHayVotosAnteriores{}) {
		return err2
	} else if err2 != nil {
		fila.Desencolar()
		return err2
	}
	return nil
}

// aplicarVoto() Funcion para aplicar el voto al partido correspondiente
func aplicarVoto(partido []Elector.Partido, voto Elector.Voto, votoEnBlanco Elector.Partido) {
	for i, alternativa := range voto.VotoPorTipo {
		if alternativa == 0 {
			// Si no eligió ninguna alternativa, se contará como voto en blanco
			votoEnBlanco.VotadoPara(Elector.TipoVoto(i))
		} else {
			// Se contará como voto electo al partido correspondiente
			partido[alternativa].VotadoPara(Elector.TipoVoto(i))
		}
	}
}

// FinalizarVoto finaliza el voto del votante actual en la fila, y devuelve un error en caso de que no se pueda realizar el voto
func finalizarVoto(fila TdaCola.Cola[Elector.Votante], partido []Elector.Partido, impugnados *int, votoEnBlanco Elector.Partido, entrada []string) error {
	// Verificar si la entrada es válida
	if len(entrada) != 1 {
		return &errores.ErrorParametros{}
	}
	// Comprobar si la fila no está vacía
	if filaEstaVacia(fila) {
		return &errores.FilaVacia{}
	}
	// Finalizar el voto
	votante := fila.VerPrimero()
	voto, err := votante.FinVoto()
	if err != nil {
		fila.Desencolar()
		return err
	}
	// Aplicar el voto al partido
	if voto.Impugnado {
		*impugnados++
	} else {
		aplicarVoto(partido, voto, votoEnBlanco)
	}
	fila.Desencolar()
	return nil
}

// Funcion para mostrar los resultados de la votacion por tipo
func mostrarResultadoPorTipo(partidos []Elector.Partido, tipo Elector.TipoVoto, votoEnBlanco Elector.Partido) {
	fmt.Println(votoEnBlanco.ObtenerResultado(tipo))
	partidos = partidos[1:]
	for _, partido := range partidos {
		fmt.Println(partido.ObtenerResultado(tipo))
	}

}

// Funcion para mostrar los votos impugnados en el programa segun la cantidad de votos impugnados
func mostrarVotosImpugnados(impugnados int) {
	if impugnados == 1 {
		fmt.Println("Votos Impugnados:", impugnados, "voto")
	} else {
		fmt.Println("Votos Impugnados:", impugnados, "votos")
	}
}

// Funcion para mostrar los resultados de la votacion
func MostrarResultados(fila TdaCola.Cola[Elector.Votante], partidos []Elector.Partido, impugnados int, votoEnBlanco Elector.Partido) {
	//Muestro los resultados
	if !fila.EstaVacia() {
		fmt.Println(error(errores.ErrorCiudadanosSinVotar{}))
	}
	fmt.Println("Presidente:")
	mostrarResultadoPorTipo(partidos, Elector.PRESIDENTE, votoEnBlanco)
	fmt.Println()
	fmt.Println("Gobernador:")
	mostrarResultadoPorTipo(partidos, Elector.GOBERNADOR, votoEnBlanco)
	fmt.Println()
	fmt.Println("Intendente:")
	mostrarResultadoPorTipo(partidos, Elector.INTENDENTE, votoEnBlanco)
	fmt.Println()
	mostrarVotosImpugnados(impugnados)
}

// Funcion para verificar el tipo de voto
func esTipoVoto(tipo string) bool {
	return tipo == "Presidente" || tipo == "Gobernador" || tipo == "Intendente"
}

// Funcion para ingresar un votante al padron y devolver un error en caso de que no se pueda realizar la operacion
func comandoIngresar(entradaUsuario []string, padron []Elector.Votante, fila TdaCola.Cola[Elector.Votante]) {
	votante, err := ingresarVotante(entradaUsuario, padron)
	if err == nil {
		fmt.Println("OK")
		fila.Encolar(votante)
	} else {
		fmt.Println(err.Error())
	}
}

// Funcion para realizar el voto del votante actual en la fila, y devolver un error en caso de que no se pueda realizar el voto
func comandoVotar(entradaUsuario []string, partidos []Elector.Partido, fila TdaCola.Cola[Elector.Votante]) {
	err := ingresarVoto(entradaUsuario, partidos, fila)
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println(err.Error())
	}
}

// Funcion para deshacer el voto del votante actual en la fila, y devolver un error en caso de que no se pueda realizar el voto
func comandoDeshacer(entradaUsuario []string, fila TdaCola.Cola[Elector.Votante]) {
	err := deshacerVoto(fila)
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println(err.Error())
	}
}

// Funcion para finalizar el voto del votante actual en la fila, y devolver un error en caso de que no se pueda realizar el voto
func comandoFinVotar(entradaUsuario []string, partidos []Elector.Partido, fila TdaCola.Cola[Elector.Votante], impugnados *int, votoEnBlanco Elector.Partido) {
	err := finalizarVoto(fila, partidos, impugnados, votoEnBlanco, entradaUsuario)
	if err == nil {
		fmt.Println("OK")
	} else {
		fmt.Println(err.Error())
	}
}

// Funcion para procesar la entrada del usuario
func ProcesarEntrada(entradaUsuario []string, padron []Elector.Votante, partidos []Elector.Partido, fila TdaCola.Cola[Elector.Votante], votoEnBlanco Elector.Partido, impugnados *int) {
	switch entradaUsuario[0] {
	case "ingresar":
		// Funcionalidad de ingresar
		comandoIngresar(entradaUsuario, padron, fila)
	case "votar":
		// Funcionalidad de votar
		comandoVotar(entradaUsuario, partidos, fila)
	case "deshacer":
		// Funcionalidad de deshacer
		comandoDeshacer(entradaUsuario, fila)
	case "fin-votar":
		// Funcionalidad de fin-votar
		comandoFinVotar(entradaUsuario, partidos, fila, impugnados, votoEnBlanco)
	default:
		fmt.Println(&errores.ErrorComandoInvalido{})
	}
}
