from library import libreria_tp as lb


def ciclo_euleriano(grafo, v):
    """PRE: Recibe un grafo y un vertice
    POST: Devuelve el ciclo euleriano y el tiempo total, None caso contrario"""
    if lb.es_conexo(grafo):
        camino, peso = hierholzer(grafo, v)
        if camino != None:
            return unica_lista(camino, []), peso
    return None, None


def hierholzer(grafo, v):
    """Devuelve el camino euleriano y el peso total"""
    lista = []
    aristas = set()
    peso = 0

    lista.append(v)
    camino, peso = _hierholzer(grafo, v, v, aristas, peso, lista)
    return camino, peso


def _hierholzer(grafo, v, origen, aristas, peso, lista):
    """Devuelve el camino euleriano y el peso total"""
    for w in grafo.adyacentes(v):
        if len(grafo.adyacentes(w)) % 2 != 0:
            return None, None
        if (v, w) not in aristas:
            aristas.add((v, w))
            aristas.add((w, v))
            peso += grafo.peso(v, w)
            lista.append(w)
            if w == origen:
                for c in lista:
                    for x in grafo.adyacentes(c):
                        if (c, x) not in aristas:
                            nueva = []
                            nueva.append(c)
                            camino2, peso = _hierholzer(grafo, c, origen, aristas, peso, nueva)
                            lista[lista.index(c)] = camino2
                return lista, peso
            else:
                return _hierholzer(grafo, w, origen, aristas, peso, lista)
    return lista, peso


def unica_lista(dato, lista_plana):
    """PRE: Recibe una lista de listas
    POST: Devuelve una lista con los elementos de la lista de listas"""
    for elemento in dato:
        if type(elemento) == list:
            unica_lista(elemento, lista_plana)
        else:
            lista_plana.append(elemento)
    return lista_plana
