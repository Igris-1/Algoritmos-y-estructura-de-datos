from library import auxiliares, libreria_tp, Euleriano


def lectura_comandos(entrada):
    """
    DOC: Recibe una entrada de tipos str y devuelve dos listas
    """
    entrada = entrada.rstrip().split(" ")
    comando = entrada[0]
    parametros = entrada[1:]

    parametros = " ".join(parametros).split(",")
    for i in range(len(parametros)):
        parametros[i] = parametros[i].lstrip()

    if comando == "ir":
        return comando, parametros
    if comando == "itinerario":
        return comando, parametros
    if comando == "viaje":
        return comando, parametros
    if comando == "reducir_caminos":
        return comando, parametros


## comando 1
def ir(desde, hasta, ruta, grafo, latitudes):
    """
    PRE: Recibe dos vertices y un grafo
    POST: Imprime el camino minimo entre los vertices y el tiempo total
    """
    if auxiliares.pertenece([desde, hasta], grafo):
        camino_minimo, tiempo = libreria_tp.camino_minimo(grafo, desde, hasta)
        if camino_minimo is None:
            print("No se encontro recorrido")
        else:
            camino = " -> ".join(camino_minimo)
            print(camino)
            print(f"Tiempo total: {tiempo}")
            auxiliares.guardar_kml(ruta, camino_minimo, latitudes)
    else:
        print("No se encontro recorrido")


## comando 2
def itinerario(ruta, grafo):
    """
    Pre: Recibe un grafo
    Post: Imprime el itinerario de las ciudades
    """
    grafo_reco = auxiliares.recomendaciones(ruta)
    if not libreria_tp.obtener_ciclo_dfs(grafo_reco):
        orden_reco, visitados = libreria_tp.orden_topologico_dfs(grafo_reco)
        for v in grafo.obtener_vertices():
            if v not in visitados:
                orden_reco.append(v)
        print(" -> ".join(orden_reco))
    else:
        print("No se encontro recorrido")


## comando 3
def viaje(origen, archivo, grafo, latitudes):
    """
    PRE: recibe un vertice, y un grafo
    POST: devuelve el ciclo y el tiempo empleado
    """
    if auxiliares.pertenece([origen], grafo):
        recorrido, tiempo = Euleriano.ciclo_euleriano(grafo, origen)
        if recorrido:
            print(" -> ".join(recorrido))
            print(f"Tiempo total: {tiempo}")
            auxiliares.guardar_viaje_kml(archivo, recorrido, latitudes)
        else:
            print("No se encontro recorrido")
    else:
        print("No se encontro recorrido")


## comando 4
def reducir_caminos(grafo, ruta, latitudes):
    """Pre: Recibe un grafo y una ruta
    Post: Imprime el peso total del arbol de tendido minimo"""
    ciudades, peso_total = libreria_tp.tm_prim(grafo)
    auxiliares.guardar_pj(ciudades, ruta, len(ciudades), latitudes)
    print(f"Peso total: {peso_total}")
