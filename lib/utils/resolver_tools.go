package utils

import (
	"crypto/rand"
	"math/big"
	"otic/resolvers/scalars"
	"reflect"
	"strings"

	"github.com/google/uuid"
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

func GenerateTokenFromUUID(lenght int, escapeHyphen bool) (token string) {
	uuidToken := uuid.New()
	token = uuidToken.String()

	if escapeHyphen {
		token = strings.ReplaceAll(token, "-", "")
	}
	if lenght >= 4 && lenght <= len(token) {
		token = token[:lenght]
	}

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

func ParseArayDBObj(obj any) (parsedArrayObj []any) {
	if reflect.TypeOf(obj).Kind() != reflect.Slice {
		return
	}

	value := reflect.ValueOf(obj)
	for i := 0; i < value.Len(); i++ {
		parsedArrayObj = append(parsedArrayObj, ParseDBObj(value.Index(i)))
	}

	return
}

func ParseDBObj(obj any) (parsedObj any) {
	if obj == nil {
		return
	}

	mapObj := map[string]interface{}{}
	for i := 0; i < reflect.TypeOf(obj).NumField(); i++ {
		fieldName := reflect.TypeOf(obj).Field(i).Name
		fieldName = strings.ToLower(string(fieldName[0])) + fieldName[1:]
		fieldValue := reflect.ValueOf(obj).Field(i).Interface()
		switch val := fieldValue.(type) {
		case scalars.DateTime:
			mapObj[fieldName] = int64(val)
		default:
			if reflect.TypeOf(val).Kind() == reflect.Struct {
				val = ParseDBObj(val)
			}
			mapObj[fieldName] = val
		}
	}
	parsedObj = mapObj
	return
}
