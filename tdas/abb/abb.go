package diccionario

import (
	TDAPila "tdas/pila"
)

type abb[K comparable, V any] struct {
	raiz       *nodo[K, V]
	cantidad   int
	funcionCmp func(K, K) int
}

type nodo[K comparable, V any] struct {
	clave   K
	dato    V
	hijoIzq *nodo[K, V]
	hijoDer *nodo[K, V]
}

// crearNodo crea un nodo del arbol
func crearNodo[K comparable, V any](clave K, dato V) *nodo[K, V] {
	return &nodo[K, V]{clave: clave, dato: dato}
}

// CrearABB Inicializa un arbol vacio con la funcion de comparacion pasada por parametro
func CrearABB[K comparable, V any](funcionCmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{funcionCmp: funcionCmp}
}

// buscarNodo se encarga de buscar al nodo asociado a la clave que se pasa por parametro, y  al padre del mismo
func (abb *abb[K, V]) buscarNodo(act *nodo[K, V], padre *nodo[K, V], clave K) (*nodo[K, V], *nodo[K, V]) {
	if act == nil {
		return act, padre
	}
	if abb.funcionCmp(act.clave, clave) < 0 {
		return abb.buscarNodo(act.hijoDer, act, clave)
	} else if abb.funcionCmp(act.clave, clave) > 0 {
		return abb.buscarNodo(act.hijoIzq, act, clave)
	}
	return act, padre

}

// Guardar guarda el par clave-dato en el Diccionario. Si la clave ya se encontraba, se actualiza el dato asociado
func (abb *abb[K, V]) Guardar(clave K, dato V) {
	act, padre := abb.buscarNodo(abb.raiz, nil, clave)
	if act != nil {
		act.dato = dato
	} else {
		nuevo := crearNodo(clave, dato)
		if padre == nil {
			abb.raiz = nuevo
		} else if abb.funcionCmp(clave, padre.clave) < 0 {
			padre.hijoIzq = nuevo
		} else if abb.funcionCmp(clave, padre.clave) > 0 {
			padre.hijoDer = nuevo
		}
		abb.cantidad++
	}
}

// Pertenece determina si una clave ya se encuentra en el abb, o no
func (abb *abb[K, V]) Pertenece(clave K) bool {
	nodo, _ := abb.buscarNodo(abb.raiz, nil, clave)
	return nodo != nil
}

// Obtener devuelve el dato asociado a una clave. Si la clave no pertenece, debe entrar en pánico con mensaje
// 'La clave no pertenece al diccionario'
func (abb *abb[K, V]) Obtener(clave K) V {
	nodo, _ := abb.buscarNodo(abb.raiz, nil, clave)
	if nodo == nil {
		panic("La clave no pertenece al diccionario")
	}
	return nodo.dato
}

// cantHijos determina la cantidad de hijos que tiene un nodo
func (nodo *nodo[K, V]) cantHijos() int {
	resultado := 0
	if nodo.hijoDer != nil {
		resultado++
	}
	if nodo.hijoIzq != nil {
		resultado++
	}
	return resultado
}

// Borrar borra del Diccionario la clave indicada, devolviendo el dato que se encontraba asociado. Si la clave no
// pertenece al diccionario, debe entrar en pánico con un mensaje 'La clave no pertenece al diccionario'
func (abb *abb[K, V]) Borrar(clave K) V {
	act, padre := abb.buscarNodo(abb.raiz, nil, clave)
	if act == nil {
		panic("La clave no pertenece al diccionario")
	}
	dato := act.dato
	hijos := act.cantHijos()
	if hijos == 0 {
		abb.borrarAuxiliar(nil, padre, clave)
	}
	if hijos == 1 {
		if act.hijoIzq != nil {
			abb.borrarAuxiliar(act.hijoIzq, padre, act.clave)
		} else {
			abb.borrarAuxiliar(act.hijoDer, padre, act.clave)
		}
	}
	if hijos == 2 {
		abb.borrarConDosHijos(act)
	}
	abb.cantidad--
	return dato
}

// borrarAuxiliar se encarga de manejar el caso de borrar un nodo sin hijos o con un solo hijo
func (abb *abb[K, V]) borrarAuxiliar(hijo, padre *nodo[K, V], claveBorrado K) {
	if padre == nil {
		abb.raiz = hijo
	} else {
		if abb.funcionCmp(claveBorrado, padre.clave) < 0 {
			padre.hijoIzq = hijo
		} else if abb.funcionCmp(claveBorrado, padre.clave) > 0 {
			padre.hijoDer = hijo
		}
	}
}

// borrarConDosHijos se encarga de buscar un reemplazante para el nodo y reemplazar los datos con el nodo correspondiente
func (abb *abb[K, V]) borrarConDosHijos(act *nodo[K, V]) {
	reemplazante := abb.buscarReemplazante(act)
	claveReemplazar := reemplazante.clave
	datoReemplazar := abb.Borrar(reemplazante.clave)
	act.clave = claveReemplazar
	act.dato = datoReemplazar
	abb.cantidad++
}

// buscarReemplazante se mueve un hijo para la derecha y despues hacia la izquierda hasta llegar a un nodo sin hijos izquierdos
// para asi llegar a lo que seria el siguiente elemento inorder del arbol
func (abb *abb[K, V]) buscarReemplazante(act *nodo[K, V]) *nodo[K, V] {
	reemplazo := act.hijoDer
	for reemplazo.hijoIzq != nil {
		reemplazo = reemplazo.hijoIzq
	}
	return reemplazo
}

// Cantidad devuelve la cantidad de elementos dentro del hash
func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

// Iterar itera el arbol completo en recorrido INORDER (se procesa la raiz entre los nodos)
func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	abb.raiz.iterarAuxiliar(nil, nil, f, abb.funcionCmp)
}

// IterarRango itera solo incluyendo a los elementos que se encuentren comprendidos en el rango indicado,
// incluyéndolos en caso de encontrarse
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	abb.raiz.iterarAuxiliar(desde, hasta, visitar, abb.funcionCmp)
}

// iterarAuxiliar se encarga de que el arbol se recorra correctamente, respetando el rango recibido,
// verifica que cada nodo visitado devuelva true, y corta toda la iteracion en caso contrario
func (nodo *nodo[K, V]) iterarAuxiliar(desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if nodo == nil {
		return true
	}
	if nodo.hijoIzq != nil && (desde == nil || cmp(*desde, nodo.clave) <= 0) {
		if !nodo.hijoIzq.iterarAuxiliar(desde, hasta, visitar, cmp) {
			return false
		}
	}
	if (desde == nil || cmp(*desde, nodo.clave) <= 0) && (hasta == nil || cmp(*hasta, nodo.clave) >= 0) {
		if !visitar(nodo.clave, nodo.dato) {
			return false
		}
	}
	if nodo.hijoDer != nil && (hasta == nil || cmp(*hasta, nodo.clave) >= 0) {
		if !nodo.hijoDer.iterarAuxiliar(desde, hasta, visitar, cmp) {
			return false
		}
	}
	return true
}

// Implementacion del iterador externo del ABB
type iteradorABB[K comparable, V any] struct {
	arbol      abb[K, V]
	recorrido  TDAPila.Pila[*nodo[K, V]]
	funcionCmp func(K, K) int
	desde      *K
	hasta      *K
}

// Iterador devuelve un IterDiccionario para este Diccionario
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return abb.IteradorRango(nil, nil)
}

// IteradorRango crea un IterDiccionario que solo itere por las claves que se encuentren en el rango indicado
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iteradorABB[K, V])
	recorrido := TDAPila.CrearPilaDinamica[*nodo[K, V]]()
	iter.desde, iter.hasta = desde, hasta
	iter.funcionCmp = abb.funcionCmp
	iter.recorrido = iter.recorridoIterador(abb.raiz, recorrido)
	return iter
}

// HaySiguiente devuelve si hay más datos para ver. Esto es, si en el lugar donde se encuentra parado
// el iterador hay un elemento.
func (iter *iteradorABB[K, V]) HaySiguiente() bool {
	return !iter.recorrido.EstaVacia()
}

// VerActual devuelve la clave y el dato del elemento actual en el que se encuentra posicionado el iterador.
// Si no HaySiguiente, debe entrar en pánico con el mensaje 'El iterador termino de iterar'
func (iter *iteradorABB[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.recorrido.VerTope()
	return nodo.clave, nodo.dato
}

// Siguiente si HaySiguiente avanza al siguiente elemento en el hash. Si no HaySiguiente, entonces debe
// entrar en pánico con mensaje 'El iterador termino de iterar'
func (iter *iteradorABB[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.recorrido.Desapilar()
	if nodo.hijoDer != nil {
		iter.recorridoIterador(nodo.hijoDer, iter.recorrido)
	}
}

// recorridoIterador apila el nodo que se pasa por parametro y el subarbol izquierdo del mismo
func (iter *iteradorABB[K, V]) recorridoIterador(nodo *nodo[K, V], recorrido TDAPila.Pila[*nodo[K, V]]) TDAPila.Pila[*nodo[K, V]] {
	if nodo == nil {
		return recorrido
	}
	if (iter.desde == nil || iter.funcionCmp(*iter.desde, nodo.clave) <= 0) && (iter.hasta == nil || iter.funcionCmp(*iter.hasta, nodo.clave) >= 0) {
		recorrido.Apilar(nodo)
	} else if iter.desde == nil || iter.funcionCmp(*iter.desde, nodo.clave) >= 0 {
		recorrido = iter.recorridoIterador(nodo.hijoDer, recorrido)
	}
	recorrido = iter.recorridoIterador(nodo.hijoIzq, recorrido)
	return recorrido
}
