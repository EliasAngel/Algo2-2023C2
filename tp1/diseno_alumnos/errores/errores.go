package errores

import "fmt"

type ErrorLeerArchivo struct{}

func (e ErrorLeerArchivo) Error() string {
	return "ERROR: Lectura de archivos"
}

// Implemento un tipo de error personalizado para Error de Parametros
type ErrorParametros struct{}

func (e ErrorParametros) Error() string {
	return "ERROR: Faltan parámetros"
}

// Implemento un tipo de error personalizado para Cadena invalida
type CadenaInvalida struct{}

func (e *CadenaInvalida) Error() string {
	return "ERROR: Cadena invalida"
}

// Implemento un tipo de error personalizado para DNI invalido
type DNIError struct{}

func (e DNIError) Error() string {
	return "ERROR: DNI incorrecto"
}

// Implemento un tipo de error personalizado para Votante fuera de rango
type DNIFueraPadron struct{}

func (e DNIFueraPadron) Error() string {
	return "ERROR: DNI fuera del padrón"
}

// Implemento un tipo de error personalizado para Error de Fila Vacia
type FilaVacia struct{}

func (e FilaVacia) Error() string {
	return "ERROR: Fila vacía"
}

type ErrorVotanteFraudulento struct {
	Dni int
}

func (e ErrorVotanteFraudulento) Error() string {
	return fmt.Sprintf("ERROR: Votante FRAUDULENTO: %d", e.Dni)
}

// Implemento un tipo de error personalizado para Error de Voto Invalido
type ErrorTipoVoto struct{}

func (e ErrorTipoVoto) Error() string {
	return "ERROR: Tipo de voto inválido"
}

// Implemento un tipo de error personalizado para Error de Alternativa Invalida
type ErrorAlternativaInvalida struct{}

func (e ErrorAlternativaInvalida) Error() string {
	return "ERROR: Alternativa inválida"
}

// Implemento un tipo de error personalizado para Error de comando invalido
type ErrorComandoInvalido struct{}

func (e ErrorComandoInvalido) Error() string {
	return "ERROR: Comando inválido"
}

type ErrorNoHayVotosAnteriores struct{}

func (e ErrorNoHayVotosAnteriores) Error() string {
	return "ERROR: Sin voto a deshacer"
}

type ErrorCiudadanosSinVotar struct{}

func (e ErrorCiudadanosSinVotar) Error() string {
	return "ERROR: Ciudadanos sin terminar de votar"
}
