package lista_test

import (
	"github.com/stretchr/testify/require"
	TDALista "tdas/lista"
	"testing"
)

// Crear un struct de personajes de Shrek
type PersonajeShrek struct {
	nombre    string
	edad      int
	profesion string
	peliculas []string
}

const (
	MAXIMO = 1000000
)

// Función para inicializar un slice de personajes de Shrek
func CrearSlicePersonajesShrek() []PersonajeShrek {
	return []PersonajeShrek{
		{"Shrek", 35, "Ogro", []string{"Shrek", "Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Fiona", 30, "Princesa", []string{"Shrek", "Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Burro", 29, "Burro", []string{"Shrek", "Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Gato con Botas", 32, "Espadachín", []string{"Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Dragona", 100, "Dragón", []string{"Shrek", "Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Gingy", 5, "Galleta de Jengibre", []string{"Shrek", "Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Lord Farquaad", 40, "Humano", []string{"Shrek"}},
		{"Pinocchio", 10, "Marioneta", []string{"Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Los Hijos de Burro", 5, "Burros Traviesos", []string{"Shrek 2", "Shrek Tercero"}},
		{"Hada Madrina", 150, "Hada Madrina", []string{"Shrek 2", "Shrek Tercero"}},
		{"Príncipe Encantador", 28, "Príncipe", []string{"Shrek 2", "Shrek Tercero", "Shrek Para Siempre"}},
		{"Rumpelstiltskin", 150, "Duende", []string{"Shrek Para Siempre"}},
		{"Capitán Garfio", 50, "Pirata", []string{"Shrek Tercero"}},
		{"Cenicienta", 25, "Princesa", []string{"Shrek 2"}},
		{"Blancanieves", 15, "Princesa", []string{"Shrek 2"}},
		{"Rapunzel", 25, "Princesa", []string{"Shrek Tercero"}},
		{"Los Tres Cerditos", 5, "Cerditos", []string{"Shrek Tercero"}},
		{"Merlín", 8000, "Mago", []string{"Shrek Tercero"}},
	}
}

// Prueba si una lista está vacía y si se comporta como tal de manera generica para cualquier tipo de dato
func PruebaVacia[T any](t *testing.T, lista TDALista.Lista[T]) {
	require.Equal(t, true, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
	require.Panics(t, func() { lista.VerPrimero() })
	require.Panics(t, func() { lista.VerUltimo() })
	require.Panics(t, func() { lista.BorrarPrimero() })
}

// Crea una Lista Enlazada vacía, y prueba si esta se comporta como tal.
func TestListaVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaVacia(t, lista)
}

// Inserta un elemento al principio de la lista y prueba si la lista se comporta como tal.(Generica)
func PruebaInsertarPrimero[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem2)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, elem2, lista.VerPrimero())
}

// TestInsertarPrimeroInt() Inserta un elemento al principio de la lista y prueba si la lista se comporta como tal.(Prueba con enteros)
func TestInsertarPrimeroInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarPrimero(t, lista, 1, 2)
}

// TestInsertarPrimeroString() Inserta un elemento al principio de la lista y prueba si la lista se comporta como tal.(Prueba con cadenas)
func TestInsertarPrimeroString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarPrimero(t, lista, "Hola", "Mundo")
}

// TestInsertarPrimeroPersonajesShrek() Inserta un elemento al principio de la lista y prueba si la lista se comporta como tal.(Prueba con personajes de Shrek)
func TestInsertarPrimeroPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarPrimero(t, lista, personajes[0], personajes[1])
}

// Inserta un elemento al final de la lista y prueba si la lista se comporta como tal.(Prueba Generica)
func PruebaInsertarUltimo[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem1)
	lista.InsertarUltimo(elem2)
	require.False(t, lista.EstaVacia())
	require.Equal(t, 3, lista.Largo())
	require.Equal(t, elem2, lista.VerUltimo())
}

// TestInsertarUltimoInt() Inserta un elemento al final de la lista y prueba si la lista se comporta como tal.(Prueba con enteros)
func TestInsertarUltimoInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarUltimo(t, lista, 1, 2)
}

// TestInsertarUltimoString() Inserta un elemento al final de la lista y prueba si la lista se comporta como tal.(Prueba con cadenas)
func TestInsertarUltimoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarUltimo(t, lista, "Hola", "Mundo")
}

// TestInsertarUltimoPersonajesShrek() Inserta un elemento al final de la lista y prueba si la lista se comporta como tal.(Prueba con personajes de Shrek)
func TestInsertarUltimoPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarUltimo(t, lista, personajes[0], personajes[1])
}

// PruebaBorrarPrimerElemento borra el primer elemento de la lista , si la lista esta vacia debe lanzar el panic con el
// mensaje de error "La lista esta vacia" (Prueba Generica)
func PruebaBorrarPrimerElemento[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem2)
	lista.BorrarPrimero()
	require.False(t, lista.EstaVacia())
	require.Equal(t, 1, lista.Largo())
	require.Equal(t, elem1, lista.VerPrimero())
	lista.BorrarPrimero()
	require.True(t, lista.EstaVacia())
	require.Equal(t, 0, lista.Largo())
}

// TestBorrarPrimerElementoInt borra el primer elemento de la lista , si la lista esta vacia debe lanzar el panic con el
// mensaje de error "La lista esta vacia" (Prueba con enteros)
func TestBorrarPrimerElementoInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaBorrarPrimerElemento(t, lista, 1, 2)
}

// TestBorrarPrimerElementoString borra el primer elemento de la lista , si la lista esta vacia debe lanzar el panic con el
// mensaje de error "La lista esta vacia" (Prueba con cadenas)
func TestBorrarPrimerElementoString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaBorrarPrimerElemento(t, lista, "Hola", "Mundo")
}

// TestBorrarPrimerElementoPersonajesShrek() borra el primer elemento de la lista , si la lista esta vacia debe lanzar el panic con el
// mensaje de error "La lista esta vacia" (Prueba con personajes de Shrek)
func TestBorrarPrimerElementoPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaBorrarPrimerElemento(t, lista, personajes[0], personajes[1])
}

// InsertarMedio si la lista no esta vacia y tiene mas de dos elementos incerta un elemento en el medio de la lista
func InsertarMedio[T any](t *testing.T, lista TDALista.Lista[T], elem T) {
	//Inserta en el medio de la lista un elemento
	iter := lista.Iterador()
	for i := 0; i < lista.Largo(); i++ {
		iter.Siguiente()
	}
	iter.Insertar(elem)
	require.Equal(t, elem, iter.VerActual())
}

// PruebaInsertarMedio inserta un elemento en la posicion del medio de la lista y verifica si esta en la posicion correcta
// (Prueba Generica)
func PruebaInsertarMedio[T any](t *testing.T, lista TDALista.Lista[T], elem1, medio, elem3 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem3)
	lista.InsertarUltimo(elem3)
	InsertarMedio(t, lista, medio)
	require.False(t, lista.EstaVacia())
}

// TestInsertarMedioInt() inserta un elemento en la posicion del medio de la lista y verifica si esta en la posicion correcta (Prueba con enteros)
func TestInsertarMedioInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarMedio(t, lista, 1, 2, 3)
}

// TestInsertarMedioString() inserta un elemento en la posicion del medio de la lista y verifica si esta en la posicion correcta (Prueba con cadenas)
func TestInsertarMedioString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarMedio(t, lista, "Hola", "Mundo", "!")
}

// TestInsertarMedioPersonajesShrek() inserta un elemento en la posicion del medio de la lista y verifica si esta en la posicion correcta (Prueba con personajes de Shrek)
func TestInsertarMedioPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarMedio(t, lista, personajes[0], personajes[1], personajes[2])
}

// Prueba InsertarFinal inserta un elemento cuando el iterador esta al final verificando que es equivalente a insertar al final
func PruebaInsertarFinal[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2, elem3 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem2)
	iter := lista.Iterador()
	iter.Insertar(elem3)
	require.Equal(t, elem3, iter.VerActual())
}

// TestInsertarFinalInt() inserta un elemento cuando el iterador esta al final verificando que es equivalente a insertar al final (Prueba con enteros)
func TestInsertarFinalInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarFinal(t, lista, 1, 2, 3)
}

// TestInsertarFinalString inserta un elemento cuando el iterador esta al final verificando que es equivalente a insertar al final (Prueba con cadenas)
func TestInsertarFinalString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarFinal(t, lista, "Hola", "Mundo", "!")
}

// TestInsertarFinalPersonajesShrek inserta un elemento cuando el iterador esta al final verificando que es equivalente a insertar al final (Prueba con personajes de Shrek)
func TestInsertarFinalPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarFinal(t, lista, personajes[0], personajes[1], personajes[2])
}

// PruebaBorrarPrimerElementoIterador borra el primer elemento de la lista cuando se crea un iterador cambiandose el
// primer elemento de la lista.
func PruebaBorrarPrimerElementoIterador[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem2)
	iter := lista.Iterador()
	iter.Borrar()
	require.Equal(t, elem1, lista.VerPrimero())
}

// TestBorrarPrimerElementoIteradorInt borra el primer elemento de la lista cuando se crea un iterador cambiandose el
// primer elemento de la lista. (Prueba con enteros)
func TestBorrarPrimerElementoIteradorInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaBorrarPrimerElementoIterador(t, lista, 1, 2)
}

// TestBorrarPrimerElementoIteradorString borra el primer elemento de la lista cuando se crea un iterador cambiandose el
// primer elemento de la lista. (Prueba con cadenas)
func TestBorrarPrimerElementoIteradorString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaBorrarPrimerElementoIterador(t, lista, "Hola", "Mundo")
}

// TestBorrarPrimerElementoIteradorPersonajesShrek borra el primer elemento de la lista cuando se crea un iterador cambiandose el
// primer elemento de la lista. (Prueba con personajes de Shrek)
func TestBorrarPrimerElementoIteradorPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaBorrarPrimerElementoIterador(t, lista, personajes[0], personajes[1])
}

// PruebaInsertarNVeces() Prueba que inserta n veces un elemento en la lista (Prueba Generica)
func PruebaInsertarNVeces[T any](t *testing.T, lista TDALista.Lista[T], elem T, n int) {
	iter := lista.Iterador()
	for i := 0; i < n; i++ {
		iter.Insertar(elem)
	}
}

// PruebaBorrarNVeces() Prueba que borra n veces un elemento en la lista (Prueba Generica)
func PruebaBorrarNVeces[T any](t *testing.T, lista TDALista.Lista[T], n int) {
	iter := lista.Iterador()
	for i := 0; i < n; i++ {
		iter.Borrar()
	}
}

// PruebaVolumen() Prueba de volumen para verificar que la lista se comporta correctamente con una gran cantidad de datos (Prueba Generica)
func PruebaVolumen[T any](t *testing.T, lista TDALista.Lista[T], elem T, n int) {
	PruebaInsertarNVeces(t, lista, elem, n)
	require.Equal(t, n, lista.Largo())
	PruebaBorrarNVeces(t, lista, n)
	require.Equal(t, 0, lista.Largo())
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
}

// TestVolumenInt() Prueba de volumen para verificar que la lista se comporta correctamente con una gran cantidad de datos (Prueba con enteros)
func TestVolumenInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaVolumen(t, lista, 1, MAXIMO)
}

// TestVolumenString() Prueba de volumen para verificar que la lista se comporta correctamente con una gran cantidad de datos (Prueba con cadenas)
func TestVolumenString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaVolumen(t, lista, "Hola", MAXIMO)
}

// TestVolumenPersonajesShrek() Prueba de volumen para verificar que la lista se comporta correctamente con una gran cantidad de datos (Prueba con personajes de Shrek)
func TestVolumenPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaVolumen(t, lista, personajes[0], MAXIMO)
}

// PruebaBorrarUltimoElementoIterador borra el ultimo elemento de la lista con el iterador cambiandose el ultimo elemento de la lista.
func PruebaBorrarUltimoElementoIterador[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2, elem3 T) {
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem2)
	lista.InsertarUltimo(elem3)
	aux := 0

	iter := lista.Iterador()

	// Itera la lista
	for iter.HaySiguiente() {
		aux++
		if aux == lista.Largo() { //Si se llega a la ultima posicion se borra ese elemento
			iter.Borrar()
			break
		}
		iter.Siguiente()

	}
	require.Equal(t, elem2, lista.VerUltimo())
}

// PruebaBorrarUltimoElementoIterador borra el ultimo elemento de la lista con el iterador cambiandose el ultimo elemento de la lista. (Prueba con enteros)
func TestBorrarUltimoElementoIteradorInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaBorrarUltimoElementoIterador(t, lista, 1, 2, 3)
}

// PruebaBorrarUltimoElementoIterador borra el ultimo elemento de la lista con el iterador cambiandose el ultimo elemento de la lista. (Prueba con cadenas)
func TestBorrarUltimoElementoIteradorString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaBorrarUltimoElementoIterador(t, lista, "Hola", "Mundo", "!")
}

// PruebaBorrarUltimoElementoIterador borra el ultimo elemento de la lista con el iterador cambiandose el ultimo elemento de la lista. (Prueba con personajes de Shrek)
func TestBorrarUltimoElementoIteradorPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaBorrarUltimoElementoIterador(t, lista, personajes[0], personajes[1], personajes[2])
}

// PruebaInsertarPrimeroConIterador() inserta un elemento con el iterador en la primer posicion y comprueba que el primer elemento de la lista sea ese
func PruebaInsertarPrimeroConIterador[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2 T) {
	lista.InsertarPrimero(elem1)
	lista.InsertarPrimero(elem2)
	iter := lista.Iterador()
	iter.Insertar(elem1)
	require.Equal(t, elem1, lista.VerPrimero())
}

// TestInsertarPrimeroConIterador inserta un elemento con el iterador en la primer posicion y comprueba que el primer elemento de la lista sea ese (Prueba con enteros)
func TestInsertarPrimeroIteradorInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarPrimeroConIterador(t, lista, 1, 2)
}

// TestInsertarPrimeroConIterador inserta un elemento con el iterador en la primer posicion y comprueba que el primer elemento de la lista sea ese (Prueba con cadenas)
func TestInsertarPrimeroIteradorString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarPrimeroConIterador(t, lista, "Hola", "Mundo")
}

// TestInsertarPrimeroConIterador inserta un elemento con el iterador en la primer posicion y comprueba que el primer elemento de la lista sea ese (Prueba con personajes de Shrek)
func TestInsertarPrimeroIteradorPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarPrimeroConIterador(t, lista, personajes[0], personajes[1])
}

// PruebaInsertarUltimoIterador inserta un elemento con el iterador en la ultima posicion y comprueba que el ultimo elemento de la lista sea ese(Prueba Generica)
func PruebaInsertarUltimoIterador[T any](t *testing.T, lista TDALista.Lista[T], elem1, elem2, elem3 T) {
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem2)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(elem3)
	require.EqualValues(t, elem3, lista.VerUltimo())
}

// TestInsertarUltimoConIterador inserta un elemento con el iterador en la ultima posicion y comprueba que el ultimo elemento de la lista sea ese (Prueba con enteros)
func TestInsertarUltimoIteradorInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarUltimoIterador(t, lista, 1, 2, 3)
}

// TestInsertarUltimoConIterador inserta un elemento con el iterador en la ultima posicion y comprueba que el ultimo elemento de la lista sea ese (Prueba con cadenas)
func TestInsertarUltimoIteradorString(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[string]()
	PruebaInsertarUltimoIterador(t, lista, "Hola", "Mundo", "!")
}

// TestInsertarUltimoConIterador inserta un elemento con el iterador en la ultima posicion y comprueba que el ultimo elemento de la lista sea ese (Prueba con personajes de Shrek)
func TestInsertarUltimoIteradorPersonajesShrek(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[PersonajeShrek]()
	personajes := CrearSlicePersonajesShrek()
	PruebaInsertarUltimoIterador(t, lista, personajes[0], personajes[1], personajes[2])
}

// PruebaBorrarMedioIterador borra el elemento del medio de la lista con el iterador y comprueba que el elemento borrado no este en la lista
func PruebaBorrarMedioIterador(t *testing.T, lista TDALista.Lista[int], elem1, medio, elem3 int) {
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(medio)
	lista.InsertarUltimo(elem3)

	iter := lista.Iterador()

	for iter.HaySiguiente() {
		valor := iter.VerActual()
		if valor == medio {
			break
		}
		iter.Siguiente()
	}
	//Guardo el elemento que se va a borrar
	elemBorrado := iter.VerActual()

	//lo elimino
	iter.Borrar()

	//Verifico que el elemento borrado no este en la lista
	iter = lista.Iterador()
	for iter.HaySiguiente() {
		require.NotEqual(t, elemBorrado, iter.VerActual())
		iter.Siguiente()
	}

}

// TestBorrarMedioIteradorInt borra el elemento del medio de la lista con el iterador y comprueba que el elemento borrado no este en la lista (Prueba con enteros)
func TestBorrarMedioIteradorInt(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaBorrarMedioIterador(t, lista, 1, 2, 3)
}

// PruebaIteradorInterno() Prueba el iterador interno pasa una funcion que suma los pares de la lista y comprueba que el resultado sea el correcto (Prueba Enteros)
func PruebaIteradorInterno(t *testing.T, lista TDALista.Lista[int], elem1, elem2, elem3, elem4, sumaPares int) {
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem2)
	lista.InsertarUltimo(elem3)
	lista.InsertarUltimo(elem4)
	resultado := 0
	lista.Iterar(func(n int) bool { //le pedimos que sume los pares a la variable resultado
		if n%2 == 0 {
			resultado += n
		}
		return true
	})
	require.EqualValues(t, resultado, sumaPares)
}

// TestIteradorInterno() Prueba el iterador interno pasa una funcion que suma los pares de la lista y comprueba que el resultado sea el correcto (Prueba con enteros)
func TestIteradorInterno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaIteradorInterno(t, lista, 1, 2, 3, 4, 6)
}

// PruebaIteradorInternoCondicionCorte() Prueba el iterador interno pasa una funcion que suma los pares de la lista y comprueba que el resultado sea el correcto (Prueba con distintos tipos de datos)
func PruebaIteradorInternoCondicionCorte(t *testing.T, lista TDALista.Lista[any], elem1, elem2, elem3, elem4 interface{}, posEsperada int) {
	// Inserta elementos de diferentes tipos en la lista
	lista.InsertarUltimo(elem1)
	lista.InsertarUltimo(elem2)
	lista.InsertarUltimo(elem3)
	lista.InsertarUltimo(elem4)
	aux := 0
	lista.Iterar(func(elem interface{}) bool {
		if aux == posEsperada {
			return false // Termina la iteración cuando alcanza la posición esperada
		}
		aux++
		return true
	})
	require.EqualValues(t, posEsperada, aux)
}

// TestIteradorInternoCondicionCorte() Prueba el iterador interno pasa una funcion que suma los pares de la lista y comprueba que el resultado sea el correcto (Prueba con distintos tipos de datos)
func TestIteradorInternoCondicionCorte(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	PruebaIteradorInternoCondicionCorte(t, lista, 1, 6, [4]int{1, 2, 3, 4}, "Hola", 3)
}

// PruebaIteradorInternoCondicionCorte() Prueba el iterador interno pasa una funcion que suma los pares de la lista y comprueba que el resultado sea el correcto (Prueba con Mascotas)
func TestIteradorInternoCondicionCorteMascotas(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	mascotas := CrearSlicePersonajesShrek()
	PruebaIteradorInternoCondicionCorte(t, lista, mascotas[0], mascotas[1], mascotas[2], mascotas[3], 2)
}

// TestBorrarUno Prueba borrar un elemento de una lista con un solo elemento y la lista queda en un estado correcto
func TestBorrarUno(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	lista.InsertarPrimero(1)
	iter := lista.Iterador()
	iter.Borrar()
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.True(t, lista.EstaVacia())
}

// TestVaciarListaConIteradores() Prueba vaciar una lista con iteradores y la lista queda en un estado correcto
func TestVaciarListaConIteradores(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaInsertarNVeces(t, lista, 1, 3)
	largoInicial := lista.Largo()
	iter := lista.Iterador()
	for i := largoInicial; i > 0; i-- {
		aux := 1
		iter = lista.Iterador()
		for iter.HaySiguiente() {
			if aux == lista.Largo() {
				iter.Borrar()
				break
			}
			iter.Siguiente()
			aux++
		}
	}
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerUltimo() })
	require.PanicsWithValue(t, "La lista esta vacia", func() { lista.VerPrimero() })
	require.True(t, lista.EstaVacia())
	require.EqualValues(t, 0, lista.Largo())
	iter = lista.Iterador()
	iter.Insertar(1)
	require.True(t, lista.VerUltimo() == 1)
	require.True(t, lista.VerPrimero() == 1)
	require.EqualValues(t, 1, lista.Largo())

}

// TestInsertarUnElemListaNoVacia Prueba insertar un elemento en una lista no vacia y la lista queda en un estado correcto
func TestInsertarUnElemListaNoVacia(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[any]()
	lista.InsertarUltimo(2)
	lista.InsertarUltimo(3)
	iter := lista.Iterador()
	for iter.HaySiguiente() {
		iter.Insertar(1)
		break
	}
	require.EqualValues(t, 1, lista.VerPrimero())
}

// PruebaIteradorBorraListaConUnicoElemento Prueba borrar en una lista con un único elemento, y la lista queda en un estado correcto
// para insertar luego con primitivas de lista(Prueba Generica)
func PruebaIteradorBorrarConUnicoElemento[T any](t *testing.T, lista TDALista.Lista[T], elem T) {
	lista.InsertarPrimero(elem)
	iter := lista.Iterador()
	iter.Borrar()
	require.False(t, iter.HaySiguiente())
	iter.Insertar(elem)
	require.EqualValues(t, elem, lista.VerPrimero())
	require.False(t, lista.EstaVacia())
}

// TestIteradorBorraListaConUnicoElemento Prueba borrar en una lista con un único elemento, y la lista queda en un estado correcto
// para insertar luego con primitivas de lista
func TestIteradorBorraListaConUnicoElemento(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	PruebaIteradorBorrarConUnicoElemento(t, lista, 1)
}

func TestAnteKUltimo(t *testing.T) {
	lista := TDALista.CrearListaEnlazada[int]()
	for i := 1; i <= 10; i++ {
		lista.InsertarUltimo(i)
	}
	dato := lista.AnteKUltimo(4)
	require.Equal(t, dato, 4)
}
