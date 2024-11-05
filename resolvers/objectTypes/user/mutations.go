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
package user

import (
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/models"
	"otic/models/enums"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (o *User) validateUniqueField(m map[string]any, userID any) (err definitionError.GQLError) {
	errr := map[string]definitionError.ErrorDescriptor{"email": gqlErrors.ERROR_EMAIL_EXISTS, "phoneNumber": gqlErrors.ERROR_PHONE_NUMBER_EXISTS}
	for key, value := range m {
		if value != nil && value.(string) != "" {
			usersWhere := bson.M{"email": value.(string)}
			if userID != nil {
				usersWhere["_id"] = bson.M{"$ne": userID.(primitive.ObjectID)}
			}
			users, rerr := o.model.Count(usersWhere, nil)
			if users > 0 || rerr != nil {
				lib.Logs.System.Warning().Println(errr[key])
				err = definitionError.NewError(errr[key], nil)
				return
			}
		}
	}
	return
}

func (o *User) createUserMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	dbUsers, _ := o.model.Count(nil, nil)
	input["role"] = enums.ROLEENUM_ADMIN
	if dbUsers > 0 {
		input["role"] = enums.ROLEENUM_USER
		if err = o.validateUniqueField(map[string]any{"email": input["email"].(string), "phoneNumber": input["phoneNumber"]}, nil); err != nil {
			return
		}
		if input["jobTitle"] == nil {
			lib.Logs.System.Warning().Println(gqlErrors.ERROR_JOB_TITLE_IS_REQUERIDED)
			err = definitionError.NewError(gqlErrors.ERROR_JOB_TITLE_IS_REQUERIDED, nil)
			return
		}
		if _, err = o.readJobTitle(input["jobTitle"].(primitive.ObjectID)); err != nil {
			return
		}
	}
	if input["password"], err = o.hashPassword(input["password"].(string)); err != nil {
		return
	}

	r, rerr := o.model.Create(input, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(rerr.Error())
		err = definitionError.NewError(gqlErrors.ERROR_INSERT_USER_IN_DB, definitionError.ExtensionError{"DB error": rerr.Error()})
	}

	return
}

func (o *User) readJobTitle(jobTitleID primitive.ObjectID) (dbjobTitle models.JobTitle, err definitionError.GQLError) {
	jobTitleWhere := map[string]any{"_id": jobTitleID}
	jobTitle, rerr := o.jobTitleModel.Read(jobTitleWhere, nil)
	if rerr != nil || len(jobTitle.([]models.JobTitle)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_JOB_TITLE_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_JOB_TITLE_IN_DB, nil)
		return
	}
	dbjobTitle = jobTitle.([]models.JobTitle)[0]

	return
}

func (o *User) hashPassword(inputPassword string) (hashedPassword string, hashErr definitionError.GQLError) {
	hashedPassword, err := utils.Argon2.Generate(inputPassword)
	if err != nil {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_PASSWORD_NOT_ENCRYPTED)
		hashErr = definitionError.NewError(gqlErrors.ERROR_PASSWORD_NOT_ENCRYPTED, nil)
	}

	return
}

func (o *User) updateUserMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	input := info.Args["input"].(map[string]any)
	userWhere := map[string]any{"_id": input["_id"].(primitive.ObjectID)}
	user, rerr := o.model.Read(userWhere, nil)
	if rerr != nil || len(user.([]models.User)) == 0 {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_QUERY_USER_IN_DB)
		err = definitionError.NewError(gqlErrors.ERROR_QUERY_USER_IN_DB, nil)
		return
	}

	if (input["email"] != nil && input["email"].(string) != "") || (input["phoneNumber"] != nil && input["phoneNumber"].(string) != "") {
		if err = o.validateUniqueField(map[string]any{"email": input["email"], "phoneNumber": input["phoneNumber"]}, nil); err != nil {
			return
		}
	}

	if input["password"] != nil && input["password"].(string) != "" {
		if input["password"], err = o.hashPassword(input["password"].(string)); err != nil {
			return
		}
	}

	if input["jobTitle"] != nil {
		if _, err = o.readJobTitle(input["jobTitle"].(primitive.ObjectID)); err != nil {
			return
		}
	}

	result, rerr := o.model.Update(input, userWhere, nil)
	if rerr != nil {
		lib.Logs.System.Warning().Println(rerr.Error())
		err = definitionError.NewError(gqlErrors.ERROR_UPDATE_USER_IN_DB, definitionError.ExtensionError{"DB error": rerr.Error()})
		return
	}

	r = result.([]models.User)[0]

	return
}
func (o *User) deleteUserMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	return
}
func (o *User) resetUserPasswordMutation(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {

	return
}
