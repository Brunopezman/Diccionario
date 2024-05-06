package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_TAM_INICIAL int = 7
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
	for iter := dic.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		parClaveValor := iter.VerActual()
		if parClaveValor.clave == clave {
			iter.Borrar()
			iter.Insertar(*CrearParClaveValor(clave, dato))
			return
		}
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

func (dic *hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for pos := 0; pos < dic.tam; pos++ {
		if dic.tabla[pos] != nil {
			for iterLista := dic.tabla[pos].Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
				if !visitar(iterLista.VerActual().clave, iterLista.VerActual().valor) {
					break
				}
			}
		}
	}

}

func (dic *hash[K, V]) Iterador() IterDiccionario[K, V] {

	pos := 0

	for pos < dic.tam {
		if dic.tabla[pos] == nil {
			pos += 1
		} else {
			return &iterDiccionario[K, V]{
				parActual:   dic.tabla[pos].VerPrimero(),
				pos:         pos,
				diccionario: dic,
			}
		}
	}
	return nil

}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter != nil

	//return iter.pos < iter.diccionario.tam || iter.diccionario.tabla[iter.pos].Iterador().HaySiguiente()
}

func (iter *iterDiccionario[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}

	return iter.parActual.clave, iter.parActual.valor
}

func (iter *iterDiccionario[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador terminó de iterar")
	}

	for iterLista := iter.diccionario.tabla[iter.pos].Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
		parClaveValor := iterLista.VerActual()
		if iter.parActual.clave == parClaveValor.clave {
			iterLista.Siguiente()
			if iterLista.HaySiguiente() {
				iter.parActual = iterLista.VerActual()
				return
			}
		}
	}

	for iter.pos < iter.diccionario.tam {
		if iter.diccionario.tabla[iter.pos] != nil {
			iter.parActual = iter.diccionario.tabla[iter.pos].VerPrimero()
			return
		}
		iter.pos++
	}

	if iter.pos > iter.diccionario.tam {
		iter = nil
	}

}

func (dic *hash[K, V]) redimensionar(n, m int) {
	factor = n / m

}
