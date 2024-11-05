package lib

import "github.com/pjmd89/gogql/lib/gql"

var (
	TYPE_COLLEGEDEPENDENCY      gql.TypeName = "CollegeDependency"
	TYPE_DEVICE                 gql.TypeName = "Device"
	TYPE_EDGECOLLEGEDEPENDENCY  gql.TypeName = "EdgeCollegeDependency"
	TYPE_EDGEDEVICE             gql.TypeName = "EdgeDevice"
	TYPE_EDGEJOBAREA            gql.TypeName = "EdgeJobArea"
	TYPE_EDGEJOBTITLE           gql.TypeName = "EdgeJobTitle"
	TYPE_EDGERESOLVERACTIVITY   gql.TypeName = "EdgeResolverActivity"
	TYPE_EDGETASK               gql.TypeName = "EdgeTask"
	TYPE_EDGETECHNICALDIAGNOSIS gql.TypeName = "EdgeTechnicalDiagnosis"
	TYPE_EDGEUSER               gql.TypeName = "EdgeUser"
	TYPE_JOBAREA                gql.TypeName = "JobArea"
	TYPE_JOBTITLE               gql.TypeName = "JobTitle"
	TYPE_LOGGED                 gql.TypeName = "Logged"
	TYPE_MUTATION               gql.TypeName = "Mutation"
	TYPE_QUERY                  gql.TypeName = "Query"
	TYPE_RESOLVERACTIVITY       gql.TypeName = "ResolverActivity"
	TYPE_TASK                   gql.TypeName = "Task"
	TYPE_TECHNICALDIAGNOSIS     gql.TypeName = "TechnicalDiagnosis"
	TYPE_UPLOAD                 gql.TypeName = "Upload"
	TYPE_USER                   gql.TypeName = "User"
)

var (
	RESOLVER_CANCELUPLOAD             gql.ResolverName = "cancelUpload"
	RESOLVER_CHECKFILESIZE            gql.ResolverName = "checkFileSize"
	RESOLVER_CREATECOLLEGEDEPENDENCY  gql.ResolverName = "createCollegeDependency"
	RESOLVER_CREATEDEVICES            gql.ResolverName = "createDevices"
	RESOLVER_CREATEJOBAREA            gql.ResolverName = "createJobArea"
	RESOLVER_CREATEJOBTITLE           gql.ResolverName = "createJobTitle"
	RESOLVER_CREATERESOLVERACTIVITY   gql.ResolverName = "createResolverActivity"
	RESOLVER_CREATETASK               gql.ResolverName = "createTask"
	RESOLVER_CREATETECHNICALDIAGNOSIS gql.ResolverName = "createTechnicalDiagnosis"
	RESOLVER_DELETECOLLEGEDEPENDENCY  gql.ResolverName = "deleteCollegeDependency"
	RESOLVER_DELETEDEVICE             gql.ResolverName = "deleteDevice"
	RESOLVER_DELETEUSER               gql.ResolverName = "deleteUser"
	RESOLVER_GETCOLLEGEDEPENDENCIES   gql.ResolverName = "getCollegeDependencies"
	RESOLVER_GETDEVICES               gql.ResolverName = "getDevices"
	RESOLVER_GETJOBAREAS              gql.ResolverName = "getJobAreas"
	RESOLVER_GETJOBTITLES             gql.ResolverName = "getJobTitles"
	RESOLVER_GETRESOLVERACTIVITIES    gql.ResolverName = "getResolverActivities"
	RESOLVER_GETTASKS                 gql.ResolverName = "getTasks"
	RESOLVER_GETTECHNICALDIAGNOSIS    gql.ResolverName = "getTechnicalDiagnosis"
	RESOLVER_GETUSERS                 gql.ResolverName = "getUsers"
	RESOLVER_LOGOUT                   gql.ResolverName = "logout"
	RESOLVER_UPDATECOLLEGEDEPENDENCY  gql.ResolverName = "updateCollegeDependency"
	RESOLVER_UPDATEDEVICE             gql.ResolverName = "updateDevice"
	RESOLVER_UPDATEJOBAREA            gql.ResolverName = "updateJobArea"
	RESOLVER_UPDATEJOBTITLE           gql.ResolverName = "updateJobTitle"
	RESOLVER_UPDATETASK               gql.ResolverName = "updateTask"
	RESOLVER_UPDATEUSER               gql.ResolverName = "updateUser"
	RESOLVER_UPLOAD                   gql.ResolverName = "upload"
)

var Auth = gql.Grant{
	TYPE_MUTATION: {
		TYPE_COLLEGEDEPENDENCY: {
			RESOLVER_CREATECOLLEGEDEPENDENCY: {`admin`, `user`},
			RESOLVER_DELETECOLLEGEDEPENDENCY: {`admin`, `user`},
			RESOLVER_UPDATECOLLEGEDEPENDENCY: {`admin`, `user`},
		},
		TYPE_DEVICE: {
			RESOLVER_CREATEDEVICES: {`admin`, `user`},
			RESOLVER_DELETEDEVICE:  {`admin`, `user`},
			RESOLVER_UPDATEDEVICE:  {`admin`, `user`},
		},
		TYPE_JOBAREA: {
			RESOLVER_CREATEJOBAREA: {`admin`, `user`},
			RESOLVER_UPDATEJOBAREA: {`admin`, `user`},
		},
		TYPE_JOBTITLE: {
			RESOLVER_CREATEJOBTITLE: {`admin`, `user`},
			RESOLVER_UPDATEJOBTITLE: {`admin`, `user`},
		},
		TYPE_RESOLVERACTIVITY: {
			RESOLVER_CREATERESOLVERACTIVITY: {`admin`, `user`},
		},
		TYPE_TASK: {
			RESOLVER_CREATETASK: {`admin`, `user`},
			RESOLVER_UPDATETASK: {`admin`, `user`},
		},
		TYPE_TECHNICALDIAGNOSIS: {
			RESOLVER_CREATETECHNICALDIAGNOSIS: {`admin`, `user`},
		},
		TYPE_UPLOAD: {
			RESOLVER_CANCELUPLOAD: {`admin`, `user`},
			RESOLVER_UPLOAD:       {`admin`, `user`},
		},
		TYPE_USER: {
			RESOLVER_DELETEUSER: {`admin`, `user`},
			RESOLVER_UPDATEUSER: {`admin`, `user`},
		},
	},
	TYPE_QUERY: {
		TYPE_EDGECOLLEGEDEPENDENCY: {
			RESOLVER_GETCOLLEGEDEPENDENCIES: {`admin`, `user`},
		},
		TYPE_EDGEDEVICE: {
			RESOLVER_GETDEVICES: {`admin`, `user`},
		},
		TYPE_EDGEJOBAREA: {
			RESOLVER_GETJOBAREAS: {`admin`, `user`},
		},
		TYPE_EDGEJOBTITLE: {
			RESOLVER_GETJOBTITLES: {`admin`, `user`},
		},
		TYPE_EDGERESOLVERACTIVITY: {
			RESOLVER_GETRESOLVERACTIVITIES: {`admin`, `user`},
		},
		TYPE_EDGETASK: {
			RESOLVER_GETTASKS: {`admin`, `user`},
		},
		TYPE_EDGETECHNICALDIAGNOSIS: {
			RESOLVER_GETTECHNICALDIAGNOSIS: {`admin`, `user`},
		},
		TYPE_EDGEUSER: {
			RESOLVER_GETUSERS: {`admin`, `user`},
		},
		TYPE_LOGGED: {
			RESOLVER_LOGOUT: {`admin`, `user`},
		},
		TYPE_UPLOAD: {
			RESOLVER_CHECKFILESIZE: {`admin`, `user`},
		},
	},
}
