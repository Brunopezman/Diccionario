package diccionario

import (
	TDAPila "tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izq   *nodoAbb[K, V]
	der   *nodoAbb[K, V]
	clave K
	dato  V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	cmp      func(K, K) int
}

type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearIterador[K comparable, V any](ab *abb[K, V], desde *K, hasta *K) *iterAbb[K, V] {
	return &iterAbb[K, V]{abb: ab, pila: TDAPila.CrearPilaDinamica[*nodoAbb[K, V]](), desde: desde, hasta: hasta}
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func CrearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izq: nil, der: nil, clave: clave, dato: dato}
}

func (ab *abb[K, V]) buscarPuntero(clave K, nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	//dir nodo vacia: No tiene ningun hijo
	if *nodo == nil {
		return nodo
	}
	// Elemento encontrado
	if ab.cmp(clave, (*nodo).clave) == 0 {
		return nodo
	}
	// Existe raiz, entonces comparo a izq, despues der
	if ab.cmp(clave, (*nodo).clave) < 0 {
		return ab.buscarPuntero(clave, &(*nodo).izq)
	} else {
		return ab.buscarPuntero(clave, &(*nodo).der)
	}
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	puntero := ab.buscarPuntero(clave, &ab.raiz)
	if *puntero == nil {
		//Si el arbol no tiene raiz
		*puntero = CrearNodoABB(clave, dato)
		ab.cantidad++
	} else {
		(*puntero).dato = dato
	}
}

func (ab *abb[K, V]) Pertenece(clave K) bool {
	return *(ab.buscarPuntero(clave, &ab.raiz)) != nil
}

func (ab *abb[K, V]) Obtener(clave K) V {
	puntero := ab.buscarPuntero(clave, &ab.raiz)
	if *puntero != nil {
		return (*puntero).dato
	}
	panic("La clave no pertenece al diccionario")
}

func (ab *abb[K, V]) Borrar(clave K) V {
	puntero := ab.buscarPuntero(clave, &ab.raiz)
	if *puntero == nil {
		panic("La clave no pertenece al diccionario")
	}
	eliminado := (*puntero).dato
	hijos := ab.cantidadHijos(puntero)
	// Caso 0 hijos, Caso 1 hijo, Caso 2 hijos
	if hijos == 0 {
		*puntero = nil
	} else if hijos == 1 {
		nodo := ab.obtenerHijo(puntero)
		*puntero = *nodo
	} else {
		nodo := ab.buscarReemplazo(&(*puntero).izq)
		clave, dato := (*nodo).clave, (*nodo).dato
		(*puntero).clave = clave
		(*puntero).dato = dato
	}

	ab.cantidad--
	return eliminado
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

func (ab *abb[K, V]) cantidadHijos(nodo **nodoAbb[K, V]) int {
	if (*nodo).izq == nil && (*nodo).der == nil {
		return 0
	} else if (*nodo).izq != nil && (*nodo).der == nil || (*nodo).izq == nil && (*nodo).der != nil {
		return 1
	} else {
		return 2
	}
}

func (ab *abb[K, V]) obtenerHijo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izq == nil {
		return &(*nodo).der
	}
	return &(*nodo).izq
}

func (ab *abb[K, V]) buscarReemplazo(nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).der == nil {
		return nodo
	}
	return ab.buscarReemplazo(&(*nodo).der)
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	iter := CrearIterador(ab, nil, nil)
	actual := ab.raiz
	for actual != nil {
		iter.pila.Apilar(actual)
		actual = actual.izq
	}
	return iter
}
func (a *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	_Iterar(visitar, &a.raiz)
}

func _Iterar[K comparable, V any](visitar func(clave K, dato V) bool, nodo **nodoAbb[K, V]) bool {
	if *nodo == nil {
		return true
	}
	if !_Iterar(visitar, &(*nodo).izq) || !visitar((*nodo).clave, (*nodo).dato) || !_Iterar(visitar, &(*nodo).der) {
		return false
	}
	return true
}

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := CrearIterador(ab, desde, hasta)
	actual := ab.raiz
	for actual != nil {
		if ab.cmp(actual.clave, *desde) == 1 {
			iter.pila.Apilar(actual)
			actual = actual.izq
		}
	}
	return iter
}

func (a *abb[K, V]) IterarRango(desde, hasta *K, visitar func(clave K, dato V) bool) {
	_IterarRango(&a.raiz, desde, hasta, visitar, a.cmp)
}

func _IterarRango[K comparable, V any](nodo **nodoAbb[K, V], desde, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) {
	if *nodo == nil {
		return
	}
	if desde == nil || cmp(*desde, (*nodo).clave) < 0 {
		_IterarRango(&(*nodo).izq, desde, hasta, visitar, cmp)
	}
	if desde == nil || cmp(*desde, (*nodo).clave) <= 0 && (hasta == nil || cmp(*hasta, (*nodo).clave) >= 0) {
		visitar((*nodo).clave, (*nodo).dato)
	}
	if hasta == nil || cmp(*hasta, (*nodo).clave) > 0 {
		_IterarRango(&(*nodo).der, desde, hasta, visitar, cmp)
	}

}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	nodo := iter.pila.Desapilar()
	if nodo.der != nil {
		if iter.desde == nil || iter.abb.cmp(nodo.der.clave, *iter.hasta) == -1 {
			iter.pila.Apilar(nodo.der)
		}

		actual := nodo.der.izq
		for actual != nil {
			if iter.desde == nil || iter.abb.cmp(actual.clave, *iter.desde) == 1 {
				iter.pila.Apilar(actual)
			}
		}
	}
}
