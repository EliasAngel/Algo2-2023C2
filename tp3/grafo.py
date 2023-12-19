import random


class Grafo:

    # Crea un grafo a partir de los valores pasados por parametro, dirigido bool y vertices
    # iniciales []
    def __init__(self, dirigido, vertices_iniciales):
        self.dirigido = dirigido
        self.vertices = dict.fromkeys(vertices_iniciales, {})
        for v in vertices_iniciales:
            self.vertices[v] = {}

    # Itera los vertices del grafo
    def __iter__(self):
        return iter(self.vertices)

    # Agrega vertices al grafo, y crea un diccionario para sus posteriores aristas
    def agregar_vertice(self, v):
        if v not in self.vertices:
            self.vertices[v] = {}

    # Borra un vertice de la lista de vertices, y las arists correspondientes al vertice
    def borrar_vertice(self, vertice):
        self.vertices.pop(vertice)

    # Conecta dos vertices con una arista, con un peso que se pasa por parametro, default 1
    def agregar_arista(self, v, w, peso=1):
        if not self.dirigido:
            self.vertices[w][v] = peso
        self.vertices[v][w] = peso

    # Borra una arista que conecta dos vertices
    def borrar_arista(self, v, w):
        if v in self.vertices and w in self.vertices[v]:
            self.vertices[v].pop(w)
            if not self.dirigido:
                self.vertices[w].pop(v)

    # Verifica si dos aristas estan unidas
    def estan_unidos(self, v, w):
        return w in self.vertices[v]

    # Devuelve el peso de la arista entre dos vertices
    def peso_arista(self, v, w):
        return self.vertices[v].get(w)

    # Obtiene todos los vertices del grafo
    def obtener_vertices(self):
        return list(self.vertices)

    # Devuelve un vertice aleatorio del grafo
    def vertice_aleatorio(self):
        return random.choice(list(self.vertices.keys()))

    # Devuelve una lista con todos los adyacentes del vertice pasado por parametro
    def adyacentes(self, v):
        adyacentes = []
        if v in self.vertices:
            for w in self.vertices[v].keys():
                adyacentes.append(w)
        return adyacentes
