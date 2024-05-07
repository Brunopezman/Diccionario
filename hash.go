package diccionario

import (
	"fmt"
	TDALista "tdas/lista"
)

const (
	_TAM_INICIAL int = 7
	_FACTOR_AGRANDAR int = 2
	_FACTOR_ACHICAR int = 2
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

func (h *hash[K, V]) Guardar(clave K, dato V) {
	pos := fhash(clave, _TAM_INICIAL)

	if h.cantidad == 3* (h.tam - 1) {
		redimensionar(_FACTOR_AGRANDAR)
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
	pos := fhash(clave, _TAM_INICIAL)
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
	pos := fhash(clave, _TAM_INICIAL)
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
	pos := fhash(clave, _TAM_INICIAL)
	
	if h.cantidad ==  (h.tam - 1) / 4 {
		redimensionar(_FACTOR_ACHICAR)
	}
	
	if h.tabla[pos] != nil {
		for iter := h.tabla[pos].Iterador(); iter.HaySiguiente(); iter.Siguiente() {
			parClaveValor := iter.VerActual()
			if parClaveValor.clave == clave {
				iter.Borrar()
				h.cantidad--
				return parClaveValor.valor
			}
		}
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
					break
				}
			}
		}
	}

}

func (h *hash[K, V]) Iterador() IterDiccionario[K, V] {
	iterador := &iterDiccionario[K, V]{pos:0, diccionario:h}
	for iterador.pos < h.tam {
		if h.tabla[pos] == nil {
			iterador.pos += 1
		} else {
			iterador.parActual = h.tabla[iterador.pos].VerPrimero()
			}
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
	iter.pos++
	for iter.pos < iter.diccionario.tam {
		if iter.diccionario.tabla[iter.pos] != nil {
			iter.parActual = iter.diccionario.tabla[iter.pos].VerPrimero()
			return
		}
		iter.pos++
	}

}

func (h *hash[K, V]) redimensionar(factor int) {

	tablaActual := h.tabla
	tamActual := h.tam

	h.tabla:make([]TDALista.Lista[parClaveValor[K, V]], h.tam*factor)),
	h.tam = h.tam*factor
	h.cantidad: 0

	for pos:=0; pos < tamActual; pos++ {
		lista := tablaActual[pos]
		h.Guardar(lista.parClaveValor)
	}
}
