/*
 * Generado por gqlgenerate.
 *
 * Este archivo puede contener errores, de ser asi, coloca el issue en el repositorio de github
 * https://github.com/pjmd89/gogql
 *
 * Estos arvhivos corren riesgo de sobreescritura, por ese motivo gqlgnerate crea una carpeta llamada generate, asi que,
 * copia todas las carpetas que estan dentro de la carpeta generate y pegalas en la carpeta raiz de tu proyecto.
 *
 * gqlgenerate no creara archivos en la carpeta raiz de tu modulo porque puedes sufrir perdida de informacion.
 */
package enums

type OperatorEnum string

const (
	OPERATORENUM_EQ     OperatorEnum = "eq"
	OPERATORENUM_GT                  = "gt"
	OPERATORENUM_GTE                 = "gte"
	OPERATORENUM_LT                  = "lt"
	OPERATORENUM_LTE                 = "lte"
	OPERATORENUM_NE                  = "ne"
	OPERATORENUM_ALL                 = "all"
	OPERATORENUM_IN                  = "in"
	OPERATORENUM_NIN                 = "nin"
	OPERATORENUM_REGEX               = "regex"
	OPERATORENUM_SIZE                = "size"
	OPERATORENUM_MOD                 = "mod"
	OPERATORENUM_EXISTS              = "exists"
)
