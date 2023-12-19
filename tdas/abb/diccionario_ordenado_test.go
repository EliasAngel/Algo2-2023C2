package diccionario_test

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"strings"
	TDAABB "tdas/diccionario"
	"testing"
)

const (
	MAX_VOLUMEN = 1000
	MED_VOLUMEN = 100
)

var TAMS_ABB_VOLUMEN = []int{10, 100, 1000, 10000}

type PeliculasDisney struct {
	NombrePelicula string
	AñoEstreno     int
}

func comparar(a, b int) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

func compararStr(a, b string) int {
	if a < b {
		return -1
	} else if a > b {
		return 1
	}
	return 0
}

// compararPeliculas compara dos peliculas de Disney por su nombre, sin importar mayúsculas y minúsculas.
func compararPeliculas(p1, p2 PeliculasDisney) int {
	nombre1 := strings.ToLower(p1.NombrePelicula)
	nombre2 := strings.ToLower(p2.NombrePelicula)

	if nombre1 < nombre2 {
		return -1
	} else if nombre1 > nombre2 {
		return 1
	}
	return 0
}

////Test de la creacion de un abb

func TestABBVacio(t *testing.T) {
	t.Log("Comprueba que un abb vacio no tenga elementos")
	abb := TDAABB.CrearABB[int, int](comparar)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener(7) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar(7) })
}

func TestDiccionarioOrdenadoClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un abb vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abbint := TDAABB.CrearABB[int, int](comparar)
	require.EqualValues(t, 0, abbint.Cantidad())
	require.False(t, abbint.Pertenece(7))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbint.Obtener(7) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbint.Borrar(7) })
}

func TestUnElementoAbb(t *testing.T) {
	t.Log("Comprueba que un Abb con un elemento tiene esa Clave, unicamente")
	abbint := TDAABB.CrearABB[int, int](comparar)
	abbint.Guardar(11, 062011)
	require.EqualValues(t, 1, abbint.Cantidad())
	require.True(t, abbint.Pertenece(11))
	require.False(t, abbint.Pertenece(7))
	require.EqualValues(t, 062011, abbint.Obtener(11))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbint.Obtener(7) })
}

func TestAbbGuardarMuchos(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el abb, y se comprueba que en todo momento funciona acorde")
	claves := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	valores := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
	abbint := TDAABB.CrearABB[int, int](comparar)
	guardarClavesYValores(t, abbint, claves, valores)
}

func guardarClavesYValores(t *testing.T, abb TDAABB.DiccionarioOrdenado[int, int], claves, valores []int) {
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
		require.EqualValues(t, i+1, abb.Cantidad())
		require.EqualValues(t, valores[i], abb.Obtener(claves[i]))
	}
}

func sobrescribirValores(t *testing.T, abb TDAABB.DiccionarioOrdenado[int, int], claves, nuevosValores []int) {
	for i := range claves {
		abb.Guardar(claves[i], nuevosValores[i])
		require.EqualValues(t, len(claves), abb.Cantidad())
		require.EqualValues(t, nuevosValores[i], abb.Obtener(claves[i]))
	}
}

func TestSobrescribir_Guardar(t *testing.T) {
	t.Log("Guarda un par de claves y valores, luego sobrescribe los valores y verifica que se hayan reemplazado correctamente")
	claves := []int{1, 2, 3, 4, 5, 6}
	valores := []int{10, 20, 30, 40, 50, 60}
	abbint := TDAABB.CrearABB[int, int](comparar)
	guardarClavesYValores(t, abbint, claves, valores)
	nuevosValores := []int{100, 200, 300, 400, 500, 600}
	sobrescribirValores(t, abbint, claves, nuevosValores)
}

func TestAbbDiccionarioBorrar(t *testing.T) {
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	for i := len(claves) - 1; i >= 0; i-- {
		clave := claves[i]
		require.True(t, abbint.Pertenece(clave))
		valorBorrado := abbint.Borrar(clave)
		require.EqualValues(t, valores[i], valorBorrado)
		require.False(t, abbint.Pertenece(clave))
	}
	require.EqualValues(t, 0, abbint.Cantidad())
}

func TestAbbBorrarUnHijo(t *testing.T) {
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 6}
	valores := []int{100, 50, 20}
	guardarClavesYValores(t, abbint, claves, valores)
	abbint.Borrar(5)
	abbint.Borrar(10)
	abbint.Borrar(6)
	require.False(t, abbint.Pertenece(10))
	require.False(t, abbint.Pertenece(6))
	require.False(t, abbint.Pertenece(5))
}

func TestAbbReutlizacionDeBorrado(t *testing.T) {
	t.Log("Prueba de si se borra correctamente un elemento, y luego se vuelve a guardar, se sobreescribe el valor")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	nuevaClave := 5
	nuevoValor := 55
	abbint.Guardar(nuevaClave, nuevoValor)
	valorRecuperado := abbint.Obtener(nuevaClave)
	require.EqualValues(t, nuevoValor, valorRecuperado)
	require.EqualValues(t, len(claves), abbint.Cantidad())
}

func TestBusquedaDeMuchosElementos(t *testing.T) {
	t.Log("Prueba de busqueda de elementos en un abb")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	for i := range claves {
		require.True(t, abbint.Pertenece(claves[i]))
		require.EqualValues(t, valores[i], abbint.Obtener(claves[i]))
	}
}

//Test de Comparacion de Recorridos

func TestABBRecorridoDeElementosCompletos(t *testing.T) {
	t.Log("Comprueba que al insertar elementos en un ABB y luego iterar sobre ellos, se recorren todos los elementos en el orden correcto (inorden)")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	clavesEsperadas := []int{3, 5, 7, 10, 12, 15, 18}
	clavesOrdenadas := []int{}
	guardarClavesYValores(t, abbint, claves, valores)
	iter := abbint.Iterador()
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesOrdenadas = append(clavesOrdenadas, clave)
		iter.Siguiente()
	}
	require.Equal(t, clavesEsperadas, clavesOrdenadas)
}

func completarPeliculas() []PeliculasDisney {
	peliculas := []PeliculasDisney{
		{"Blancanieves y los 7 Enanitos", 2003},
		{"La Cenicienta", 2004},
		{"La Bella Durmiente", 2005},
		{"La Sirenita", 2006},
		{"Shrek 2", 2007},
		{"La Bella y la Bestia", 2008},
		{"Aladdin", 2009},
		{"El Rey León", 2010},
		{"Pocahontas", 2011},
		{"Mulan", 2012},
		{"Enredados", 2014},
		{"Frozen", 2015},
	}
	return peliculas
}

func TestABBClavesConStruct(t *testing.T) {
	t.Log("Prueba de busqueda de elementos en un abb con claves de tipo struct")
	abbstruct := TDAABB.CrearABB[PeliculasDisney, PeliculasDisney](compararPeliculas)
	peliculas := completarPeliculas()
	for pelicula := range peliculas {
		abbstruct.Guardar(peliculas[pelicula], peliculas[pelicula])
		require.EqualValues(t, pelicula+1, abbstruct.Cantidad())
		require.EqualValues(t, peliculas[pelicula], abbstruct.Obtener(peliculas[pelicula]))
	}
	for i := range peliculas {
		require.True(t, abbstruct.Pertenece(peliculas[i]))
		require.EqualValues(t, peliculas[i], abbstruct.Obtener(peliculas[i]))
	}
	require.EqualValues(t, len(peliculas), abbstruct.Cantidad())
}
func TestABBClavesConString(t *testing.T) {
	t.Log("Prueba de busqueda de elementos en un abb con claves de tipo string")
	abbstring := TDAABB.CrearABB[string, string](compararStr)
	claves := []string{"a", "b", "c", "d", "e", "f", "g"}
	valores := []string{"aa", "bb", "cc", "dd", "ee", "ff", "gg"}
	for i := range claves {
		abbstring.Guardar(claves[i], valores[i])
		require.EqualValues(t, i+1, abbstring.Cantidad())
		require.EqualValues(t, valores[i], abbstring.Obtener(claves[i]))
	}
	for i := range claves {
		require.True(t, abbstring.Pertenece(claves[i]))
		require.EqualValues(t, valores[i], abbstring.Obtener(claves[i]))
	}
	require.EqualValues(t, len(claves), abbstring.Cantidad())
}

// Test de iteradores internos
func TestABBIteradorInternoClaves(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas por el iterador interno")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	abbint.Iterar(func(clave int, valor int) bool {
		require.Contains(t, claves, clave)
		return true
	})
	require.EqualValues(t, len(claves), abbint.Cantidad())
	require.NotEqualValues(t, 0, abbint.Cantidad())
}

func TestABBIteradorInternoValores(t *testing.T) {
	t.Log("Valida que todos los valores sean recorridos por el iterador interno")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	factorial := 1
	f := &factorial
	abbint.Iterar(func(_ int, valor int) bool {
		*f *= valor
		return true
	})
	require.EqualValues(t, 100*50*150*30*70*120*180, factorial)
}

//Test de iteradores externos e internos vacios

func TestABBIteradorExternoEnABBVacio(t *testing.T) {
	t.Log("Comprueba que un iterador externo no se ejecute si el ABB está vacío")
	abbint := TDAABB.CrearABB[int, int](comparar)
	iter := abbint.Iterador()
	if iter.HaySiguiente() {
		panic("No debería ejecutarse")
	}
}

func TestABBIteradorInternoEnABBVacio(t *testing.T) {
	t.Log("Comprueba que un iterador interno no se ejecute si el ABB está vacío")
	abbint := TDAABB.CrearABB[int, int](comparar)
	var elementosIterados int
	abbint.Iterar(func(clave int, valor int) bool {
		elementosIterados++
		return true
	})
	if elementosIterados > 0 {
		panic("No debería ejecutarse")
	}
}

// Test de iteradores externos e internos

func TestABBIterarUnElemento(t *testing.T) {
	t.Log("Prueba de iterador interno, se agrega unos elemento y se verifica que se itere correctamente")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{1}
	valores := []int{10}
	guardarClavesYValores(t, abbint, claves, valores)
	iter := abbint.Iterador()
	require.True(t, iter.HaySiguiente())
	actual, _ := iter.VerActual()
	require.EqualValues(t, claves[0], actual)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorInterno(t *testing.T) {
	t.Log("Prueba de iterador interno, se agregan 1000 elementos y se verifica si el iterador funciona correctamente")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := make([]int, MED_VOLUMEN)
	valores := make([]int, MED_VOLUMEN)
	for i := 0; i < MED_VOLUMEN; i++ {
		claves[i] = i
		valores[i] = i * 10
	}
	guardarClavesYValores(t, abbint, claves, valores)
	abbint.Iterar(func(clave int, valor int) bool {
		require.Contains(t, claves, clave)
		return true
	})
}

func TestIteradorInternoCorte(t *testing.T) {
	t.Log("Itera el arbol inorder y corta la iteracion en un punto ''complicado''")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 4, 6, 15, 14, 16}
	valores := []int{4, 6, 8, 2, 6, 9, 1}
	guardarClavesYValores(t, abbint, claves, valores)
	suma := 0
	abbint.Iterar(func(clave int, valor int) bool {
		if clave <= 14 {
			suma += valor
			return true
		} else {
			return false
		}
	})
	require.EqualValues(t, 29, suma)
}

func TestIteradorInternoCortePrimerElemento(t *testing.T) {
	t.Log("Itera el primer elemento inorder y corta la iteracion")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 4, 6, 15, 14, 16}
	valores := []int{4, 6, 8, 2, 6, 9, 1}
	guardarClavesYValores(t, abbint, claves, valores)
	suma := 0
	abbint.Iterar(func(clave int, valor int) bool {
		suma += valor
		return false
	})
	require.EqualValues(t, 8, suma)
}

func TestIterarRangoCombinaciones(t *testing.T) {
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{3, 1, 2, 4, 5, 6, 7}
	valores := []int{4, 6, 8, 2, 6, 9, 1}
	guardarClavesYValores(t, abbint, claves, valores)
	inicio, final := 2, 5
	iter := abbint.IteradorRango(&inicio, &final)
	clave, _ := iter.VerActual()
	require.EqualValues(t, 2, clave)
}

func TestCalcularSumaDeTodosLosElementos(t *testing.T) {
	t.Log("Prueba de iterador interno, se agregan 100 elementos y se calcula la suma de todos los elementos")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := make([]int, MED_VOLUMEN)
	valores := make([]int, MED_VOLUMEN)
	for i := 0; i < MED_VOLUMEN; i++ {
		claves[i] = i
		valores[i] = i * 10
	}
	guardarClavesYValores(t, abbint, claves, valores)
	suma := 0
	abbint.Iterar(func(_ int, valor int) bool {
		suma += valor
		return true
	})
	require.EqualValues(t, 49500, suma)
}

//Test de iteradores externos e internos con y sin condicion de corte

func TestABBIteracionConCondicionesDeCorte(t *testing.T) {
	t.Log("Comprueba que la iteración interna en un árbol con condición de corte funciona correctamente")
	abb := TDAABB.CrearABB[int, int](comparar)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	caso := [2]int{}
	lista := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		if dato > 1000 {
			return false
		}
		lista = append(lista, clave)
		caso[0]++
		caso[1] = caso[1] + dato
		return true
	})
	require.NotContains(t, lista, 156687)
	require.Equal(t, 1, caso[0])
}

func TestABBIteracionSinCondicionesDeCorte(t *testing.T) {
	t.Log("Comprueba que la iteración interna en un árbol sin una condición de corte recorre todos los elementos")
	abb := TDAABB.CrearABB[int, int](comparar)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	for i := range claves {
		abb.Guardar(claves[i], valores[i])
	}
	caso := [2]int{}
	lista := []int{}
	abb.Iterar(func(clave int, dato int) bool {
		lista = append(lista, clave)
		caso[0]++
		caso[1] = caso[1] + dato
		return true
	})
	require.Equal(t, len(claves), caso[0])
}

//Test de iteradores externos e internos con rango

func TestABBIteradorInternoRango(t *testing.T) {
	t.Log("Comprueba que se pueda iterar correctamente dentro de un rango específico de claves en un ABB")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	inicio, fin := 5, 15
	clavesEnRango := []int{}
	funcion := func(clave int, _ int) bool {
		clavesEnRango = append(clavesEnRango, clave)
		return true
	}
	abbint.IterarRango(&inicio, &fin, funcion)
	clavesEsperadas := []int{5, 7, 10, 12, 15}
	require.Equal(t, clavesEsperadas, clavesEnRango)
}

func TestABBIteracionExternaConRango(t *testing.T) {
	t.Log("Comprueba que se pueda iterar externamente dentro de un rango específico de claves en un ABB")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 15, 3, 7, 12, 18}
	valores := []int{100, 50, 150, 30, 70, 120, 180}
	guardarClavesYValores(t, abbint, claves, valores)
	inicio, fin := 5, 15
	iter := abbint.IteradorRango(&inicio, &fin)
	clavesEnRango := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesEnRango = append(clavesEnRango, clave)
		iter.Siguiente()
	}
	clavesEsperadas := []int{5, 7, 10, 12, 15}
	require.Equal(t, clavesEsperadas, clavesEnRango)
}

func TestABBIteracionExternaSinDesde(t *testing.T) {
	t.Log("Comprueba que funcione correctamente el iterador si se pasa un solo valor del rango")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	guardarClavesYValores(t, abbint, claves, valores)
	fin := 25
	iter := abbint.IteradorRango(nil, &fin)
	clavesEnRango := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesEnRango = append(clavesEnRango, clave)
		iter.Siguiente()
	}
	clavesEsperadas := []int{1, 2, 3, 5, 6, 7, 11, 25}
	require.Equal(t, clavesEsperadas, clavesEnRango)
}

func TestABBIteracionExternaSinHasta(t *testing.T) {
	t.Log("Comprueba que funcione correctamente el iterador si se pasa un solo valor del rango")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{1, 432, 25, 5, 2, 3, 6, 7, 11, 71}
	valores := []int{152, 35, 523, 4, 15615, 156687, 156, 165, 12315, 0}
	guardarClavesYValores(t, abbint, claves, valores)
	inicio := 25
	iter := abbint.IteradorRango(&inicio, nil)
	clavesEnRango := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesEnRango = append(clavesEnRango, clave)
		iter.Siguiente()
	}
	clavesEsperadas := []int{25, 71, 432}
	require.Equal(t, clavesEsperadas, clavesEnRango)
}

func TestIteradorRangoCasoBorde(t *testing.T) {
	t.Log("Esta prueba valida que si tengo un ABB con preorder [10, 5, 7, 9], y creo un iterador por rangodesde 8, que empiece en 9 y siga con 10.")
	abbint := TDAABB.CrearABB[int, int](comparar)
	claves := []int{10, 5, 7, 9}
	valores := []int{1, 2, 3, 4}
	guardarClavesYValores(t, abbint, claves, valores)
	inicio := 8
	iter := abbint.IteradorRango(&inicio, nil)
	clavesEnRango := []int{}
	for iter.HaySiguiente() {
		clave, _ := iter.VerActual()
		clavesEnRango = append(clavesEnRango, clave)
		iter.Siguiente()
	}
	clavesEsperadas := []int{9, 10}
	require.Equal(t, clavesEsperadas, clavesEnRango)
}

//Test de Volumenes

func ejecutarPruebaVolumenABB(b *testing.B, n int) {
	abbint := TDAABB.CrearABB[int, int](comparar)

	claves := make([]int, n)
	valores := make([]int, n)

	for i := 0; i < n; i++ {
		valores[i] = i
		claves[i] = i
		abbint.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, abbint.Cantidad(), "La cantidad de elementos es incorrecta")

	ok := true
	for i := 0; i < n; i++ {
		ok = abbint.Pertenece(claves[i])
		if !ok {
			break
		}
		ok = abbint.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, abbint.Cantidad(), "La cantidad de elementos es incorrecta")

	for i := 0; i < n; i++ {
		ok = abbint.Borrar(claves[i]) == valores[i]
		if !ok {
			break
		}
		ok = !abbint.Pertenece(claves[i])
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, abbint.Cantidad())
}

func BenchmarkABB(b *testing.B) {
	b.Log("Prueba de stress del ABB. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves generadas, " +
		"y que luego podemos borrar sin problemas")

	for _, n := range TAMS_ABB_VOLUMEN {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenABB(b, n)
			}
		})
	}
}
