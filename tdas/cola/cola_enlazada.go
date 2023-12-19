package cola

type nodoCola[T any] struct {
	dato T
	prox *nodoCola[T]
}

type colaEnlazada[T any] struct {
	primero *nodoCola[T]
	ultimo  *nodoCola[T]
}

func crearNodo[T any](dato T) *nodoCola[T] {
	nodo := new(nodoCola[T])
	nodo.prox = nil
	nodo.dato = dato
	return nodo
}

func CrearColaEnlazada[T any]() Cola[T] {
	cola := new(colaEnlazada[T])
	cola.primero = nil
	cola.ultimo = nil
	return cola
}
func (c *colaEnlazada[T]) EstaVacia() bool {
	return c.primero == nil
}

func (c *colaEnlazada[T]) VerPrimero() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	return c.primero.dato
}
func (c *colaEnlazada[T]) Encolar(elem T) {
	nuevo := crearNodo(elem)
	if c.EstaVacia() {
		c.primero = nuevo
	} else {
		c.ultimo.prox = nuevo
	}
	c.ultimo = nuevo
}

func (c *colaEnlazada[T]) Desencolar() T {
	if c.EstaVacia() {
		panic("La cola esta vacia")
	}
	elem := c.primero.dato
	if c.primero == c.ultimo { //Cuando queda un solo elemento
		c.ultimo = nil
	}
	c.primero = c.primero.prox
	return elem
}

func (c *colaEnlazada[T]) Multiprimeros(k int) []T {
	resultado := make([]T, k)
	for i := 0; i < k; i++ {
		if c.primero == nil {
			return resultado[:i]
		}
		resultado[i] = c.Desencolar()
	}
	return resultado
}
