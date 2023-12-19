package diccionario

import (
	"fmt"
)

type estado int

const (
	VACIO estado = iota
	OCUPADO
	BORRADO
	LA_SEPTIMA       = 7
	TAM_INICIAL      = LA_SEPTIMA
	MAX_FACTOR_CARGA = 0.75
	MIN_FACTOR_CARGA = 0.25
	FACTOR_REDIM     = 2
)

type celdaHash[K comparable, V any] struct {
	clave  K
	dato   V
	estado estado
}

type hashCerrado[K comparable, V any] struct {
	tabla    []celdaHash[K, V]
	tam      int
	cant     int
	borrados int
}

type iteradorHashCerrado[K comparable, V any] struct {
	diccionario *hashCerrado[K, V]
	posActual   int
}

// convertirABytes convierte una clave a un slice de bytes
func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

// CrearTabla crea una tabla de hash cerrado
func crearTabla[K comparable, V any](nuevoT int) []celdaHash[K, V] {
	return make([]celdaHash[K, V], nuevoT)
}

// CrearHash crea un diccionario con diccionario cerrado
func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hashCerrado[K, V]{
		tabla: crearTabla[K, V](TAM_INICIAL),
		tam:   TAM_INICIAL,
	}
}

//Primitivas del diccionario

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (h *hashCerrado[K, V]) Guardar(clave K, dato V) {
	if (float64(h.cant+h.borrados) / float64(h.tam)) > MAX_FACTOR_CARGA {
		h.cambiarTamanio(FACTOR_REDIM * h.tam)
	}
	pos := h.buscarClave(clave)
	if h.tabla[pos].estado == VACIO {
		h.tabla[pos] = celdaHash[K, V]{clave, dato, OCUPADO}
		h.cant++
	}
	h.tabla[pos].dato = dato
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (h *hashCerrado[K, V]) Borrar(clave K) V {
	pos := h.buscarClave(clave)
	h.estaDisponiblePanic(pos)
	h.tabla[pos].estado = BORRADO
	h.cant--
	h.borrados++
	valor := h.tabla[pos].dato
	if (float64(h.cant)/float64(h.tam)) > MAX_FACTOR_CARGA && h.tam > TAM_INICIAL {
		h.cambiarTamanio(h.tam / FACTOR_REDIM)
	}
	return valor
}

// Pertenece determina si una clave ya se encuentra en el diccionario, o no
func (h *hashCerrado[K, V]) Pertenece(clave K) bool {
	pos := h.buscarClave(clave)
	return h.tabla[pos].estado == OCUPADO
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (h *hashCerrado[K, V]) Obtener(clave K) V {
	pos := h.buscarClave(clave)
	h.estaDisponiblePanic(pos)
	return h.tabla[pos].dato
}

// Cantidad devuelve la cantidad de elementos dentro del diccionario
func (h *hashCerrado[K, V]) Cantidad() int {
	return h.cant
}

// Iterar itera internamente el diccionario, aplicando la función pasada por parámetro a todos los elementos del
// mismo
func (h *hashCerrado[K, V]) Iterar(f func(clave K, dato V) bool) {
	for i := 0; i < h.tam; i++ {
		if h.tabla[i].estado == OCUPADO {
			if !f(h.tabla[i].clave, h.tabla[i].dato) {
				break
			}
		}
	}
}

// Iterador devuelve un IterDiccionario para este Diccionario
func (h *hashCerrado[K, V]) Iterador() IterDiccionario[K, V] {
	iter := &iteradorHashCerrado[K, V]{h, 0}
	if h.tabla[0].estado != OCUPADO {
		iter.posActual = iter.avanzarASiguienteOcupado()
	}
	return iter
}

// Funciones Auxiliares del diccionario

// hashFNV1a() es una función de hash no criptográfica que produce un hash de 64 bits.
// Referencias:
// http://www.isthe.com/chongo/tech/comp/fnv/index.html#FNV-1
// https://en.wikipedia.org/wiki/Fowler%E2%80%93Noll%E2%80%93Vo_hash_function
// https://golangprojectstructure.com/hash-functions-go-code/
func (h *hashCerrado[K, V]) hashFNV1a(dato []byte) uint64 {
	var hash, cavani uint64 = 0xcbf29ce484222325, 0x00000100000001b3
	for _, b := range dato {
		hash ^= uint64(b)
		hash *= cavani
	}
	return hash
}

// transferirElementos transfiere los elementos de la tabla anterior a la nueva tabla
func (h *hashCerrado[K, V]) transferirElementos(tablaAnterior []celdaHash[K, V]) {
	for _, celda := range tablaAnterior {
		if celda.estado == OCUPADO {
			h.Guardar(celda.clave, celda.dato)
		}
	}
}

// cambiarTamanio cambia el tamaño de la tabla de diccionario
func (h *hashCerrado[K, V]) cambiarTamanio(nuevoT int) {
	tabla_anterior := h.tabla
	nuevaTabla := crearTabla[K, V](nuevoT)
	h.tabla = nuevaTabla
	h.tam = nuevoT
	h.cant = 0
	h.transferirElementos(tabla_anterior)
}
func (h *hashCerrado[K, V]) estaDisponiblePanic(pos int) {
	if h.tabla[pos].estado == VACIO {
		panic("La clave no pertenece al diccionario")
	}
}

// buscarClave busca la clave en la tabla de diccionario y devuelve la posicion donde se encuentra
func (h *hashCerrado[K, V]) buscarClave(clave K) int {
	posInicial := int(h.hashFNV1a(convertirABytes(clave)) % uint64(h.tam))
	for !(h.tabla[posInicial].estado == VACIO) && !(h.tabla[posInicial].clave == clave && h.tabla[posInicial].estado == OCUPADO) {
		posInicial = (posInicial + 1) % h.tam
	}
	return posInicial
}

//TDA Iterador

//Primitivas del iterador

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
func (iter *iteradorHashCerrado[K, V]) HaySiguiente() bool {
	return iter.posActual < iter.diccionario.tam
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter *iteradorHashCerrado[K, V]) VerActual() (K, V) {
	iter.terminoDeIterar()
	celda := iter.diccionario.tabla[iter.posActual]
	return celda.clave, celda.dato
}

// Siguiente avanza al siguiente elemento en el diccionario. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador terminó de iterar'
func (iter *iteradorHashCerrado[K, V]) Siguiente() {
	iter.terminoDeIterar()
	iter.posActual = iter.avanzarASiguienteOcupado()
}

//Fuinciones Auxiliares de iterador

// avanzarASiguienteOcupado avanza el iterador a la siguiente posicion ocupada
func (iter *iteradorHashCerrado[K, V]) avanzarASiguienteOcupado() int {
	for i := iter.posActual + 1; i < iter.diccionario.tam; i++ {
		if iter.diccionario.tabla[i].estado == OCUPADO {
			return i
		}
	}
	return iter.diccionario.tam
}

// terminoDeIterar determina si el iterador termino de iterar
func (iter *iteradorHashCerrado[K, V]) terminoDeIterar() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
}
