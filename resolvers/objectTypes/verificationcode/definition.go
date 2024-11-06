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
package verificationcode

import (
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
)

type VerificationCode struct {
	model models.VerificationCode
}

func NewVerificationCode(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &VerificationCode{
		model: models.VerificationCode{},
	}
	o.(*VerificationCode).model.Init(models.VerificationCode{}, db)
	return o
}
func (o *VerificationCode) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "verifyCode":
			r, err = o.verifyCodeQuery(info)
		}
	case "mutation":
		switch info.Resolver {
		case "sendVerificationCode":
			r, err = o.sendVerificationCodeMutation(info)
		}
	}
	return
}
func (o *VerificationCode) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}
