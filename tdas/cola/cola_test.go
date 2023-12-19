package cola_test

import (
	"github.com/stretchr/testify/require"
	TDACola "tdas/cola"
	"testing"
)

func TestColaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
	require.True(t, cola.EstaVacia())
	cola.Encolar(1)
	require.EqualValues(t, 1, cola.VerPrimero())
	cola.Encolar(4)
	require.EqualValues(t, 1, cola.VerPrimero())
}

func TestColaNoVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	cola.Encolar(1)
	require.False(t, cola.EstaVacia())
	require.EqualValues(t, 1, cola.VerPrimero())
	require.EqualValues(t, 1, cola.Desencolar())
	require.True(t, cola.EstaVacia())
}

func TestInvarianteColaMismo(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[int]()
	cola.Encolar(1)
	require.True(t, cola.VerPrimero() == 1)
	cola.Encolar(2)
	require.True(t, cola.VerPrimero() == 1)
	cola.Encolar(3)
	require.True(t, cola.VerPrimero() == 1)
	require.EqualValues(t, 1, cola.Desencolar())
	require.EqualValues(t, 2, cola.Desencolar())
	require.EqualValues(t, 3, cola.Desencolar())
}

func TestInvarianteColaDistinto(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	cola.Encolar("Hola")
	require.True(t, cola.VerPrimero() == "Hola")
	cola.Encolar(1)
	require.True(t, cola.VerPrimero() == "Hola")
	cola.Encolar([4]int{1, 2, 3, 4})
	require.True(t, cola.VerPrimero() == "Hola")
	require.EqualValues(t, cola.Desencolar(), "Hola")
	require.EqualValues(t, cola.Desencolar(), 1)
	require.EqualValues(t, cola.Desencolar(), [4]int{1, 2, 3, 4})
}

func TestVolumen(t *testing.T) {
	volumen := 10000
	cola := TDACola.CrearColaEnlazada[any]()
	for i := 0; i <= volumen; i++ {
		cola.Encolar(i)
		require.EqualValues(t, 0, cola.VerPrimero())
	}
	for i := 0; i <= volumen; i++ {
		require.EqualValues(t, cola.Desencolar(), i)
	}
	require.True(t, cola.EstaVacia())
}

func TestColaSeComportaVacia(t *testing.T) {
	cola := TDACola.CrearColaEnlazada[any]()
	for i := 0; i <= 5; i++ {
		cola.Encolar(i)
	}
	for !cola.EstaVacia() {
		cola.Desencolar() //Se vacía la cola
	}
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.Desencolar() })
	require.PanicsWithValue(t, "La cola esta vacia", func() { cola.VerPrimero() })
}

func helados([]struct{})  {

}