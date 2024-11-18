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
	"otic/models"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/resolvers"
	"github.com/pjmd89/goutils/dbutils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	model                 models.User
	jobTitleModel         models.JobTitle
	verificationCodeModel models.VerificationCode
}

func NewUser(db dbutils.DBInterface) (o resolvers.ObjectTypeInterface) {
	o = &User{
		model:                 models.User{},
		jobTitleModel:         models.JobTitle{},
		verificationCodeModel: models.VerificationCode{},
	}
	o.(*User).model.Init(models.User{}, db)
	o.(*User).jobTitleModel.Init(models.JobTitle{}, db)
	o.(*User).verificationCodeModel.Init(models.VerificationCode{}, db)
	return o
}
func (o *User) Resolver(info resolvers.ResolverInfo) (r resolvers.DataReturn, err definitionError.GQLError) {
	switch info.Operation {
	case "query":
		switch info.Resolver {
		case "edges":
			r, err = o.edges(info)
		case "user", "recordUser", "receiverUser", "resolverUser", "diagnosticUser", "autor":
			r, err = o.user(info)
		}
	case "mutation":
		switch info.Resolver {
		case "createUser":
			r, err = o.createUserMutation(info)
		case "updateUser":
			r, err = o.updateUserMutation(info)
		case "deleteUser":
			r, err = o.deleteUserMutation(info)
		case "resetUserPassword":
			r, err = o.resetUserPasswordMutation(info)
		case "assignPermissions":
			r, err = o.assignPermissions(info)
		case "user", "recordUser", "receiverUser", "resolverUser", "diagnosticUser", "autor":
			r, err = o.user(info)
		}
	}
	return
}
func (o *User) Subscribe(info resolvers.ResolverInfo) (r bool) {
	return
}

func CheckPermissions(userID primitive.ObjectID, moduleName, resolver string, db dbutils.DBInterface) (r bool) {
	userObj := NewUser(db).(*User)
	userWhere := bson.M{
		"_id":                    userID,
		"permissions.name":       moduleName,
		"permissions.operations": bson.M{"$in": bson.A{resolver}},
	}
	if result, _ := userObj.model.Read(userWhere, nil); len(result.([]models.User)) > 0 {
		r = true
	}
	return
}
