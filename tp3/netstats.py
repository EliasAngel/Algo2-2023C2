#!/usr/bin/python3

import funciones_tp
import grafo
import sys

sys.setrecursionlimit(500000)

COMANDOS = ["listar_operaciones", "camino",
            "mas_importantes", "conectados",
            "ciclo", "lectura",
            "diametro", "rango",
            "navegacion", "clustering",
            "comunidad"]


def cargar_grafo(archivo):
    grafo_nuevo = grafo.Grafo(dirigido=True, vertices_iniciales=[])
    with open(archivo) as archivo:
        for linea in archivo:
            linea = linea.strip().split("\t")
            grafo_nuevo.agregar_vertice(linea[0])
        archivo.seek(0)
        for linea in archivo:
            linea = linea.strip().split("\t")
            for link in linea[1:]:
                grafo_nuevo.agregar_arista(linea[0], link)
    return grafo_nuevo


def comando():
    while True:
        com = input()
        com = com.split()
        if com[0] in COMANDOS:
            if com == "diametro":
                return com, None
            if com == "clustering":
                if len(com) == 1:
                    return com, None
            parametros = com[1:]
            parametros = " ".join(parametros)
            parametros = parametros.split(",")
            return com, parametros


def main():
    ruta_archivo = sys.argv[1]
    grafo = cargar_grafo(ruta_archivo)
    while True:
        com, param = comando()
        if com[0] == "listar_operaciones":
            funciones_tp.listar_operaciones()
        elif com[0] == "camino":
            cam, costo = funciones_tp.camino_mas_corto(grafo, param[0], param[1])
            print(cam)
            if costo is not None:
                print("Costo:", costo)
        elif com[0] == "mas_importantes":
            pass
        elif com[0] == "conectados":
            componente = funciones_tp.conectividad(grafo, param[0])
            print(", ".join(componente))
        elif com[0] == "ciclo":
            pass
        elif com[0] == "lectura":
            camino = funciones_tp.lectura_2_am(grafo, param)
            if camino:
                print(", ".join(camino))
            else:
                print("No existe forma de leer las paginas en orden")
        elif com[0] == "diametro":
            camino, diametro = funciones_tp.diametro(grafo)
            print(camino)
            print("Costo:", diametro)
        elif com[0] == "rango":
            print(funciones_tp.todos_en_rango(grafo, param[0], param[1]))
        elif com[0] == "comunidad":
            resultado = funciones_tp.label_propagation(grafo, param[0])
            resultado = (", ".join(resultado))
            print(resultado)
        elif com[0] == "navegacion":
            print(funciones_tp.navegacion_primero(grafo, param[0]))
        elif com[0] == "clustering":
            resultado = funciones_tp.clustering(grafo, param[0])
            print(f"{resultado:.3f}")


main()
