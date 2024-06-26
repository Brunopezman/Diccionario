package diccionario_test

import (
	"fmt"
	"math/rand"
	"strings"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

var TAMS_VOLUMEN_ABB = []int{12500, 25000, 50000, 100000, 200000, 400000}

func func_cmp_int(e1, e2 int) int {
	if e1 == e2 {
		return 0
	}
	if e1 > e2 {
		return 1
	}
	return -1
}

func func_cmp_str(s1, s2 string) int {
	if len(s1) == len(s2) {
		return 0
	}
	if len(s1) > len(s2) {
		return 1
	}
	return -1
}
func TestDiccionarioAbbVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[string, int](func_cmp_str)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece("A"))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("A") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("A") })
}
func TestDiccionarioAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un ABB vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	abb := TDADiccionario.CrearABB[string, string](func_cmp_str)
	require.False(t, abb.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abb.Borrar("") })

	abbNum := TDADiccionario.CrearABB[int, string](func_cmp_int)
	require.False(t, abbNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { abbNum.Borrar(0) })
}

func TestUnElementoAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	dic.Guardar(1, 10)
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(1))
	require.False(t, dic.Pertenece(2))
	require.EqualValues(t, 10, dic.Obtener(1))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(2) })
}

func TestDiccionarioAbbGuardar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se comprueba que en todo momento funciona acorde")
	clave1 := 1
	clave2 := 2
	clave3 := 3
	valor1 := 100
	valor2 := 200
	valor3 := 300
	claves := []int{clave1, clave2, clave3}
	valores := []int{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	require.EqualValues(t, 1, dic.Cantidad())
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))

	require.False(t, dic.Pertenece(claves[1]))
	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[1], valores[1])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))

	require.False(t, dic.Pertenece(claves[2]))
	dic.Guardar(claves[2], valores[2])
	require.True(t, dic.Pertenece(claves[0]))
	require.True(t, dic.Pertenece(claves[1]))
	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, 3, dic.Cantidad())
	require.EqualValues(t, valores[0], dic.Obtener(claves[0]))
	require.EqualValues(t, valores[1], dic.Obtener(claves[1]))
	require.EqualValues(t, valores[2], dic.Obtener(claves[2]))
}

func TestReemplazoDatoAbb(t *testing.T) {
	t.Log("Guarda un par de claves, y luego vuelve a guardar, buscando que el dato se haya reemplazado")
	clave := 1
	clave2 := 2
	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	dic.Guardar(clave, 100)
	dic.Guardar(clave2, 200)
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 100, dic.Obtener(clave))
	require.EqualValues(t, 200, dic.Obtener(clave2))
	require.EqualValues(t, 2, dic.Cantidad())

	dic.Guardar(clave, 10)
	dic.Guardar(clave2, 20)
	require.True(t, dic.Pertenece(clave))
	require.True(t, dic.Pertenece(clave2))
	require.EqualValues(t, 2, dic.Cantidad())
	require.EqualValues(t, 10, dic.Obtener(clave))
	require.EqualValues(t, 20, dic.Obtener(clave2))
}
func TestReemplazoDatoHopscotchAbb(t *testing.T) {
	t.Log("Guarda bastantes claves, y luego reemplaza sus datos. Luego valida que todos los datos sean " +
		"correctos. Para una implementación Hopscotch, detecta errores al hacer lugar o guardar elementos.")

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	for i := 0; i < 500; i++ {
		dic.Guardar(i, i)
	}
	for i := 0; i < 500; i++ {
		dic.Guardar(i, 2*i)
	}
	ok := true
	for i := 0; i < 500 && ok; i++ {
		ok = dic.Obtener(i) == 2*i
	}
	require.True(t, ok, "Los elementos no fueron actualizados correctamente")
}

func TestDiccionarioAbbBorrar(t *testing.T) {
	t.Log("Guarda algunos pocos elementos en el diccionario, y se los borra, revisando que en todo momento " +
		"el diccionario se comporte de manera adecuada")
	clave1 := 1
	clave2 := 2
	clave3 := 3
	valor1 := 100
	valor2 := 200
	valor3 := 300
	claves := []int{clave1, clave2, clave3}
	valores := []int{valor1, valor2, valor3}

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	require.False(t, dic.Pertenece(claves[0]))
	require.False(t, dic.Pertenece(claves[0]))
	dic.Guardar(claves[0], valores[0])
	dic.Guardar(claves[1], valores[1])
	dic.Guardar(claves[2], valores[2])

	require.True(t, dic.Pertenece(claves[2]))
	require.EqualValues(t, valores[2], dic.Borrar(claves[2]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[2]) })
	require.EqualValues(t, 2, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[2]))

	require.True(t, dic.Pertenece(claves[0]))
	require.EqualValues(t, valores[0], dic.Borrar(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[0]) })
	require.EqualValues(t, 1, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[0]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[0]) })

	require.True(t, dic.Pertenece(claves[1]))
	require.EqualValues(t, valores[1], dic.Borrar(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(claves[1]) })
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(claves[1]))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(claves[1]) })
}

func TestReutlizacionDeBorradosAbb(t *testing.T) {
	t.Log("Prueba de caja blanca: reinsertando un elemento borrado")
	abb := TDADiccionario.CrearABB[string, string](func_cmp_str)
	clave := "hola"
	abb.Guardar(clave, "mundo!")
	abb.Borrar(clave)
	require.EqualValues(t, 0, abb.Cantidad())
	require.False(t, abb.Pertenece(clave))
	abb.Guardar(clave, "mundooo!")
	require.True(t, abb.Pertenece(clave))
	require.EqualValues(t, 1, abb.Cantidad())
	require.EqualValues(t, "mundooo!", abb.Obtener(clave))
}

func TestGuardarYBorrarRepetidasVecesAbb(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar ")

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i)
		require.True(t, dic.Pertenece(i))
		dic.Borrar(i)
		require.False(t, dic.Pertenece(i))
	}
}
func buscarInt(clave int, claves []int) int {
	for i, c := range claves {
		if c == clave {
			return i
		}
	}
	return -1
}
func TestIteradorInternoClavesABB(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := 1
	clave2 := 2
	clave3 := 3
	claves := []int{clave1, clave2, clave3}
	dic := TDADiccionario.CrearABB[int, *int](func_cmp_int)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)

	cs := []int{0, 0, 0}
	cantidad := 0
	cantPtr := &cantidad

	dic.Iterar(func(clave int, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarInt(cs[0], claves))
	require.NotEqualValues(t, -1, buscarInt(cs[1], claves))
	require.NotEqualValues(t, -1, buscarInt(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

func TestIteradorInternoValoresABB(t *testing.T) {
	t.Log("Valida que los datos sean recorridas correctamente (y una única vez) con el iterador interno")
	clave1 := 1
	clave2 := 2
	clave3 := 3
	clave4 := 4
	clave5 := 5

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	dic.Guardar(clave1, 6)
	dic.Guardar(clave2, 2)
	dic.Guardar(clave3, 3)
	dic.Guardar(clave4, 4)
	dic.Guardar(clave5, 5)

	factorial := 1
	ptrFactorial := &factorial
	dic.Iterar(func(_ int, dato int) bool {
		*ptrFactorial *= dato
		return true
	})

	require.EqualValues(t, 720, factorial)
}

func TestPruebaIterarTrasBorradosABB(t *testing.T) {

	clave1 := 1
	clave2 := 2
	clave3 := 3

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)
	dic.Guardar(clave1, 100)
	dic.Guardar(clave2, 200)
	dic.Guardar(clave3, 300)
	dic.Borrar(clave1)
	dic.Borrar(clave2)
	dic.Borrar(clave3)
	iter := dic.Iterador()

	require.False(t, iter.HaySiguiente())
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	dic.Guardar(clave1, 10)
	iter = dic.Iterador()

	require.True(t, iter.HaySiguiente())
	c1, v1 := iter.VerActual()
	require.EqualValues(t, clave1, c1)
	require.EqualValues(t, 10, v1)
	iter.Siguiente()
	require.False(t, iter.HaySiguiente())
}

func TestVolumenIteradorCorteABB(t *testing.T) {
	t.Log("Prueba de volumen de iterador interno, para validar que siempre que se indique que se corte" +
		" la iteración con la función visitar, se corte")

	dic := TDADiccionario.CrearABB[int, int](func_cmp_int)

	/* Inserta 'n' parejas en el abb */
	for i := 0; i < 10000; i++ {
		dic.Guardar(i, i)
	}

	seguirEjecutando := true
	siguioEjecutandoCuandoNoDebia := false

	dic.Iterar(func(c int, v int) bool {
		if !seguirEjecutando {
			siguioEjecutandoCuandoNoDebia = true
			return false
		}
		if c%100 == 0 {
			seguirEjecutando = false
			return false
		}
		return true
	})

	require.False(t, seguirEjecutando, "Se tendría que haber encontrado un elemento que genere el corte")
	require.False(t, siguioEjecutandoCuandoNoDebia,
		"No debería haber seguido ejecutando si encontramos un elemento que hizo que la iteración corte")
}

func ejecutarPruebaVolumenAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)
	clavesOrd := make([]string, n)
	valoresOrd := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		clavesOrd[i] = fmt.Sprintf("%08d", i)
		valoresOrd[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})
	/* Inserta 'n' parejas en el ABB */
	for i := range claves {
		dic.Guardar(claves[i], valores[i])
	}

	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que devuelva los valores correctos */
	ok := true
	for i := 0; i < n; i++ {
		ok = dic.Pertenece(clavesOrd[i])
		require.EqualValues(b, dic.Obtener(clavesOrd[i]), valoresOrd[i])
		if !ok {
			break
		}
		ok = dic.Obtener(claves[i]) == valores[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Pertenece y Obtener con muchos elementos no funciona correctamente")
	require.EqualValues(b, n, dic.Cantidad(), "La cantidad de elementos es incorrecta")

	/* Verifica que borre y devuelva los valores correctos */
	for i := 0; i < n; i++ {
		ok = dic.Borrar(clavesOrd[i]) == valoresOrd[i]
		if !ok {
			break
		}
	}

	require.True(b, ok, "Borrar muchos elementos no funciona correctamente")
	require.EqualValues(b, 0, dic.Cantidad())
}

func BenchmarkDiccionarioOrdenado(b *testing.B) {
	b.Log("Prueba de stress del Diccionario. Prueba guardando distinta cantidad de elementos (muy grandes), " +
		"ejecutando muchas veces las pruebas para generar un benchmark. Valida que la cantidad " +
		"sea la adecuada. Luego validamos que podemos obtener y ver si pertenece cada una de las claves geeneradas, " +
		"y que luego podemos borrar sin problemas")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebaVolumenAbb(b, n)
			}
		})
	}
}

func ejecutarPruebasVolumenIteradorAbb(b *testing.B, n int) {
	dic := TDADiccionario.CrearABB[string, *int](strings.Compare)

	claves := make([]string, n)
	valores := make([]int, n)
	clavesOrd := make([]string, n)
	valoresOrd := make([]int, n)

	for i := 0; i < n; i++ {
		claves[i] = fmt.Sprintf("%08d", i)
		valores[i] = i
		clavesOrd[i] = fmt.Sprintf("%08d", i)
		valoresOrd[i] = i
	}
	rand.Shuffle(n, func(i, j int) {
		claves[i], claves[j] = claves[j], claves[i]
		valores[i], valores[j] = valores[j], valores[i]
	})
	/* Inserta 'n' parejas en el ABB */
	for i := 0; i < n; i++ {
		dic.Guardar(claves[i], &valores[i])
	}

	// Prueba de iteración sobre las claves almacenadas.
	iter := dic.Iterador()
	require.True(b, iter.HaySiguiente())

	ok := true
	var i int
	var clave string
	var valor *int

	for i = 0; i < n; i++ {
		if !iter.HaySiguiente() {
			ok = false
			break
		}
		c1, v1 := iter.VerActual()
		clave = c1
		valor = v1
		if clave == "" || valor == nil || clavesOrd[i] != clave || valoresOrd[i] != *valor {
			ok = false
			break
		}
		*valor = n
		iter.Siguiente()
	}
	require.True(b, ok, "Iteracion en volumen no funciona correctamente")
	require.EqualValues(b, n, i, "No se recorrió todo el largo")
	require.False(b, iter.HaySiguiente(), "El iterador debe estar al final luego de recorrer")

	ok = true
	for i = 0; i < n; i++ {
		if valores[i] != n {
			ok = false
			break
		}
	}
	require.True(b, ok, "No se cambiaron todos los elementos")
}

func BenchmarkIteradorAbb(b *testing.B) {
	b.Log("Prueba de stress del Iterador del Diccionario. Prueba guardando distinta cantidad de elementos " +
		"(muy grandes) b.N elementos, iterarlos todos sin problemas. Se ejecuta cada prueba b.N veces para generar " +
		"un benchmark")
	for _, n := range TAMS_VOLUMEN_ABB {
		b.Run(fmt.Sprintf("Prueba %d elementos", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				ejecutarPruebasVolumenIteradorAbb(b, n)
			}
		})
	}
}

func TestIteradorInternoRangoClavesABB(t *testing.T) {
	t.Log("Valida que todas las claves sean recorridas (y una única vez) con el iterador interno")
	clave1 := 1
	clave2 := 2
	clave3 := 3
	clave4 := 4
	clave5 := 5
	claves := []int{clave1, clave2, clave3, clave4, clave5}
	dic := TDADiccionario.CrearABB[int, *int](func_cmp_int)
	dic.Guardar(claves[0], nil)
	dic.Guardar(claves[1], nil)
	dic.Guardar(claves[2], nil)
	dic.Guardar(claves[3], nil)
	dic.Guardar(claves[4], nil)

	cs := []int{0, 0, 0, 0, 0}
	cantidad := 0
	cantPtr := &cantidad
	desde := 2
	hasta := 4

	dic.IterarRango(&desde, &hasta, func(clave int, dato *int) bool {
		cs[cantidad] = clave
		*cantPtr = *cantPtr + 1
		return true
	})

	require.EqualValues(t, 3, cantidad)
	require.NotEqualValues(t, -1, buscarInt(cs[0], claves))
	require.NotEqualValues(t, -1, buscarInt(cs[1], claves))
	require.NotEqualValues(t, -1, buscarInt(cs[2], claves))
	require.NotEqualValues(t, cs[0], cs[1])
	require.NotEqualValues(t, cs[0], cs[2])
	require.NotEqualValues(t, cs[2], cs[1])
}

