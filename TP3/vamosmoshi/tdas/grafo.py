class Grafo:
    def __init__(self, dirigido=False):
        self.dirigido = dirigido
        self.vertices = []
        self.adyacente = {}

    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices.append(v)
            self.adyacente[v] = {}

    def obtener_vertices(self):
        return self.vertices

    def adyacentes(self, v):
        return list(self.adyacente[v])

    def agregar_arista(self, v, w, peso=1):
        self.adyacente[v][w] = peso
        if not self.dirigido:
            self.adyacente[w][v] = peso

    def es_arista(self, v, w):
        return w in self.adyacente[v]

    def peso(self, v, w):
        if self.es_arista(v, w):
            return self.adyacente[v][w]

    def __str__(self):
        return str(self.adyacente)
