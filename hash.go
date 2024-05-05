package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_TAM_INICIAL int = 10
)

type parClaveValor[K comparable, V any] struct {
	clave K
	valor V
}

type hash[K comparable, V any] struct {
	tabla    []TDALista.Lista[parClaveValor[K, V]]
	tam      int
	cantidad int
}

type iterDiccionario[K comparable, V any] struct {
	parActual   parClaveValor[K, V]
	pos         int
	diccionario *hash[K, V]
}

func CrearParClaveValor[K comparable, V any](clave K, valor V) *parClaveValor[K, V] {
	return &parClaveValor[K, V]{clave: clave, valor: valor}
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	return &hash[K, V]{
		tabla:    make([]TDALista.Lista[parClaveValor[K, V]], _TAM_INICIAL),
		tam:      _TAM_INICIAL,
		cantidad: 0,
	}
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
	pos := fhash(clave, _TAM_INICIAL)
	if dic.tabla[pos] == nil {
		dic.tabla[pos] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	dic.tabla[pos].InsertarUltimo(*CrearParClaveValor(clave, dato))
	dic.cantidad++

}

func (dic *hash[K, V]) Pertenece(clave K) bool {
	pos := fhash(clave, _TAM_INICIAL)
	if dic.tabla[pos] != nil {
		for iter := dic.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				return true
			}
		}
	}
	return false

}

func (dic *hash[K, V]) Obtener(clave K) V {
	pos := fhash(clave, _TAM_INICIAL)
	if dic.tabla[pos] != nil {
		for iter := dic.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				return parClaveValor.valor
			}
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (dic *hash[K, V]) Borrar(clave K) V {
	pos := fhash(clave, _TAM_INICIAL)
	if dic.tabla[pos] != nil {
		for iter := dic.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				iter.Borrar()
				dic.cantidad--
				return parClaveValor.valor
			}
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
	pos := 0

	for dic.tabla[pos] == nil {
		pos += 1

	}
	actual := dic.tabla[pos].VerPrimero()
	return &iterDiccionario[K, V]{parActual: actual,
		pos:         pos,
		diccionario: dic}
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {

	return iter.pos != iter.diccionario.cantidad
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.parActual.clave, iter.parActual.valor
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

}
