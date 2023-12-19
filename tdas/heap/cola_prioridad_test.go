package cola_prioridad_test

import (
	"github.com/stretchr/testify/require"
	TDAHeap "tdas/heap"
	"testing"
)

var TAMS_VOLUMEN = []int{12500, 25000, 50000, 100000, 200000, 400000}

type PersonajeShrek struct {
	Nombre string
	Edad   int
}

func intCmp(a, b int) int {
	return a - b
}

func strCmp(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func personajeCmp(a, b PersonajeShrek) int {
	return a.Edad - b.Edad
}
func TestHeapVacio(t *testing.T) {
	t.Log("Crea un heap vacio y verifica que el mismo se comporta como tal")
	h := TDAHeap.CrearHeap(intCmp)
	require.True(t, h.EstaVacia())
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
}
func TestHeapArrVacio(t *testing.T) {
	t.Log("Crea un heap vacio a partir de un arreglo y verifica que el mismo se comporta como tal")
	h := TDAHeap.CrearHeapArr([]int{}, intCmp)
	require.True(t, h.EstaVacia())
	require.Equal(t, 0, h.Cantidad())
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.VerMax() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { h.Desencolar() })
}

func TestHeapEncolar(t *testing.T) {
	t.Log("Encola elementos en un heap y verifica que se encolen correctamente")
	h := TDAHeap.CrearHeap(intCmp)
	h.Encolar(1)
	require.Equal(t, 1, h.Cantidad())
	h.Encolar(2)
	require.Equal(t, 2, h.Cantidad())
	h.Encolar(3)
	require.Equal(t, 3, h.Cantidad())
	require.Equal(t, 3, h.VerMax())
}

func TestHeapDesencolar(t *testing.T) {
	t.Log("Encola elementos en un heap y verifica que se desencolen correctamente")
	h := TDAHeap.CrearHeap(intCmp)
	h.Encolar(1)
	h.Encolar(2)
	h.Encolar(3)
	require.Equal(t, 3, h.Desencolar())
	require.Equal(t, 2, h.Desencolar())
	require.Equal(t, 1, h.Desencolar())
	require.Equal(t, 0, h.Cantidad())
}

func TestHeapVolumenEncolar(t *testing.T) {
	t.Log("Encola muchos elementos en un heap y verifica que se encolen correctamente")
	for _, tam := range TAMS_VOLUMEN {
		h := TDAHeap.CrearHeap(intCmp)
		for i := 0; i < tam; i++ {
			h.Encolar(i)
		}
		require.Equal(t, tam, h.Cantidad())
	}
}

func TestHeapVolumenDesencolar(t *testing.T) {
	t.Log("Encola muchos elementos en un heap y verifica que se desencolen correctamente")
	for _, tam := range TAMS_VOLUMEN {
		h := TDAHeap.CrearHeap(intCmp)
		for i := 0; i < tam; i++ {
			h.Encolar(i)
		}
		for i := tam - 1; i >= 0; i-- {
			require.Equal(t, i, h.Desencolar())
		}
		require.Equal(t, 0, h.Cantidad())
	}
}

func TestHeapSort(t *testing.T) {
	t.Log("Ordena un arreglo de enteros utilizando heapsort")
	arreglo := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	arregloOrd := []int{1, 2, 3, 5, 6, 7, 11, 25, 71, 432}
	TDAHeap.HeapSort(arreglo, intCmp)
	require.Equal(t, arregloOrd, arreglo)
	//ahora de mayor a menor
	TDAHeap.HeapSort(arreglo, func(a, b int) int { return b - a })
	require.Equal(t, []int{432, 71, 25, 11, 7, 6, 5, 3, 2, 1}, arreglo)
}

func TestHeapSortVolumen(t *testing.T) {
	t.Log("Ordena arreglos de enteros de mayor tamaño utilizando heapsort")
	for _, tam := range TAMS_VOLUMEN {
		arreglo := make([]int, tam)
		for i := 0; i < tam; i++ {
			arreglo[i] = tam - i
		}
		TDAHeap.HeapSort(arreglo, intCmp)
		for i := 0; i < tam; i++ {
			require.Equal(t, i+1, arreglo[i])
		}
	}
}

func TestHeapPrioridadDuplicada(t *testing.T) {
	t.Log("Encola elementos con prioridad duplicada y verifica que se encolen correctamente")
	h := TDAHeap.CrearHeap(intCmp)
	h.Encolar(1)
	h.Encolar(1)
	h.Encolar(1)
	require.Equal(t, 3, h.Cantidad())
	require.Equal(t, 1, h.VerMax())
	require.Equal(t, 1, h.Desencolar())
	require.Equal(t, 1, h.VerMax())
	require.Equal(t, 1, h.Desencolar())
	require.Equal(t, 1, h.VerMax())
	require.Equal(t, 1, h.Desencolar())
	require.Equal(t, 0, h.Cantidad())
}

func TestHeapHeapArray(t *testing.T) {
	t.Log("Crea un heap a partir de un arreglo y verifica que se comporte correctamente")
	h := TDAHeap.CrearHeapArr([]int{1, 2, 3}, intCmp)
	require.Equal(t, 3, h.Cantidad())
	require.Equal(t, 3, h.VerMax())
	require.Equal(t, 3, h.Desencolar())
	require.Equal(t, 2, h.VerMax())
	require.Equal(t, 2, h.Desencolar())
	require.Equal(t, 1, h.VerMax())
	require.Equal(t, 1, h.Desencolar())
	require.Equal(t, 0, h.Cantidad())
}

func CrearPersonaje(nombre string, edad int) PersonajeShrek {
	return PersonajeShrek{
		Nombre: nombre,
		Edad:   edad,
	}
}

func CompletarPersonajes() []PersonajeShrek {
	personajes := []PersonajeShrek{
		CrearPersonaje("Shrek", 30),
		CrearPersonaje("Fiona", 28),
		CrearPersonaje("Burro", 25),
		CrearPersonaje("Lord Farquaad", 40),
		CrearPersonaje("Gingerbread Man", 5),
		CrearPersonaje("Dragon", 500),
	}
	return personajes
}

func TestHeapDePersonajes(t *testing.T) {
	t.Log("Crea un heap de personajes y verifica que se comporte correctamente")
	personajes := CompletarPersonajes()
	h := TDAHeap.CrearHeap(personajeCmp)
	for _, personaje := range personajes {
		h.Encolar(personaje)
	}
	require.Equal(t, 6, h.Cantidad())
	require.Equal(t, personajes[5], h.VerMax())
	require.Equal(t, personajes[5], h.Desencolar())
	require.Equal(t, personajes[3], h.VerMax())
	require.Equal(t, personajes[3], h.Desencolar())
	require.Equal(t, personajes[0], h.VerMax())
	require.Equal(t, personajes[0], h.Desencolar())
	require.Equal(t, personajes[1], h.VerMax())
	require.Equal(t, personajes[1], h.Desencolar())
	require.Equal(t, personajes[2], h.VerMax())
	require.Equal(t, personajes[2], h.Desencolar())
	require.Equal(t, personajes[4], h.VerMax())
	require.Equal(t, personajes[4], h.Desencolar())
	require.Equal(t, 0, h.Cantidad())
}

func TestHeapString(t *testing.T) {
	t.Log("Crea un heap de strings y verifica que se comporte correctamente")
	h := TDAHeap.CrearHeap(strCmp)
	entradas := []string{"Boca", "va", "a", "ganar", "la", "libertadores"}
	for _, entrada := range entradas {
		h.Encolar(entrada)
	}
	require.Equal(t, 6, h.Cantidad())
	require.Equal(t, "va", h.VerMax())
	require.Equal(t, "va", h.Desencolar())
	require.Equal(t, "libertadores", h.VerMax())
	require.Equal(t, "libertadores", h.Desencolar())
	require.Equal(t, "la", h.VerMax())
	require.Equal(t, "la", h.Desencolar())
	require.Equal(t, "ganar", h.VerMax())
	require.Equal(t, "ganar", h.Desencolar())
	require.Equal(t, "a", h.VerMax())
	require.Equal(t, "a", h.Desencolar())
	require.Equal(t, "Boca", h.VerMax())
	require.Equal(t, "Boca", h.Desencolar())
	require.Equal(t, 0, h.Cantidad())

}
