package votos

import (
	"rerepolez/diseno_alumnos/errores"
	TDAPila "tdas/pila"
)

type votanteImplementacion struct {
	dni       int
	voto      Voto
	historial TDAPila.Pila[Voto]
	haVotado  bool
}

// CrearVotante genera un votante
func CrearVotante(dni int) Votante {
	return &votanteImplementacion{dni: dni, historial: TDAPila.CrearPilaDinamica[Voto]()}
}

// LeerDNI devuelve el dni del votante
func (votante votanteImplementacion) LeerDNI() int {
	return votante.dni
}

// Votar recibe el cargo y la lista que se quieren votar, y actualiza el voto segun corresponda
// En caso de que el votante ya haya votado, se devuelve el respectivo error
func (votante *votanteImplementacion) Votar(tipo TipoVoto, alternativa int) error {
	if votante.haVotado {
		return &errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	votante.historial.Apilar(votante.voto)
	if alternativa == LISTA_IMPUGNA {
		votante.voto.Impugnado = true
	} else {
		votante.voto.VotoPorTipo[tipo] = alternativa
	}
	return nil
}

// Deshacer deshace el ultimo voto emitido, y vuelve al anterior en caso de que exista,
// en caso de que no existan votos anteriores o el votante ya haya votado, se devuelve un error
func (votante *votanteImplementacion) Deshacer() error {
	if votante.haVotado {
		return &errores.ErrorVotanteFraudulento{Dni: votante.dni}
	}
	if votante.historial.EstaVacia() {
		return &errores.ErrorNoHayVotosAnteriores{}
	}
	votante.voto = votante.historial.Desapilar()
	return nil
}

// FinVoto devuelve el ultimo voto emitido, y actualiza el haVotado del votante
func (votante *votanteImplementacion) FinVoto() (Voto, error) {
	//El ultimo voto emitido es el que se considera valido
	if votante.haVotado {
		return votante.voto, &errores.ErrorVotanteFraudulento{votante.dni}
	}
	votante.haVotado = true
	return votante.voto, nil
}
