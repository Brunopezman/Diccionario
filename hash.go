package diccionario

import "fmt"

type hash[K comparable, V any] struct {
	clave    K
	valor    V
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	actual      K
	siguiente   K
	diccionario *hash[K, V]
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{clave: nil, valor: nil, cantidad: 0}
}

func (dic *hash[K, V]) Guardar(clave K, dato V) {}

func (dic *hash[K, V]) Pertenece(clave K) bool {
	return false
}

func (dic *hash[K, V]) Obtener(clave K) V {}

func (dic *hash[K, V]) Borrar(clave K) V {}

func (dic *hash[K, V]) Cantidad() int {
	return dic.cantidad
}

func (dic *hash[K, V]) Iterar(func(clave K, dato V) bool) {}

func (dic *hash[K, V]) Iterador() IterDiccionario[K, V] {}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {}

func (iter *iterDiccionario[K, V]) Siguiente() {}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprint("%v", clave))
}
