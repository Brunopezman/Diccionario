package diccionario

import TDAPila "tdas/pila"

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
	pila   TDAPila.Pila[*nodoAbb[K, V]]
	actual *nodoAbb[K, V]
}

// Funciones de creacion
func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func crearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izq: nil, der: nil, clave: clave, dato: dato}
}

// Primitivas diccionario
func (ab *abb[K, V]) Guardar(clave K, dato V) {
	puntero := ab.buscarPuntero(clave, &ab.raiz)
	if *puntero == nil {
		//Si el arbol no tiene raiz
		*puntero = crearNodoABB(clave, dato)
		ab.cantidad++
	} else {
		// si ya existe, actualizo el dato del nodo
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
	hijos := ab.cantidadHijos(puntero)
	eliminado := (*puntero).dato
	// Caso 0 hijos, Caso 1 hijo, Caso 2 hijos
	if hijos == 0 {
		*puntero = nil
	} else if hijos == 1 {
		nodo := ab.obtenerHijo(puntero)
		*puntero = *nodo
	} else {
		nodo := ab.buscarReemplazo(&(*puntero).izq)
		clave, dato := (*nodo).clave, (*nodo).dato
		*nodo = (*nodo).izq
		(*puntero).clave = clave
		(*puntero).dato = dato
	}

	ab.cantidad--
	return eliminado
}

func (ab *abb[K, V]) Cantidad() int {
	return ab.cantidad
}

// Funciones soporte para las primitivas

func (ab *abb[K, V]) buscarPuntero(clave K, nodoPadre **nodoAbb[K, V]) **nodoAbb[K, V] {
	//dir nodo vacia: No tiene ningun hijo
	if *nodoPadre == nil {
		return nodoPadre
	}
	// Existe raiz, entonces comparo a izq, despues der
	if ab.cmp(clave, (*nodoPadre).clave) < 0 {
		if (*nodoPadre).izq == nil || ab.cmp(clave, (*nodoPadre).izq.clave) == 0 {
			//  Quiere decir que en realidad es hoja ahi, por lo tanto la retorno
			return &(*nodoPadre).izq
		}
		return ab.buscarPuntero(clave, &(*nodoPadre).izq)
	} else if ab.cmp(clave, (*nodoPadre).clave) > 0 {
		if (*nodoPadre).der == nil || ab.cmp(clave, (*nodoPadre).der.clave) == 0 {
			return &(*nodoPadre).der
		}
		return ab.buscarPuntero(clave, &(*nodoPadre).der)
	} else {
		return nodoPadre
	}
}

func (ab *abb[K, V]) cantidadHijos(nodo **nodoAbb[K, V]) int {
	if (*nodo).izq != nil && (*nodo).der != nil {
		return 2
	} else if (*nodo).izq == nil && (*nodo).der == nil {
		return 0
	} else {
		return 1
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

// Iteradores internos
func (ab *abb[K, V]) Iterar(visitar func(clave K, dato V) bool) {
	ab.IterarRango(nil, nil, visitar)
}

func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if ab.raiz == nil {
		return
	}
	if desde == nil {
		primerNodo := primero(&ab.raiz)
		desde = &(*primerNodo).clave
	}
	if hasta == nil {
		ultimoNodo := ultimo(&ab.raiz)
		hasta = &(*ultimoNodo).clave
	}
	_IterarRango(&ab.raiz, desde, hasta, visitar, ab.cmp)
}

func _IterarRango[K comparable, V any](nodo **nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, cmp func(K, K) int) bool {
	if *nodo == nil {
		return true
	}

	if cmp((*nodo).clave, *desde) < 0 {
		return _IterarRango(&(*nodo).der, desde, hasta, visitar, cmp)
	} else if cmp((*nodo).clave, *hasta) > 0 {
		return _IterarRango(&(*nodo).izq, desde, hasta, visitar, cmp)
	} else if cmp((*nodo).clave, *hasta) == 0 && cmp((*nodo).clave, *desde) == 0 {
		visitar((*nodo).clave, (*nodo).dato)
		return false
	}

	if !_IterarRango(&(*nodo).izq, desde, hasta, visitar, cmp) {
		return false
	}

	if !visitar((*nodo).clave, (*nodo).dato) {
		return false
	}

	return _IterarRango(&(*nodo).der, desde, hasta, visitar, cmp)
}

// Iterador Externo
func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	iter := new(iterAbb[K, V])
	iter.pila = TDAPila.CrearPilaDinamica[*nodoAbb[K, V]]()
	iter.pila.Apilar(nil)

	if ab.raiz != nil {
		if desde == nil {
			primerNodo := primero(&ab.raiz)
			desde = &(*primerNodo).clave
		}
		if hasta == nil {
			ultimoNodo := ultimo(&ab.raiz)
			hasta = &(*ultimoNodo).clave
		}

		apilarNodosPorRango(&ab.raiz, desde, hasta, iter, ab.cmp)
	}

	iter.actual = iter.pila.Desapilar()
	return iter
}

func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

// Primitivas de diccionario
func (iter *iterAbb[K, V]) HaySiguiente() bool {
	return iter.actual != nil
}

func (iter *iterAbb[K, V]) VerActual() (K, V) {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iter.actual.clave, iter.actual.dato
}

func (iter *iterAbb[K, V]) Siguiente() {
	if !iter.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iter.actual = iter.pila.Desapilar()
}

// Otras funciones
func apilarNodosPorRango[K comparable, V any](nodo **nodoAbb[K, V], desde *K, hasta *K, iter *iterAbb[K, V], cmp func(K, K) int) {
	if *nodo == nil {
		return
	}
	if iter.pila.VerTope() != nil {
		if cmp(iter.pila.VerTope().clave, *desde) <= 0 {
			return
		}
	}
	apilarNodosPorRango(&(*nodo).der, desde, hasta, iter, cmp)
	if cmp((*nodo).clave, *desde) >= 0 && cmp((*nodo).clave, *hasta) <= 0 {
		iter.pila.Apilar(*nodo)
	}
	apilarNodosPorRango(&(*nodo).izq, desde, hasta, iter, cmp)

}

func ultimo[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).der == nil {
		return nodo
	}
	return ultimo(&(*nodo).der)
}

func primero[K comparable, V any](nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	if (*nodo).izq == nil {
		return nodo
	}
	return primero(&(*nodo).izq)
}

