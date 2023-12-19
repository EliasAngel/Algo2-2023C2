package lista

// Creo un struct nodo que contiene un valor y un puntero al siguiente nodo.
type nodo[T any] struct {
	valor T
	sig   *nodo[T]
}

type listaEnlazada[T any] struct {
	primero *nodo[T]
	ultimo  *nodo[T]
	largo   int
}

// CrearNodo crea un nodo con el valor recibido y lo devuelve.
func crearNodo[T any](valor T) *nodo[T] {
	return &nodo[T]{valor, nil}
}

// TDA Lista Enlazada

// CrearListaEnlazada Crea una lista enlazada vacia
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{nil, nil, 0}
}

// EstaVacia bool devuelve true si la lista esta vacia en caso contrario devuelve false
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.largo == 0
}

// InsertarPrimero inserta un elemento al principio de la lista
func (lista *listaEnlazada[T]) InsertarPrimero(valor T) {
	nuevo := crearNodo(valor)
	if lista.EstaVacia() {
		lista.ultimo = nuevo
	} else {
		nuevo.sig = lista.primero
	}
	lista.primero = nuevo
	lista.largo++
}

// InsertarUltimo inserta un elemento al final de la lista
func (lista *listaEnlazada[T]) InsertarUltimo(valor T) {
	nuevo := crearNodo(valor)
	if lista.EstaVacia() {
		lista.primero = nuevo
	} else {
		lista.ultimo.sig = nuevo
	}
	lista.ultimo = nuevo
	lista.largo++
}

func (lista listaEnlazada[T]) listaVaciaPanic() {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
}

// BorrarPrimero borra el primer elemento de la lista, si la lista se encuentra vacia entra en panico con
// un mensaje de error "La lista esta vacia"
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	lista.listaVaciaPanic()
	dato := lista.primero.valor
	lista.primero = lista.primero.sig
	//Verificar si la lista quedo vacia
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return dato
}

// VerPrimero devuelve el primer elemento de la lista si la lista esta vacia entra en panico con un
// mensaje de error "La lista esta vacia"
func (lista listaEnlazada[T]) VerPrimero() T {
	lista.listaVaciaPanic()
	return lista.primero.valor
}

// VerUltimo devuelve el ultimo elemento de la lista si la lista esta vacia entra en panico con un
// mensaje de error "La lista esta vacia"
func (lista listaEnlazada[T]) VerUltimo() T {
	lista.listaVaciaPanic()
	return lista.ultimo.valor
}

// Largo devuelve el largo de la lista
func (lista listaEnlazada[T]) Largo() int {
	return lista.largo
}

// Iterar itera sobre la lista y ejecuta la funcion pasada por parametro en cada elemento de la lista si la
// funcion devuelve false se corta la iteracion
func (lista listaEnlazada[T]) Iterar(visitar func(T) bool) {
	actual := lista.primero
	for actual != nil {
		if !visitar(actual.valor) {
			break
		}
		actual = actual.sig
	}
}

func (lista listaEnlazada[T]) AnteKUltimo(k int) T {
	actual := lista.primero
	for i := 0; i <= k; i++ {
		actual = actual.sig
	}
	devolver := lista.primero
	actual = actual.sig
	for actual != nil {
		actual = actual.sig
		devolver = devolver.sig
	}
	return devolver.valor
}

// Iterador devuelve un iterador de la lista
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iteradorListaEnlazada[T]{lista.primero, nil, lista}
}

// TDA IteradorListaEnlazada
type iteradorListaEnlazada[T any] struct {
	actual   *nodo[T]
	anterior *nodo[T]
	lista    *listaEnlazada[T]
}

// iteradorTerminoPanic() si el iterador ya termino de iterar entra en panico con un mensaje de
// "El iterador termino de iterar"
func (iterador *iteradorListaEnlazada[T]) iteradorTerminoPanic() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

}

// VerActual() devuelve el elemento actual del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
// "El iterador termino de iterar"
func (iterador *iteradorListaEnlazada[T]) VerActual() T {
	iterador.iteradorTerminoPanic()
	return iterador.actual.valor
}

// HaySiguiente() devuelve true si el iterador tiene un elemento siguiente en caso contrario devuelve false
func (iterador *iteradorListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

// Siguiente() devuelve el siguiente elemento del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
// "El iterador termino de iterar"
func (iterador *iteradorListaEnlazada[T]) Siguiente() {
	iterador.iteradorTerminoPanic()
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.sig
}

// Insertar() inserta un elemento en la posicion actual si no esta vacio en caso contrario inserta el elemento al principio.
func (iterador *iteradorListaEnlazada[T]) Insertar(valor T) {
	nuevoNodo := crearNodo(valor)
	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.sig = nuevoNodo
	}
	nuevoNodo.sig = iterador.actual
	iterador.lista.largo++
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.actual = nuevoNodo
}

// Borrar() borra el elemento actual del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
// "El iterador termino de iterar"
func (iterador *iteradorListaEnlazada[T]) Borrar() T {
	iterador.iteradorTerminoPanic()
	valor := iterador.actual.valor
	if iterador.anterior == nil {
		//caso especial: borro el primer elemento
		iterador.lista.primero = iterador.actual.sig
		if iterador.lista.primero == nil {
			//caso especial: borro el ultimo elemento
			iterador.lista.ultimo = nil
		}
	} else {
		//caso general
		iterador.anterior.sig = iterador.actual.sig
		if iterador.actual.sig == nil {
			//caso especial: borro el ultimo elemento
			iterador.lista.ultimo = iterador.anterior
		}
	}
	//avanzo el iterador
	iterador.actual = iterador.actual.sig
	iterador.lista.largo--
	return valor
}
