package utils

import (
	"log"
	"otic/lib"
	gqlErrors "otic/lib/gql_errors"

	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/http"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Session struct {
	UserID   primitive.ObjectID
	UserRole string
}

func GetSession(sessionID string) (sessionData Session, err definitionError.GQLError) {
	sess, rerr := http.Session.GetByID(sessionID)

	if sess == nil || sess.(Session).UserID.IsZero() {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_USER_NOT_LOGGED)
		log.Println(rerr)
		err = definitionError.NewError(gqlErrors.ERROR_USER_NOT_LOGGED, nil)
		return
	}
	sessionData = sess.(Session)

	return
}
