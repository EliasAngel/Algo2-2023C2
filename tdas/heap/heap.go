package cola_prioridad

const (
	LA_SEPTIMA   = 7
	TAM_INICIAL  = LA_SEPTIMA
	FACTOR_REDIM = 2
	LIMITE_REDIM = 4
)

type comparador[T any] func(T, T) int

type colaConPrioridad[T comparable] struct {
	datos []T
	cant  int
	cmp   comparador[T]
}

// CrearHeap crea una cola con prioridad vacía.
func CrearHeap[T comparable](cmp func(T, T) int) ColaPrioridad[T] {
	return &colaConPrioridad[T]{
		datos: make([]T, TAM_INICIAL),
		cmp:   cmp,
	}
}

// CrearHeapArr crea un heap a partir de un arreglo
func CrearHeapArr[T comparable](arreglo []T, cmp func(T, T) int) ColaPrioridad[T] {
	heap := &colaConPrioridad[T]{
		datos: make([]T, len(arreglo)+1),
		cmp:   cmp,
	}
	copy(heap.datos, arreglo)
	heap.cant = len(arreglo)
	heapify(heap.datos, heap.cant, cmp)
	return heap
}

// ------------------------------------ Primitivas del Heap---------------------------------------------------------

// EstaVacia devuelve true si la la cola se encuentra vacía, false en caso contrario.
func (h *colaConPrioridad[T]) EstaVacia() bool {
	return h.cant == 0
}

// Cantidad devuelve la cantidad de elementos que hay en la cola de prioridad.
func (h *colaConPrioridad[T]) Cantidad() int {
	return h.cant
}

// Encolar Agrega un elemento al heap.
func (h *colaConPrioridad[T]) Encolar(elem T) {
	if h.cant == len(h.datos) {
		h.redimensionar(len(h.datos) * FACTOR_REDIM)
	}
	h.datos[h.cant] = elem
	h.upheap(h.datos, h.cant, h.cmp)
	h.cant++
}

// Desencolar elimina el elemento con máxima prioridad, y lo devuelve. Si está vacía, entra en pánico con un
// mensaje "La cola esta vacia"
func (h *colaConPrioridad[T]) Desencolar() T {
	h.esVacioPanic()
	if h.cant <= len(h.datos)/LIMITE_REDIM && len(h.datos) > TAM_INICIAL {
		h.redimensionar(len(h.datos) / FACTOR_REDIM)
	}
	swap(0, h.cant-1, h)
	elementoEleminado := h.datos[h.cant-1]
	h.cant--
	downheap(h.datos, 0, h.cant, h.cmp)
	return elementoEleminado
}

// VerMax devuelve el elemento con máxima prioridad. Si está vacía, entra en pánico con un mensaje
// "La cola esta vacia".
func (h *colaConPrioridad[T]) VerMax() T {
	h.esVacioPanic()
	return h.datos[0]
}

// ------------------------------------ Primitivas Auxiliares Priivadas del Heap---------------------------------------------------------
// heapify reordena el arreglo para que cumpla la condicion de heap
func heapify[T comparable](arreglo []T, cant int, cmp comparador[T]) {
	for i := cant - 1; i >= 0; i-- {
		downheap(arreglo, i, cant, cmp)
	}
}

// esVacioPanic es una funcion auxiliar que se encarga de verificar si la cola esta vacia, en caso de estarlo
func (h *colaConPrioridad[T]) esVacioPanic() {
	if h.EstaVacia() {
		panic("La cola esta vacia")
	}
}

// downheap baja el elemento en la posicion posPadre hasta que se cumpla la condicion de heap
func downheap[T comparable](arreglo []T, posPadre int, cant int, cmp comparador[T]) {
	posHijoIzq := posPadre*2 + 1
	posHijoDer := posPadre*2 + 2
	posHijoMayor := posHijoIzq
	if posHijoIzq >= cant {
		return
	}
	if posHijoDer < cant && cmp(arreglo[posHijoDer], arreglo[posHijoIzq]) > 0 {
		posHijoMayor = posHijoDer
	}
	if cmp(arreglo[posHijoMayor], arreglo[posPadre]) > 0 {
		arreglo[posHijoMayor], arreglo[posPadre] = arreglo[posPadre], arreglo[posHijoMayor]
		downheap(arreglo, posHijoMayor, cant, cmp)
	}
}

// redimensionar redimensiona el arreglo de la cola con prioridad
func (h *colaConPrioridad[T]) redimensionar(nuevoTam int) {
	nuevoArreglo := make([]T, nuevoTam)
	copy(nuevoArreglo, h.datos)
	h.datos = nuevoArreglo
}

// calculoPosPadre calcula la posicion del padre de un hijo
func (h *colaConPrioridad[T]) calculoPosPadre(posHijo int) int {
	return (posHijo - 1) / 2
}

// swap intercambia los elementos en las posiciones pos1 y pos2
func swap[T comparable](pos1 int, pos2 int, h *colaConPrioridad[T]) {
	h.datos[pos1], h.datos[pos2] = h.datos[pos2], h.datos[pos1]
}

// padreQueHijo determina si el padre es mayor que el hijo
func (h *colaConPrioridad[T]) padreQueHijo(cmp func(T, T) int, arreglo []T, posHijo, posPadre int) bool {
	return cmp(arreglo[posPadre], arreglo[posHijo]) < 0
}

// upheap sube el elemento en la posicion posHijo hasta que se cumpla la condicion de heap
func (h *colaConPrioridad[T]) upheap(arreglo []T, posHijo int, cmp comparador[T]) {
	if posHijo == 0 {
		return
	}
	posPadre := h.calculoPosPadre(posHijo)
	if h.padreQueHijo(cmp, arreglo, posHijo, posPadre) {
		swap(posPadre, posHijo, h)
		h.upheap(arreglo, posPadre, cmp)
	}
}

// ------------------------------------ HeapSort ---------------------------------------------------------

// HeapSort ordena el arreglo usando el algoritmo HeapSort
func HeapSort[T comparable](elementos []T, cmp func(T, T) int) {
	heapify(elementos, len(elementos), cmp)
	for i := len(elementos) - 1; i >= 0; i-- {
		elementos[0], elementos[i] = elementos[i], elementos[0]
		downheap(elementos, 0, i, cmp)
	}
}
