package diccionario_test

import (
	//"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func func_cmp(e1, e2 int) int {
	if e1 == e2 {
		return 0
	}
	if e1 > e2 {
		return 1
	}
	return -1
}
func TestDiccionarioAbbVacio(t *testing.T) {
	t.Log("Comprueba que Diccionario vacio no tiene claves")
	dic := TDADiccionario.CrearABB[int, int](func_cmp)
	require.EqualValues(t, 0, dic.Cantidad())
	require.False(t, dic.Pertenece(3))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener(3) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar(3) })
}
func TestDiccionarioAbbClaveDefault(t *testing.T) {
	t.Log("Prueba sobre un Hash vacío que si justo buscamos la clave que es el default del tipo de dato, " +
		"sigue sin existir")
	dic := TDADiccionario.CrearHash[string, string]()
	require.False(t, dic.Pertenece(""))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Obtener("") })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dic.Borrar("") })

	dicNum := TDADiccionario.CrearHash[int, string]()
	require.False(t, dicNum.Pertenece(0))
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Obtener(0) })
	require.PanicsWithValue(t, "La clave no pertenece al diccionario", func() { dicNum.Borrar(0) })
}

func TestUnElementoAbb(t *testing.T) {
	t.Log("Comprueba que Diccionario con un elemento tiene esa Clave, unicamente")
	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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

	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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
	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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

	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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

	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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

func TestGuardarYBorrarRepetidasVecesAbb(t *testing.T) {
	t.Log("Esta prueba guarda y borra repetidas veces. Esto lo hacemos porque un error comun es no considerar " +
		"los borrados para agrandar en un Hash Cerrado. Si no se agranda, muy probablemente se quede en un ciclo " +
		"infinito")

	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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
	dic := TDADiccionario.CrearABB[int, *int](func_cmp)
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

	dic := TDADiccionario.CrearABB[int, int](func_cmp)
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

