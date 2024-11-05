package utils

import (
	"crypto/rand"
	"math/big"
	"reflect"
	"strings"

	"github.com/pjmd89/goutils/systemutils"
)

var Argon2 *systemutils.Argon2

func GenerateToken(length int) (token string) {
	digits := make([]string, length)
	max := big.NewInt(9)
	for i := 0; i < length; i++ {
		n, _ := rand.Int(rand.Reader, max)
		digits[i] = n.String()
	}
	token = strings.Join(digits, "")

	return
}

// only works with gql tags
func GetGqlField(gqlTagName string, obj any) (fieldValue any) {

	reflectStruct := reflect.TypeOf(obj)

	if reflectStruct.Kind() != reflect.Struct {
		return
	}

	for i := 0; i < reflectStruct.NumField(); i++ {
		structField := reflectStruct.Field(i)
		tag := structField.Tag.Get("gql")
		tagValues := strings.Split(tag, ",")
		fieldName := strings.Split(tagValues[0], "=")
		if fieldName[1] == gqlTagName {
			fieldValue = reflect.ValueOf(obj).Field(i).Interface()
			break
		}
	}

	return
}
