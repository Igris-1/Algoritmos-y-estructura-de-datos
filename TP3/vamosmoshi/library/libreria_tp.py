import heapq
from tdas import pila
from tdas import cola


def camino_minimo(grafo, desde, hasta):
    """Pre: Recibe un grafo y dos vertices
    Post: Devuelve el camino minimo entre los vertices y el tiempo total
    """
    padres, _ = dijkstra(grafo, desde, hasta)
    if padres is None:
        return None, None
    camino, tiempo = reconstruir_camino(padres, hasta, grafo)
    return camino, tiempo


def reconstruir_camino(padres, hasta, grafo):
    """Reconstruye el camino mínimo entre dos ciudades.
    Devuelve una lista con las ciudades que forman el camino
    y el 'tiempo' total del mismo."""
    camino = []
    tiempo = 0
    actual = hasta
    while actual is not None:
        camino.append(actual)
        if padres[actual] is not None:
            tiempo += grafo.peso(padres[actual], actual)
        actual = padres[actual]
    camino.reverse()
    return camino, tiempo


def dijkstra(grafo, origen, destino):
    """Pre: Recibe un grafo y dos vertices
    Post: Devuelve el camino minimo entre los vertices y el peso"""
    dist = {}
    padre = {}

    for v in grafo.obtener_vertices():
        dist[v] = float("inf")
        padre[v] = None

    dist[origen] = 0
    padre[origen] = None
    heap = []
    heapq.heappush(heap, (0, origen))

    while heap:
        _, v = heapq.heappop(heap)
        if v == destino:
            return padre, dist

        for w in grafo.adyacentes(v):
            distancia_por_v = dist[v] + grafo.peso(v, w)
            if distancia_por_v < dist[w]:
                dist[w] = distancia_por_v
                padre[w] = v
                heapq.heappush(heap, (dist[w], w))
    return None, None


def obtener_ciclo_dfs(grafo):
    """Pre: Recibe un grafo
    Post: Devuelve True o false si el grafo tiene un ciclo o no"""
    visitados = {}
    padre = {}
    for v in grafo.obtener_vertices():
        if v not in visitados:
            ciclo = dfs_ciclo(grafo, v, visitados, padre)
            if not ciclo:
                return False
    return True


def dfs_ciclo(grafo, v, visitados, padre):
    """Pre: Recibe un grafo
    Post: devuelve True o False si el grafo tiene un ciclo o no"""
    visitados[v] = True
    for w in grafo.adyacentes(v):
        if w in visitados:
            if w != padre[v]:
                return True
        else:
            padre[w] = v
            ciclo = dfs_ciclo(grafo, w, visitados, padre)
            if ciclo is not None:
                return ciclo
    return False


def orden_topologico_dfs(grafo):
    """Devuelve un orden topologico del grafo"""
    visitados = set()
    _pila = pila.Pila()
    for v in grafo.obtener_vertices():
        if v not in visitados:
            visitados.add(v)
            dfs(grafo, v, visitados, _pila)
    return pila_a_lista(_pila), visitados


def dfs(grafo, v, visitados, pila):
    """Funcion auxiliar de orden_topologico_dfs"""
    for w in grafo.adyacentes(v):
        if w not in visitados:
            visitados.add(w)
            dfs(grafo, w, visitados, pila)
    pila.apilar(v)


def pila_a_lista(pila):
    """Funcion auxiliar de orden_topologico_dfs
    Devuelve una lista con los elementos de la pila"""
    lista = []
    while not pila.esta_vacia():
        lista.append(pila.desapilar())
    return lista


def tm_prim(grafo):
    """Pre: Recibe un grafo
    Post: Devuelve el arbol de tendido minimo y el peso total"""
    v = grafo.obtener_vertices()[0]
    visitados = set()
    visitados.add(v)
    heap = []

    for w in grafo.adyacentes(v):
        heapq.heappush(heap, (grafo.peso(v, w), v, w))

    arbol = []
    peso_total = 0
    while heap:
        peso, v, w = heapq.heappop(heap)
        if w in visitados:
            continue
        visitados.add(w)
        arbol.append((v, w, peso))
        peso_total += peso
        for x in grafo.adyacentes(w):
            if x not in visitados:
                heapq.heappush(heap, (grafo.peso(w, x), w, x))
    return arbol, peso_total


def bfs_gen(grafo, origen, vis):
    q = cola.Cola()
    q.encolar(origen)
    vis.add(origen)
    while not q.esta_vacia():
        v = q.desencolar()
        for w in grafo.adyacentes(v):
            if w not in vis:
                vis.add(w)
                q.encolar(w)


def es_conexo(grafo):
    visitados = set()
    cant = 0
    for v in grafo.obtener_vertices():
        if v not in visitados:
            cant += 1
            if cant == 2:
                return False
            bfs_gen(grafo, v, visitados)
    return True
