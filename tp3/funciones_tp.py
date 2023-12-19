import grafo
import funciones_grafo
import random

MAX_PAGINAS = 20
MINIMOS_ADYACENTES = 2
ITERACIONES_COMUNIDADES = 50


# Funcion auxiliar que genera un camino utilizando el diccionario de padres obtenido por bfs
def camino(padres, origen, final):
    camino = [final]
    actual = final
    while actual != origen:
        actual = padres[actual]
        camino.append(actual)
    return camino[::-1]


# Funcion auxiliar para unir una lista con el formato pedido por el tp
def unir_visitados(visitados):
    return " -> ".join(visitados)


# Imprime las operaciones incluidas en el tp
def listar_operaciones():
    print("camino")  # 1
    # print("mas_importantes")  # 3
    print("conectados")  # 2
    # print("ciclo")  # 3
    print("lectura")  # 2
    print("diametro")  # 1
    print("rango")  # 1
    print("comunidad")  # 2
    print("navegacion")  # 1
    print("clustering")  # 2


# Realiza un recorrido bfs sobre el grafo y devuelve el camino mas corto que una a las paginas
# pasadas por parametro
def camino_mas_corto(grafo, origen, fin):
    padres, orden = funciones_grafo.bfs(grafo, origen)
    if fin in orden.keys():
        cam = camino(padres, origen, fin)
        costo = len(cam)-1
        cam = (unir_visitados(cam))
        return cam, costo
    else:
        return "No se encontro recorrido", None


# Busca las componentes fuertemente conexas del grafo y devuelvela que contenga la pagina
# pasada por parametro
def conectividad(grafo, pagina):
    resultados = funciones_grafo.buscar_cfc(grafo, pagina)
    for componente in resultados:
        if pagina in componente:
            return componente


# Genera un nuevo grafo con los vertices pasados por parametro y devuelve el orden topologico
# invertido para poder cumplir la consigna
def lectura_2_am(g, paginas):
    grafo_nuevo = grafo.Grafo(dirigido=True, vertices_iniciales=paginas)
    for pagina in paginas:
        for w in g.adyacentes(pagina):
            if w in paginas:
                grafo_nuevo.agregar_arista(pagina, w)
    orden = funciones_grafo.orden_topologico(grafo_nuevo)
    if not orden:  # HAY CICLO
        return None
    return orden[::-1]


# Obtiene el diametro de toda la red buscando cual es el bfs que tenga el maximo valor en el
# diccionario de orden
def diametro(g):
    recorrido = None
    diam = 0
    origen = None
    orden_diametro = None
    for v in g:
        resultado_bfs, orden = funciones_grafo.bfs(g, v)
        if max(orden.values()) > diam:
            diam = max(orden.values())
            recorrido = resultado_bfs
            origen = v
            orden_diametro = orden
    vertice_final = max(orden_diametro, key=lambda k: orden_diametro[k])
    return (unir_visitados(camino(recorrido, origen, vertice_final))), diam


# Realiza un recorrido bfs sobre el grafo y devuelve cuantos vertices estan a x vertices de
# distancia
def todos_en_rango(grafo, pagina, rango):
    _, orden = funciones_grafo.bfs(grafo, pagina)
    contador = 0
    for valor in orden.values():
        if valor == int(rango):
            contador += 1
    return contador


# Obtiene las comunidades del grafo mediante el algoritmo label propagation y devuelve
# la comunidad a la que pertenece la pagina pasada por parametro
def label_propagation(grafo, pagina):
    label = {}
    i = 0
    comunidades = {}
    for v in grafo:
        comunidades[i] = {v}
        label[v] = i
        i += 1
    vertices = grafo.obtener_vertices()
    v_ent = funciones_grafo.obtener_vertices_entrada(grafo)
    random.shuffle(vertices)
    for _ in range(ITERACIONES_COMUNIDADES):
        for v in vertices:
            comunidad = funciones_grafo.max_freq(v_ent[v], label)
            comunidades[comunidad].add(v)
            if label[v] in comunidades[label[v]]:
                comunidades[label[v]].remove(v)
            label[v] = comunidad
    for comunidad in comunidades.values():
        if pagina in comunidad:
            return comunidad


# Parte de la pagina pasada por parametro y se mueve al vertice conectado por el primer link
# hasta que no hubiera mas links o se cumplan las 20 iteraciones
def navegacion_primero(grafo, origen):
    visitados = [origen]
    pagina = origen
    for i in range(MAX_PAGINAS):
        links = grafo.adyacentes(pagina)
        if len(links) > 0:
            pagina = links[0]
            visitados.append(pagina)
        else:
            break
    return unir_visitados(visitados)


# Calcula el coeficiente de clustering de una pagina o de el grafo, segun los parametros recibidos
def clustering(grafo, pagina):
    if pagina:
        resultado = _clustering(grafo, pagina)
        return resultado
    coeficiente_global = 0
    vertices = len(grafo.obtener_vertices())
    for v in grafo:
        coeficiente_global += _clustering(grafo, v)
    resultado = (1/vertices) * coeficiente_global
    return resultado


# Funcion auxiliar que se encarga de calcular el coeficiente de clustering de una pagina
def _clustering(grafo, pagina):
    contador = 0
    adyacentes_pagina = grafo.adyacentes(pagina)
    if len(adyacentes_pagina) < MINIMOS_ADYACENTES:
        return contador
    g_salida = len(adyacentes_pagina)
    for v in adyacentes_pagina:
        if v != pagina:
            for w in adyacentes_pagina:
                if v != w and grafo.estan_unidos(v, w):
                    contador += 1
    numerador = contador
    denominador = g_salida*(g_salida-1)
    return numerador / denominador
