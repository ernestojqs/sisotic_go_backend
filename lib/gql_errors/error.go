package gqlErrors

import "github.com/pjmd89/gogql/lib/gql/definitionError"

var (
	ERROR_EMPTY_OR_INVALID_ID              = definitionError.ErrorDescriptor{Code: "001", Message: "Empty or invalid input id", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_PHONE_NUMBER             = definitionError.ErrorDescriptor{Code: "002", Message: "Empty or invalid phone number format", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_EMAIL                    = definitionError.ErrorDescriptor{Code: "003", Message: "Empty or invalid email format", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_MINIMUM_CHARACTERS      = definitionError.ErrorDescriptor{Code: "004", Message: "The password must be a minimum of 8 characters", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_JOB_TITLE_IN_DB            = definitionError.ErrorDescriptor{Code: "005", Message: "Job title not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_JOB_AREA_IN_DB             = definitionError.ErrorDescriptor{Code: "006", Message: "Job area not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_JOB_TITLE_IS_REQUERIDED          = definitionError.ErrorDescriptor{Code: "007", Message: "The user must have a job title related to him", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_USER_IN_DB                = definitionError.ErrorDescriptor{Code: "008", Message: "Error when trying to create user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_USER_IN_DB                = definitionError.ErrorDescriptor{Code: "009", Message: "Error when trying to update user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_USER_IN_DB                 = definitionError.ErrorDescriptor{Code: "010", Message: "User not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_USER_NOT_LOGGED                  = definitionError.ErrorDescriptor{Code: "011", Message: "USER NOT LOGGED IN", Level: definitionError.LEVEL_FATAL}
	ERROR_ACCESS_DENIED                    = definitionError.ErrorDescriptor{Code: "012", Message: "Acces denied", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_JOB_TITLE_IN_DB           = definitionError.ErrorDescriptor{Code: "013", Message: "Error when trying to create job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_JOB_TITLE_IN_DB           = definitionError.ErrorDescriptor{Code: "014", Message: "Error when trying to update job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_JOB_TITLE_IN_DB           = definitionError.ErrorDescriptor{Code: "015", Message: "Error when trying to delete job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_JOB_AREA_IN_DB            = definitionError.ErrorDescriptor{Code: "016", Message: "Error when trying to create job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_JOB_AREA_IN_DB            = definitionError.ErrorDescriptor{Code: "017", Message: "Error when trying to update job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_JOB_AREA_IN_DB            = definitionError.ErrorDescriptor{Code: "018", Message: "Error when trying to delete job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_COLLEGE_DEPENDENCY_IN_DB  = definitionError.ErrorDescriptor{Code: "019", Message: "Error when trying to create college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_COLLEGE_DEPENDENCY_IN_DB  = definitionError.ErrorDescriptor{Code: "020", Message: "Error when trying to update college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_COLLEGE_DEPENDENCY_IN_DB  = definitionError.ErrorDescriptor{Code: "021", Message: "Error when trying to delete college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_COLLEGE_DEPENDENCY_IN_DB   = definitionError.ErrorDescriptor{Code: "022", Message: "College dependency not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_NOT_ENCRYPTED           = definitionError.ErrorDescriptor{Code: "023", Message: "Failed to generate hashed password", Level: definitionError.LEVEL_FATAL}
	ERROR_CANNOT_COMPARE_PASSWORDS         = definitionError.ErrorDescriptor{Code: "024", Message: "Passwords could not be compared", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_NOT_MATCH               = definitionError.ErrorDescriptor{Code: "025", Message: "Password does not match", Level: definitionError.LEVEL_FATAL}
	ERROR_MANY_LOGIN_VALUES                = definitionError.ErrorDescriptor{Code: "026", Message: "Use email or phone number, not both", Level: definitionError.LEVEL_FATAL}
	ERROR_EMAIL_EXISTS                     = definitionError.ErrorDescriptor{Code: "027", Message: "Email is not available", Level: definitionError.LEVEL_FATAL}
	ERROR_PHONE_NUMBER_EXISTS              = definitionError.ErrorDescriptor{Code: "028", Message: "Phone number is not available", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_VERIFICATION_CODE_IN_DB   = definitionError.ErrorDescriptor{Code: "029", Message: "Error when trying to create verification code in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_VERIFICATION_CODE_IN_DB    = definitionError.ErrorDescriptor{Code: "030", Message: "verification code not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_VERIFICATION_CODE_IN_DB   = definitionError.ErrorDescriptor{Code: "031", Message: "Error when trying to update verification code in database", Level: definitionError.LEVEL_FATAL}
	ERROR_CODE_IS_NOT_EQUAL                = definitionError.ErrorDescriptor{Code: "032", Message: "Error code is not equal", Level: definitionError.LEVEL_FATAL}
	ERROR_EXPIRED_CODE                     = definitionError.ErrorDescriptor{Code: "033", Message: "Error code is expired", Level: definitionError.LEVEL_FATAL}
	ERROR_EMAIL_IS_NOT_VERIFIED            = definitionError.ErrorDescriptor{Code: "034", Message: "User's email is not verified", Level: definitionError.LEVEL_FATAL}
	ERROR_LON_AND_LAT_FORMAT               = definitionError.ErrorDescriptor{Code: "035", Message: "latitude values range between 90 and -90 (180 degrees), while longitude values range between 180 and -180 (360 degrees). Check if one of these fields meets the above", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_DATETIME                 = definitionError.ErrorDescriptor{Code: "036", Message: "invalid datetime format", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_USER_IN_DB                = definitionError.ErrorDescriptor{Code: "037", Message: "Error when trying to delete user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_FILE_TYPE_BLANK                  = definitionError.ErrorDescriptor{Code: "038", Message: "file type can't be blank", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_DEVICE_IN_DB              = definitionError.ErrorDescriptor{Code: "039", Message: "Error when trying to create device in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_DEVICE_IN_DB              = definitionError.ErrorDescriptor{Code: "040", Message: "Error when trying to update device in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_DEVICE_IN_DB              = definitionError.ErrorDescriptor{Code: "041", Message: "Error when trying to delete device in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_DEVICE_IN_DB               = definitionError.ErrorDescriptor{Code: "042", Message: "Device not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_RESOLVER_ACTIVITY_IN_DB   = definitionError.ErrorDescriptor{Code: "043", Message: "Error when trying to create resolver activity in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_RESOLVER_ACTIVITY_IN_DB   = definitionError.ErrorDescriptor{Code: "044", Message: "Error when trying to update resolver activity in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_RESOLVER_ACTIVITY_IN_DB   = definitionError.ErrorDescriptor{Code: "045", Message: "Error when trying to delete resolver activity in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_RESOLVER_ACTIVITY_IN_DB    = definitionError.ErrorDescriptor{Code: "046", Message: "Resolver activity not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_TECHNICAL_DIAGNOSIS_IN_DB = definitionError.ErrorDescriptor{Code: "047", Message: "Error when trying to create technical diagnosis in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_TECHNICAL_DIAGNOSIS_IN_DB = definitionError.ErrorDescriptor{Code: "048", Message: "Error when trying to update technical diagnosis in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_TECHNICAL_DIAGNOSIS_IN_DB = definitionError.ErrorDescriptor{Code: "049", Message: "Error when trying to delete technical diagnosis in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_TECHNICAL_DIAGNOSIS_IN_DB  = definitionError.ErrorDescriptor{Code: "050", Message: "Technical diagnosis not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_TECHNICAL_DIAGNOSIS_REQUERIDED   = definitionError.ErrorDescriptor{Code: "051", Message: "If the place is local and the device will be supported, information on technical diagnostics and resolver activities is required", Level: definitionError.LEVEL_FATAL}
	ERROR_USER_ALREADY_HAS_PERMISSION      = definitionError.ErrorDescriptor{Code: "052", Message: "The user already have the provided permission", Level: definitionError.LEVEL_FATAL}
	ERROR_NOT_FOUND_PERMISSION_IN_USER     = definitionError.ErrorDescriptor{Code: "053", Message: "The user does not have the provided permission", Level: definitionError.LEVEL_FATAL}
	ERROR_USER_ACCOUNT_BLOCKED             = definitionError.ErrorDescriptor{Code: "054", Message: "The user account has been blocked", Level: definitionError.LEVEL_FATAL}
)
