#!/usr/bin/python3
from library import comandos
from library import auxiliares
import sys

sys.setrecursionlimit(10000)

from tdas import grafo

# constantes de COMANDOS
COMANDO_1 = "ir"
COMANDO_2 = "itinerario"
COMANDO_3 = "viaje"
COMANDO_4 = "reducir_caminos"


def main(grafo, latitudes):
    """
    DOC: Funcion principal del programa, se encarga de
    leer los comandos y llamar a las funciones correspondientes"""
    for entrada in sys.stdin:
        comando, params = comandos.lectura_comandos(entrada)
        if comando == COMANDO_1:
            if len(params) != 3:
                print("Error, uso: ir desde, hasta, ruta")
                continue
            comandos.ir(params[0], params[1], params[2], grafo, latitudes)
        if comando == COMANDO_2:
            if params[0] == "":
                print("Error, uso: itinerario ruta")
                continue
            comandos.itinerario(params[0], grafo)
        if comando == COMANDO_3:
            if len(params) != 2:
                print("Error, uso: viaje 'desde', 'hasta'")
                continue
            comandos.viaje(params[0], params[1], grafo, latitudes)
        if comando == COMANDO_4:
            if len(params) != 1:
                print("Error, uso: reducir_caminos ruta")
                continue
            comandos.reducir_caminos(grafo, params[0], latitudes)


if __name__ == "__main__":
    if len(sys.argv) != 2:
        print("Error, uso: ./vamosmoshi.py archivo.pj")
    else:
        archivo = sys.argv[1]
        grafo = grafo.Grafo()
        latitudes = {}
        auxiliares.agregar_ciudades(archivo, grafo, latitudes)
        main(grafo, latitudes)
