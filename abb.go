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
type iterAbb[K comparable, V any] struct {
	abb   *abb[K, V]
	pila  TDAPila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

func CrearABB[K comparable, V any](funcion_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: funcion_cmp}
}

func CrearNodoABB[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izq: nil, der: nil, clave: clave, dato: dato}
}

func (a *abb[K, V]) buscarPuntero(clave K, nodo **nodoAbb[K, V]) **nodoAbb[K, V] {
	//dir nodo vacia: No tiene ningun hijo
	if *nodo == nil {
		return nodo
	}
	//Comparo con los hijos

}

func (a *abb[K, V]) Guardar(clave K, dato V) {
	puntero := a.buscarPuntero(clave, &a.raiz)
	if *puntero == nil {
		*puntero = CrearNodoABB(clave, dato)
		a.cantidad++
	}
	(*puntero).dato = dato

}

// func (a *abb[K, V]) Pertenece(clave K) bool

// func (a *abb[K, V]) Obtener(clave K) V

// func (a *abb[K, V]) Borrar(clave K) V

func (a *abb[K, V]) Cantidad() int {
	return a.cantidad
}

func (a *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	_IterarRango(a.raiz, desde, hasta, visitar, a)
}

func _IterarRango[K comparable, V any](n *nodoAbb[K, V], desde *K, hasta *K, visitar func(clave K, dato V) bool, a *abb[K, V])

// func (a *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V]
