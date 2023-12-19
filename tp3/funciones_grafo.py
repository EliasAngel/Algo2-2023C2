from collections import deque
import pila


# Hace el recorrido bfs del grafo y devuelve un diccionario de las distancias
# Se añade un parametro n en caso de que se quiera limitar el recorrido de paginas
def bfs(grafo, origen):
    visitados = set()
    padres = {}
    orden = {}
    padres[origen] = None
    orden[origen] = 0
    cola = deque()
    cola.append(origen)
    visitados.add(origen)
    while not len(cola) == 0:
        v = cola.popleft()
        for w in grafo.adyacentes(v):
            if w not in visitados:
                padres[w] = v
                orden[w] = orden[v] + 1
                cola.append(w)
                visitados.add(w)
    return padres, orden


# Devuelve un orden topologico del grafo o None en caso de que haya ciclo
def orden_topologico(grafo):
    g_ent = grados_entrada(grafo)
    cola = deque()
    for v in grafo:
        if g_ent[v] == 0:
            cola.append(v)
    resultado = []
    while not len(cola) == 0:
        v = cola.popleft()
        resultado.append(v)
        for w in grafo.adyacentes(v):
            g_ent[w] -= 1
            if g_ent[w] == 0:
                cola.append(w)
    if len(resultado) != len(grafo.obtener_vertices()):
        return None  # HAY CICLO
    return resultado


# Obtiene los grados de entrada de los vertices del grafo
def grados_entrada(grafo):
    g_ent = {}
    for v in grafo:
        g_ent[v] = 0
    for v in grafo:
        for w in grafo.adyacentes(v):
            g_ent[w] += 1
    return g_ent


# Busca las componentes fuertemente conexas de un grafo
def buscar_cfc(grafo, pagina):
    resultados = []
    visitados = set()
    contador_global = [0]
    dfs_cfc(grafo, pagina, visitados, {}, {}, pila.Pila(), set(), resultados, contador_global)
    return resultados


# Se encarga de usar la busqueda tipo dfs para encontrar las componentes fuertemente conexas del grafo
def dfs_cfc(grafo, v, visitados, orden, mas_bajo, pila, apilados, cfcs, cont):
    orden[v] = mas_bajo[v] = cont[0]
    cont[0] += 1
    visitados.add(v)
    pila.apilar(v)
    apilados.add(v)
    for w in grafo.adyacentes(v):
        if w not in visitados:
            dfs_cfc(grafo, w, visitados, orden, mas_bajo, pila, apilados, cfcs, cont)
        if w in apilados:
            mas_bajo[v] = min(mas_bajo[v], mas_bajo[w])
    if orden[v] == mas_bajo[v]:
        nueva_cfc = []
        while True:
            w = pila.desapilar()
            apilados.remove(w)
            nueva_cfc.append(w)
            if w == v:
                break
        cfcs.append(nueva_cfc)


# Se encarga de verificar cual es la etiqueta que mas se repite en los vertices pasados por parametro
def max_freq(entradas, label):
    contador = {}
    for v in entradas:
        frecuencia = label[v]
        contador[frecuencia] = contador.get(frecuencia, 0) + 1
    return max(contador, key=contador.get)


# Devuelve un diccionario {vertice: entradas_al_vertice} con todos los vertices del grafo
def obtener_vertices_entrada(grafo):
    entradas = {}
    for v in grafo:
        entradas[v] = []
    for v in grafo:
        for w in grafo.adyacentes(v):
            entradas[w].append(v)
    return entradas
