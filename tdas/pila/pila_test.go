package pila_test

import (
	"github.com/stretchr/testify/require"
	TDAPila "tdas/pila"
	"testing"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.True(t, pila.EstaVacia())
	pila.Apilar(1)
	require.EqualValues(t, 1, pila.VerTope())
	require.EqualValues(t, 1, pila.Desapilar())
	require.True(t, pila.EstaVacia())
}

func TestInvariantePilaMismo(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(1)
	require.True(t, pila.VerTope() == 1)
	pila.Apilar(2)
	require.True(t, pila.VerTope() == 2)
	pila.Apilar(3)
	require.True(t, pila.VerTope() == 3)
	require.EqualValues(t, pila.Desapilar(), 3)
	require.EqualValues(t, pila.Desapilar(), 2)
	require.EqualValues(t, pila.Desapilar(), 1)
}

func TestInvariantePilaDistinto(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()
	pila.Apilar("Hola")
	require.True(t, pila.VerTope() == "Hola")
	pila.Apilar(2)
	require.True(t, pila.VerTope() == 2)
	pila.Apilar([2]int{1, 2})
	require.True(t, pila.VerTope() == [2]int{1, 2})
	require.EqualValues(t, pila.Desapilar(), [2]int{1, 2})
	require.EqualValues(t, pila.Desapilar(), 2)
	require.EqualValues(t, pila.Desapilar(), "Hola")
}

func TestVolumen(t *testing.T) {
	volumen := 10000
	pila := TDAPila.CrearPilaDinamica[any]()
	for i := 0; i <= volumen; i++ {
		pila.Apilar(i)
		require.True(t, pila.VerTope() == i)
	}
	for i := volumen; i >= 0; i-- {
		require.EqualValues(t, pila.Desapilar(), volumen)
		volumen--
	}
	require.True(t, pila.EstaVacia())
}

func TestPilaSeComportaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[any]()
	for i := 0; i <= 5; i++ {
		pila.Apilar(i)
		require.True(t, pila.VerTope() == i)
	}
	for pila.EstaVacia() == false {
		pila.Desapilar() //Se vuelve a una pila vacía
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
}
