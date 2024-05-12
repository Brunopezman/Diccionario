package diccionario

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

// type iterAbb[K comparable, V any] struct {
// 	abb   *abb[K, V]
// 	pila  TDAPila.Pila[*nodoAbb[K, V]]
// 	desde *K
// 	hasta *K
// }

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{cmp: funcion_cmp}
}

func CrearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izq: nil, der: nil, clave: clave, dato: dato}
}

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
	}
	return nodoPadre
}

func (ab *abb[K, V]) Guardar(clave K, dato V) {
	puntero := ab.buscarPuntero(clave, &ab.raiz)
	if *puntero == nil {
		//Si el arbol no tiene raiz
		*puntero = CrearNodoABB(clave, dato)
		ab.cantidad++
	}
	// si ya existe, actualizo el dato del nodo
	(*puntero).dato = dato

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
	// Caso 0 hijos, Caso 1 hijo, Caso 2 hijos
	if ab.cantidadHijos(puntero) == 0 {
		*puntero = nil
	} else if ab.cantidadHijos(puntero) == 1 {
		nodo := ab.obtenerHijo(puntero)
		*puntero = *nodo
	} else {
		nodo := ab.buscarReemplazo(puntero)
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

func (ab *abb[K, V]) Iterar(funcion_cmp func(clave K, dato V) bool) {}
func (ab *abb[K, V]) Iterador() IterDiccionario[K, V] {
	return ab.IteradorRango(nil, nil)
}

// func (a *iterAbb[K, V]) VerActual()    {}
// func (i *iterAbb[K, V]) HaySiguiente() {}
// func (i *iterAbb[K, V]) Siguiente()    {}

// func (ab *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
// 	_IterarRango(a.raiz, desde, hasta, visitar, a)
// }

// func _IterarRango[K comparable, V any](n *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, ab *abb[K, V]) {
// }

func (ab *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	return ab.Iterador()
}

