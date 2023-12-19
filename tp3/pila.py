class Pila:

    # Crea una pila vacía
    def __init__(self):
        self.tope = None

    # Apila un elemento en el tope de la pila
    def apilar(self, dato):
        nodo = _Nodo(dato, self.tope)
        self.tope = nodo

    # Desapila el tope de la pila y lo devuelve
    def desapilar(self):
        if self.esta_vacia():
            raise ValueError("pila vacía")
        dato = self.tope.dato
        self.tope = self.tope.prox
        return dato

    # Devuelve el tope de la pila
    def ver_tope(self):
        if self.esta_vacia():
            raise ValueError("pila vacía")
        return self.tope.dato

    # Devuelve el estado de la pila
    def esta_vacia(self):
        return self.tope is None


class _Nodo:
    def __init__(self, dato, prox=None):
        self.dato = dato
        self.prox = prox