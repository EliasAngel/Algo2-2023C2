package votos

import "fmt"

type partidoImplementacion struct {
	nombre     string
	candidatos [CANT_VOTACION]string
	votos      [CANT_VOTACION]int
}

type partidoEnBlanco struct {
	votosEnBlanco [CANT_VOTACION]int
}

// CrearPartido genera un partido a partir del nombre del mismo y sus candidatos
func CrearPartido(nombre string, candidatos [CANT_VOTACION]string) Partido {
	return &partidoImplementacion{nombre: nombre, candidatos: candidatos}
}

// CrearVotosEnBlanco genera un "partido" que cuenta los votos en blanco por cargo
func CrearVotosEnBlanco() Partido {
	return &partidoEnBlanco{}
}

// VotadoPara le añade un voto al candidato del partido para el cargo pasado por parametro
func (partido *partidoImplementacion) VotadoPara(tipo TipoVoto) {
	partido.votos[tipo]++
}

// ObtenerResultado obtiene el total de votos obtenidos por el candidato del partido
// para el cargo pasado por parametro
func (partido partidoImplementacion) ObtenerResultado(tipo TipoVoto) string {
	cantVotos := partido.votos[tipo]
	return fmt.Sprintf("%s - %s: %d %s", partido.nombre, partido.candidatos[tipo], cantVotos, esSingular(cantVotos))
}

// VotadoPara añade un voto en blanco para el cargo pasado por parametro
func (blanco *partidoEnBlanco) VotadoPara(tipo TipoVoto) {
	blanco.votosEnBlanco[tipo]++
}

// ObtenerResultado obtiene los votos en blanco para el cargo que se pasa por parametro
func (blanco partidoEnBlanco) ObtenerResultado(tipo TipoVoto) string {
	cantVotos := blanco.votosEnBlanco[tipo]
	return fmt.Sprintf("Votos en Blanco: %d %s", cantVotos, esSingular(cantVotos))
}

// esSingular() devuelve la palabra Votos en singular o plural segun corresponda
func esSingular(cantVotos int) string {
	if cantVotos == 1 {
		return "voto"
	} else {
		return "votos"
	}
}
