package lista

type Lista[T any] interface {
	//EstaVacia() bool devuelve true si la lista esta vacia en caso contrario devuelve false
	EstaVacia() bool

	//InsertarPrimero() inserta un elemento al principio de la lista y devuelve true si se inserto correctamente
	InsertarPrimero(T)

	//InsertarUltimo() inserta un elemento al final de la lista
	InsertarUltimo(T)

	//BorrarPrimero() borra el primer elemento de la lista si la lista se encuentra vacia entra en panico con
	//un mensaje de error "La lista esta vacia"
	BorrarPrimero() T

	//VerPrimero() devuelve el primer elemento de la lista si la lista esta vacia entra en panico con un
	//mensaje de error "La lista esta vacia"
	VerPrimero() T

	//VerUltimo() devuelve el ultimo elemento de la lista si la lista esta vacia entra en panico con un
	//mensaje de error "La lista esta vacia"
	VerUltimo() T

	//Largo() devuelve el largo de la lista
	Largo() int

	//Iterador() devuelve un iterador de la lista
	Iterador() IteradorLista[T]

	//Iterar() itera sobre la lista y ejecuta la funcion pasada por parametro en cada elemento de la lista si la
	//funcion devuelve false se corta la iteracion
	Iterar(visitar func(T) bool)

	AnteKUltimo(k int) T
}

type IteradorLista[T any] interface {

	//VerActual() devuelve el elemento actual del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
	//"El iterador termino de iterar"
	VerActual() T

	//HaySiguiente() devuelve true si el iterador tiene un elemento siguiente en caso contrario devuelve false
	HaySiguiente() bool

	//Siguiente() devuelve el siguiente elemento del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
	//"El iterador termino de iterar"
	Siguiente()

	// Insertar() inserta un elemento en la posicion actual si no esta vacio en caso contrario inserta el elemento al principio.
	Insertar(T)

	//Borrar() borra el elemento actual del iterador si el iterador ya termino de iterar entra en panico con un mensaje de
	//"El iterador termino de iterar"
	Borrar() T
}
