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
	"otic/models"
	"time"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
)

func (o *VerificationCode) verifyCodeQuery(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	input := info.Args["input"].(map[string]any)
	email := input["email"].(string)
	where := map[string]any{"email": email}

	result, _ := o.model.Read(where, nil)
	if (len(result.([]models.VerificationCode)) == 0) || (input["isResetPassword"] == true && result.([]models.VerificationCode)[0].User.IsZero()) {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_VERIFICATION_CODE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_VERIFICATION_CODE_IN_DB, nil)
		return
	}

	verificationCode := result.([]models.VerificationCode)[0]
	if verificationCode.SubmittedCode != input["submittedCode"].(string) {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_CODE_IS_NOT_EQUAL)
		err = definitionError.NewError(gqlErrors.ERROR_CODE_IS_NOT_EQUAL, nil)
		return
	}
	/*
		if o.isExpired(int64(verificationCode.Created)) {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_EXPIRED_CODE)
			err = definitionError.NewError(gqlErrors.ERROR_EXPIRED_CODE, nil)
			return
		}
	*/
	verificationCodeInput := map[string]interface{}{"wasVerified": true}
	updated, _ := o.model.Update(verificationCodeInput, where, nil)

	r = updated
	return

}

func (o *VerificationCode) isExpired(dateTime int64) (r bool) {

	secondsElapsed := time.Since(time.Unix(dateTime, 0)).Seconds()
	//300 represents 5 minutes in unix time
	if secondsElapsed >= 300 {
		r = true
	}
	return
}
