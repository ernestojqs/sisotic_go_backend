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
package main

import (
	"embed"
	"slices"

	"otic/lib"
	"otic/lib/generateauth"
	gqlErrors "otic/lib/gql_errors"
	"otic/lib/utils"
	"otic/resolvers/directives"
	"otic/resolvers/objectTypes/changelog"
	"otic/resolvers/objectTypes/collegedependency"
	"otic/resolvers/objectTypes/device"
	"otic/resolvers/objectTypes/edgecollegedependency"
	"otic/resolvers/objectTypes/edgedevice"
	"otic/resolvers/objectTypes/edgejobarea"
	"otic/resolvers/objectTypes/edgejobtitle"
	"otic/resolvers/objectTypes/edgeresolveractivity"
	"otic/resolvers/objectTypes/edgetask"
	"otic/resolvers/objectTypes/edgetechnicaldiagnosis"
	"otic/resolvers/objectTypes/edgeuser"
	"otic/resolvers/objectTypes/jobarea"
	"otic/resolvers/objectTypes/jobtitle"
	"otic/resolvers/objectTypes/logged"
	"otic/resolvers/objectTypes/observation"
	"otic/resolvers/objectTypes/pageinfo"
	"otic/resolvers/objectTypes/resolveractivity"
	"otic/resolvers/objectTypes/systeminfo"
	"otic/resolvers/objectTypes/task"
	"otic/resolvers/objectTypes/technicaldiagnosis"
	"otic/resolvers/objectTypes/transferevidence"
	"otic/resolvers/objectTypes/ubication"
	"otic/resolvers/objectTypes/upload"
	"otic/resolvers/objectTypes/user"
	"otic/resolvers/objectTypes/verificationcode"
	"otic/resolvers/scalars"

	"github.com/pjmd89/gogql/lib/gql"
	"github.com/pjmd89/gogql/lib/gql/definitionError"
	"github.com/pjmd89/gogql/lib/http"
	"github.com/pjmd89/gogql/lib/rest"
	"github.com/pjmd89/goutils/systemutils"
	"github.com/pjmd89/goutils/systemutils/debugmode"
	"github.com/pjmd89/mongomodel/mongomodel"
)

var (
	//go:embed "schema"
	embedFS embed.FS
)

var myConfig = lib.Config()
var systemLog = systemutils.NewLog(myConfig.SystemLog)
var accessLog = systemutils.NewLog(myConfig.AccessLog)
var logs = systemutils.Logs{System: systemLog, Access: accessLog}
var db = mongomodel.NewConn(&myConfig.DBConfigFile)
var schema = gql.Init(embedFS, "schema")

var restfull = rest.Init()
var myHttp = http.Init(logs, myConfig.HTTPConfigFile).SetGql(schema).SetRest(restfull)

func main() {
	utils.Argon2 = systemutils.NewArgon2()
	lib.MyConfig = myConfig
	lib.Logs = logs
	systemLog.Info().Println("debugmode: ", debugmode.Enabled)
	if generateauth.Enabled {
		generateauth.GenerateAuth(schema.GetSchema(), "schema")
	}
	myHttp.Start()
}
func OnDB(currentDB string, currentCollection string) (r string) {
	r = currentDB
	return
}
func httpOnSession() (r interface{}) {
	return
}
func httpCheckOrigin(url http.URL) (r bool, info interface{}) {
	r = true
	return
}
func httpOnBegin(url http.URL, httpPath *http.Path, originData interface{}, uid string) (r bool) {
	return true
}
func httpOnFinish(isErr bool, uid string) {}
func OnIntrospection() (err definitionError.GQLError) {
	return
}
func OnAuthorizate(authInfo gql.AuthorizateInfo) (err definitionError.GQLError) {
	//El el rol puede venir de una variable de sesion
	//La sesion debes definirla de manera que se adapte a tu necesidad
	access, isIn := lib.Auth[authInfo.SrcType][authInfo.DstType][authInfo.Resolver]
	if !isIn {
		return
	}
	session, err := utils.GetSession(authInfo.SessionID)
	if err != nil {
		return
	}

	if !slices.Contains(access, session.UserRole) {
		lib.Logs.System.Warning().Println(gqlErrors.ERROR_ACCESS_DENIED)
		err = definitionError.NewError(gqlErrors.ERROR_ACCESS_DENIED, nil)
		return
	}

	return
}
func init() {
	db.(*mongomodel.MongoDBConn).OnDatabase = OnDB
	myHttp.OnSession = httpOnSession
	myHttp.CheckOrigin = httpCheckOrigin
	myHttp.OnBegin = httpOnBegin
	myHttp.OnFinish = httpOnFinish
	schema.OnIntrospection = OnIntrospection
	schema.OnAuthorizate = OnAuthorizate
	schema.ObjectType("ChangeLog", changelog.NewChangeLog(db))
	schema.ObjectType("PageInfo", pageinfo.NewPageInfo(db))
	schema.ObjectType("Logged", logged.NewLogged(db))
	schema.ObjectType("ResolverActivity", resolveractivity.NewResolverActivity(db))
	schema.ObjectType("Task", task.NewTask(db))
	schema.ObjectType("User", user.NewUser(db))
	schema.ObjectType("EdgeUser", edgeuser.NewEdgeUser(db))
	schema.ObjectType("Device", device.NewDevice(db))
	schema.ObjectType("EdgeJobArea", edgejobarea.NewEdgeJobArea(db))
	schema.ObjectType("JobTitle", jobtitle.NewJobTitle(db))
	schema.ObjectType("TechnicalDiagnosis", technicaldiagnosis.NewTechnicalDiagnosis(db))
	schema.ObjectType("SystemInfo", systeminfo.NewSystemInfo(db))
	schema.ObjectType("VerificationCode", verificationcode.NewVerificationCode(db))
	schema.ObjectType("EdgeCollegeDependency", edgecollegedependency.NewEdgeCollegeDependency(db))
	schema.ObjectType("EdgeJobTitle", edgejobtitle.NewEdgeJobTitle(db))
	schema.ObjectType("Observation", observation.NewObservation(db))
	schema.ObjectType("EdgeTask", edgetask.NewEdgeTask(db))
	schema.ObjectType("TransferEvidence", transferevidence.NewTransferEvidence(db))
	schema.ObjectType("CollegeDependency", collegedependency.NewCollegeDependency(db))
	schema.ObjectType("EdgeDevice", edgedevice.NewEdgeDevice(db))
	schema.ObjectType("Ubication", ubication.NewUbication(db))
	schema.ObjectType("Upload", upload.NewUpload(db))
	schema.ObjectType("EdgeTechnicalDiagnosis", edgetechnicaldiagnosis.NewEdgeTechnicalDiagnosis(db))
	schema.ObjectType("JobArea", jobarea.NewJobArea(db))
	schema.ObjectType("EdgeResolverActivity", edgeresolveractivity.NewEdgeResolverActivity(db))
	schema.Scalar("Email", scalars.NewEmailScalar())
	schema.Scalar("PhoneNumber", scalars.NewPhoneNumberScalar())
	schema.Scalar("DateTime", scalars.NewDateTimeScalar())
	schema.Scalar("Latitude", scalars.NewLatitudeScalar())
	schema.Scalar("GeoAddress", scalars.NewGeoAddressScalar())
	schema.Scalar("Longitude", scalars.NewLongitudeScalar())
	schema.Scalar("ID", scalars.NewIDScalar())
	schema.Directive("paginate", directives.NewPaginate())
	schema.Directive("sort", directives.NewSort())
	schema.Directive("search", directives.NewSearch(schema.GetScalars()))
}
