package main

import "tdas/diccionario"

type Resultado struct {
	pais      string
	resultado string // "v" si ganamos, "d" perdimos y "e" empatamos
}

type Resumen struct {
	pais    string
	ventaja int // de partidos de ventaja sobre el contrincante
}

func paternidad(resuls []Resultado) []Resumen {
	dicAux := diccionario.CrearHash[string, int]()
	suma := 0
	resultado := make([]Resumen, 0)
	for _, partido := range resuls {
		if dicAux.Pertenece(partido.pais) {
			if partido.resultado == "v" {
				suma = 1
			} else if partido.resultado == "d" {
				suma = -1
			}
			dicAux.Guardar(partido.pais, dicAux.Obtener(partido.pais)+suma)
		} else {
			dicAux.Guardar(partido.pais, suma)
		}
		suma = 0
	}
	iter := dicAux.Iterador()
	for iter.HaySiguiente() {
		rival, ventaja := iter.VerActual()
		resultado = append(resultado, Resumen{
			pais:    rival,
			ventaja: ventaja,
		})
		iter.Siguiente()
	}
	return resultado
}

func main() {

}
