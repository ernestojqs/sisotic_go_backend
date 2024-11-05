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
package logged

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/http"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *Logged) loginQuery(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)

	if input["password"] == nil {
		r, _ = utils.GetSession(info.SessionID)
		return
	}
	userWhere := map[string]any{}
	if (input["phoneNumber"] == nil && input["email"] == nil) || ((input["phoneNumber"] != nil && input["phoneNumber"].(string) != "") && (input["email"] != nil && input["email"].(string) != "")) {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_PASSWORD_NOT_MATCH)
		err = definitionError.NewError(gqlErrors.ERROR_PASSWORD_NOT_MATCH, nil)
		return
	}

	if input["phoneNumber"] != nil && input["phoneNumber"].(string) != "" {
		userWhere["phoneNumber"] = input["phoneNumber"].(string)
	}

	if input["email"] != nil && input["email"].(string) != "" {
		userWhere["email"] = input["email"].(string)
	}

	signedUpUser, err := o.readUser(userWhere)
	if err != nil {
		return
	}

	isEqual, rerr := utils.Argon2.Compare(input["password"].(string), string(signedUpUser.Password))
	if rerr != nil {
		lib.Logs.System.Warning().Println(rerr.Error())
		err = definitionError.NewError(gqlErrors.ERROR_CANNOT_COMPARE_PASSWORDS, map[string]any{"compare Error": rerr.Error()})
		return
	}

	if !isEqual {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_PASSWORD_NOT_MATCH)
		err = definitionError.NewError(gqlErrors.ERROR_PASSWORD_NOT_MATCH, nil)
		return
	}

	sess := utils.Session{
		UserID:   signedUpUser.Id,
		UserRole: string(signedUpUser.Role),
	}

	http.Session.Set(sess)
	r = signedUpUser

	return
}

func (o *Logged) readUser(userWhere map[string]any) (signedUpUser models.User, err definitionError.GQLError) {
	user, rerr := o.userModel.Read(userWhere, nil)
	if rerr != nil || len(user.([]models.User)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_USER_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_USER_IN_DB, nil)
		return
	}
	return user.([]models.User)[0], nil
}

func (o *Logged) logoutQuery(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	session, err := utils.GetSession(info.SessionID)
	var loggedUser any
	if err == nil {
		loggedUser, _ = o.readUser(map[string]any{"_id": session.UserID})
		session = utils.Session{
			UserID:   primitive.NilObjectID,
			UserRole: "",
		}
		http.Session.Set(session)
	}
	r = loggedUser
	return
}
