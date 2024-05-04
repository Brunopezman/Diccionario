package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_CAPACIDAD_INICIAL int = 10
)

type hash[K comparable, V any] struct {
	datos    []TDALista.Lista[K]
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
	return &hash[K, V]{datos: make([]TDALista.Lista[K], _CAPACIDAD_INICIAL), cantidad: 0}
}

func convertirABytes[K comparable](clave K) []byte {
	return []byte(fmt.Sprintf("%v", clave))
}

func fhash[K comparable](clave K, capacidad int) int {
	hash := 0
	for _, c := range convertirABytes(clave) {
		hash = (31*hash + int(c)) % capacidad
	}
	return hash
}

func (dic *hash[K, V]) Guardar(clave K, dato V) {
	pos := fhash(clave, _CAPACIDAD_INICIAL)
	if dic.datos[pos].EstaVacia() {
		dic.datos[pos].InsertarPrimero(clave)
	}
	for iter := dic.datos[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		iter.Insertar(clave)
	}

	dic.cantidad++

}

func (dic *hash[K, V]) Pertenece(clave K) bool {
	pos := fhash(clave, _CAPACIDAD_INICIAL)
	for iter := dic.datos[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		elemento := iter.VerActual()
		if elemento == clave {
			return true
		}
	}
	return false

}

func (dic *hash[K, V]) Obtener(clave K) V {
	pos := fhash(clave, _CAPACIDAD_INICIAL)
	for iter := dic.datos[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		elemento := iter.VerActual()
		if elemento == clave {
			return dic.valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dic *hash[K, V]) Borrar(clave K) V {
	pos := fhash(clave, _CAPACIDAD_INICIAL)
	for iter := dic.datos[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		elemento, valor := dic.Iterador().VerActual()
		if elemento == clave {
			iter.Borrar()
			dic.cantidad--
			return valor
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dic *hash[K, V]) Cantidad() int {
	return dic.cantidad
}

func (dic *hash[K, V]) Iterar(func(clave K, dato V) bool) {

}

func (dic *hash[K, V]) Iterador() IterDiccionario[K, V] {
	return &iterDiccionario[K, V]{}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.diccionario.datos != nil
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.diccionario.clave, iter.diccionario.valor
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.actual = iter.siguiente
}
