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
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

func (o *VerificationCode) sendVerificationCodeMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	email := input["email"].(string)
	token := utils.GenerateToken(4)
	input["submittedCode"] = token
	verificationCodeWhere := map[string]any{"email": email}
	verificationCode, _ := o.model.Read(verificationCodeWhere, nil)
	if len(verificationCode.([]models.VerificationCode)) == 0 {
		result, rerr := o.model.Create(input, nil)
		if rerr != nil {
			lib.Logs.System.Error().Println(rerr.Error())
			err = definitionError.NewError(gqlErrors.ERROR_INSERT_VERIFICATION_CODE_IN_DB, definitionError.ExtensionError{"Insert Error": rerr.Error()})
			return
		}
		r = result
	} else {
		if verificationCode.([]models.VerificationCode)[0].WasVerified {
			input["wasVerified"] = false
		}
		verificationCode, rerr := o.model.Update(input, verificationCodeWhere, nil)
		if rerr != nil || len(verificationCode.([]models.VerificationCode)) == 0 {
			lib.Logs.System.Error().Println(gqlErrors.ERROR_UPDATE_VERIFICATION_CODE_IN_DB)
			err = definitionError.NewError(gqlErrors.ERROR_UPDATE_VERIFICATION_CODE_IN_DB, nil)
			return
		}
		r = verificationCode.([]models.VerificationCode)[0]
	}
	//aqui se envia el correo electrónico
	return
}
