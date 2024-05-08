package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_TAM_INICIAL     int = 13
	_FACTOR_AGRANDAR int = 3
	_FACTOR_ACHICAR  int = 2
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

func (h *hash[K, V]) inicializarTabla() {
	h.tabla = make([]TDALista.Lista[parClaveValor[K, V]], h.tam)
	h.cantidad = 0
}

func CrearHash[K comparable, V any]() Diccionario[K, V] {
	diccionario := &hash[K, V]{
		tam:      _TAM_INICIAL,
		cantidad: 0,
	}
	diccionario.inicializarTabla()
	return diccionario
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

func (h *hash[K, V]) Guardar(clave K, dato V) {
	pos := fhash(clave, h.tam)

	if h.cantidad/h.tam >= _FACTOR_AGRANDAR {
		h.redimensionar(h.tam * _FACTOR_AGRANDAR)
	}

	if h.tabla[pos] == nil {
		h.tabla[pos] = TDALista.CrearListaEnlazada[parClaveValor[K, V]]()
	}
	for iter := h.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
		parClaveValor := iter.VerActual()
		if parClaveValor.clave == clave {
			iter.Borrar()
			iter.Insertar(*CrearParClaveValor(clave, dato))
			return
		}
	}
	h.tabla[pos].InsertarUltimo(*CrearParClaveValor(clave, dato))
	h.cantidad++

}

func (h *hash[K, V]) Pertenece(clave K) bool {
	pos := fhash(clave, h.tam)
	if h.tabla[pos] != nil {
		for iter := h.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				return true
			}
		}
	}
	return false

}

func (h *hash[K, V]) Obtener(clave K) V {
	pos := fhash(clave, h.tam)
	if h.tabla[pos] != nil {
		for iter := h.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				return parClaveValor.valor
			}
		}
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hash[K, V]) Borrar(clave K) V {
	pos := fhash(clave, h.tam)

	if h.tabla[pos] != nil {
		for iterLista := h.tabla[pos].Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
			parClaveValor := iterLista.VerActual()
			if parClaveValor.clave == clave {
				iterLista.Borrar()
				h.cantidad--
				if h.tabla[pos].EstaVacia() {
					h.tabla[pos] = nil
				}
				return parClaveValor.valor
			}
		}
	}
	if h.cantidad/h.tam < _FACTOR_ACHICAR {
		h.redimensionar(h.tam / _FACTOR_ACHICAR)
	}
	panic("La clave no pertenece al diccionario")
}

func (h *hash[K, V]) Cantidad() int {
	return h.cantidad
}

func (h *hash[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	for pos := 0; pos < h.tam; pos++ {
		if h.tabla[pos] != nil {
			for iterLista := h.tabla[pos].Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
				if !visitar(iterLista.VerActual().clave, iterLista.VerActual().valor) {
					return
				}
			}
		}
	}

}

func (h *hash[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{pos: 0, diccionario: h}
	for iterador.pos < h.tam {
		if h.tabla[iterador.pos] == nil {
			iterador.pos += 1
		} else {
			iterador.parActual = h.tabla[iterador.pos].VerPrimero()
			break
		}
	}
	return iterador
}

func (iter *iterDiccionario[K, V]) HaySiguiente() bool {
	return iter.pos < iter.diccionario.tam
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

	for iterLista := iter.diccionario.tabla[iter.pos].Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
		parClaveValor := iterLista.VerActual()
		if iter.parActual.clave == parClaveValor.clave {
			iterLista.Siguiente()
			if iterLista.HaySiguiente() {
				iter.parActual = iterLista.VerActual()
				return
			}
			break
		}
	}
	iter.pos++
	for iter.pos < iter.diccionario.tam {
		if iter.diccionario.tabla[iter.pos] != nil {
			iter.parActual = iter.diccionario.tabla[iter.pos].VerPrimero()
			return
		}
		iter.pos++
	}

}

func (h *hash[K, V]) redimensionar(nuevoTam int) {

	tablaAnterior := h.tabla
	tamAnterior := h.tam

	h.tam = nuevoTam
	h.inicializarTabla()
	h.cantidad = 0

	for pos := 0; pos < tamAnterior; pos++ {
		if tablaAnterior[pos] != nil {
			lista := tablaAnterior[pos]
			for iterLista := lista.Iterador(); iterLista.HaySiguiente(); iterLista.Siguiente() {
				h.Guardar(iterLista.VerActual().clave, iterLista.VerActual().valor)
			}

		}

	}
}
