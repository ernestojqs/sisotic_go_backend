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
 *
 * Este type fue obtenido del archivo:
 * .schema/objectTypes/transfer_evidence/type.gql:1
 */
package models

type TransferEvidence struct {
	PhotoPath    string `bson:"photoPath" gql:"name=photoPath"`
	TransferType string `bson:"transferType" gql:"name=transferType"`
	Id_number    string `bson:"id_number" gql:"name=id_number"`
}
