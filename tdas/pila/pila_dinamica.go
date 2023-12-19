package pila

const _CAP_INICIAL = 10
const _REDIMENSION = 2
const _COND_ACHIQUE = 4 //Si solo un cuarto de la pila está llena, se reduce su capacidad

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func (p *pilaDinamica[T]) EstaVacia() bool {
	return p.cantidad == 0
}

func (p *pilaDinamica[T]) VerTope() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	return p.datos[p.cantidad-1]
}

func (p *pilaDinamica[T]) Apilar(elem T) {
	if cap(p.datos) == p.cantidad {
		p.redimensionar(cap(p.datos) * _REDIMENSION)
	}
	p.datos[p.cantidad] = elem
	p.cantidad++
}

func (p *pilaDinamica[T]) Desapilar() T {
	if p.EstaVacia() {
		panic("La pila esta vacia")
	}
	if cap(p.datos) == p.cantidad*_COND_ACHIQUE {
		p.redimensionar(cap(p.datos) / _REDIMENSION)
	}
	p.cantidad--
	desapilado := p.datos[p.cantidad]
	return desapilado
}

func CrearPilaDinamica[T any]() Pila[T] {
	pila := new(pilaDinamica[T])
	pila.datos = make([]T, _CAP_INICIAL)
	pila.cantidad = 0
	return pila
}

func (p *pilaDinamica[T]) redimensionar(tam int) {
	pilaAmpliada := make([]T, tam)
	copy(pilaAmpliada, p.datos)
	p.datos = pilaAmpliada
}
