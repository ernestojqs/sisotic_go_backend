package gqlErrors

import "github.com/pjmd89/gogql/lib/gql/definitionError"

var (
	ERROR_EMPTY_OR_INVALID_ID             = definitionError.ErrorDescriptor{Code: "001", Message: "Empty or invalid input id", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_PHONE_NUMBER            = definitionError.ErrorDescriptor{Code: "002", Message: "Empty or invalid phone number format", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_EMAIL                   = definitionError.ErrorDescriptor{Code: "003", Message: "Empty or invalid email format", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_MINIMUM_CHARACTERS     = definitionError.ErrorDescriptor{Code: "004", Message: "The password must be a minimum of 8 characters", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_JOB_TITLE_IN_DB           = definitionError.ErrorDescriptor{Code: "005", Message: "Job title not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_JOB_AREA_IN_DB            = definitionError.ErrorDescriptor{Code: "006", Message: "Job area not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_JOB_TITLE_IS_REQUERIDED         = definitionError.ErrorDescriptor{Code: "007", Message: "The user must have a job title related to him", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_USER_IN_DB               = definitionError.ErrorDescriptor{Code: "008", Message: "Error when trying to create user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_USER_IN_DB               = definitionError.ErrorDescriptor{Code: "009", Message: "Error when trying to update user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_USER_IN_DB                = definitionError.ErrorDescriptor{Code: "010", Message: "User not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_USER_NOT_LOGGED                 = definitionError.ErrorDescriptor{Code: "011", Message: "USER NOT LOGGED IN", Level: definitionError.LEVEL_FATAL}
	ERROR_ACCESS_DENIED                   = definitionError.ErrorDescriptor{Code: "012", Message: "Acces denied", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_JOB_TITLE_IN_DB          = definitionError.ErrorDescriptor{Code: "013", Message: "Error when trying to create job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_JOB_TITLE_IN_DB          = definitionError.ErrorDescriptor{Code: "014", Message: "Error when trying to update job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_JOB_TITLE_IN_DB          = definitionError.ErrorDescriptor{Code: "015", Message: "Error when trying to delete job title in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_JOB_AREA_IN_DB           = definitionError.ErrorDescriptor{Code: "016", Message: "Error when trying to create job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_JOB_AREA_IN_DB           = definitionError.ErrorDescriptor{Code: "017", Message: "Error when trying to update job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_JOB_AREA_IN_DB           = definitionError.ErrorDescriptor{Code: "018", Message: "Error when trying to delete job area in database", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_COLLEGE_DEPENDENCY_IN_DB = definitionError.ErrorDescriptor{Code: "019", Message: "Error when trying to create college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_COLLEGE_DEPENDENCY_IN_DB = definitionError.ErrorDescriptor{Code: "020", Message: "Error when trying to update college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_COLLEGE_DEPENDENCY_IN_DB = definitionError.ErrorDescriptor{Code: "021", Message: "Error when trying to delete college dependecy in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_COLLEGE_DEPENDENCY_IN_DB  = definitionError.ErrorDescriptor{Code: "022", Message: "College dependency not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_NOT_ENCRYPTED          = definitionError.ErrorDescriptor{Code: "023", Message: "Failed to generate hashed password", Level: definitionError.LEVEL_FATAL}
	ERROR_CANNOT_COMPARE_PASSWORDS        = definitionError.ErrorDescriptor{Code: "024", Message: "Passwords could not be compared", Level: definitionError.LEVEL_FATAL}
	ERROR_PASSWORD_NOT_MATCH              = definitionError.ErrorDescriptor{Code: "025", Message: "Password does not match", Level: definitionError.LEVEL_FATAL}
	ERROR_MANY_LOGIN_VALUES               = definitionError.ErrorDescriptor{Code: "026", Message: "Use email or phone number, not both", Level: definitionError.LEVEL_FATAL}
	ERROR_EMAIL_EXISTS                    = definitionError.ErrorDescriptor{Code: "027", Message: "Email is not available", Level: definitionError.LEVEL_FATAL}
	ERROR_PHONE_NUMBER_EXISTS             = definitionError.ErrorDescriptor{Code: "028", Message: "Phone number is not available", Level: definitionError.LEVEL_FATAL}
	ERROR_INSERT_VERIFICATION_CODE_IN_DB  = definitionError.ErrorDescriptor{Code: "029", Message: "Error when trying to create verification code in database", Level: definitionError.LEVEL_FATAL}
	ERROR_QUERY_VERIFICATION_CODE_IN_DB   = definitionError.ErrorDescriptor{Code: "030", Message: "verification code not found in database", Level: definitionError.LEVEL_FATAL}
	ERROR_UPDATE_VERIFICATION_CODE_IN_DB  = definitionError.ErrorDescriptor{Code: "031", Message: "Error when trying to update verification code in database", Level: definitionError.LEVEL_FATAL}
	ERROR_CODE_IS_NOT_EQUAL               = definitionError.ErrorDescriptor{Code: "032", Message: "Error code is not equal", Level: definitionError.LEVEL_FATAL}
	ERROR_EXPIRED_CODE                    = definitionError.ErrorDescriptor{Code: "033", Message: "Error code is expired", Level: definitionError.LEVEL_FATAL}
	ERROR_EMAIL_IS_NOT_VERIFIED           = definitionError.ErrorDescriptor{Code: "034", Message: "User's email is not verified", Level: definitionError.LEVEL_FATAL}
	ERROR_LON_AND_LAT_FORMAT              = definitionError.ErrorDescriptor{Code: "035", Message: "latitude values range between 90 and -90 (180 degrees), while longitude values range between 180 and -180 (360 degrees). Check if one of these fields meets the above", Level: definitionError.LEVEL_FATAL}
	ERROR_INVALID_DATETIME                = definitionError.ErrorDescriptor{Code: "036", Message: "invalid datetime format", Level: definitionError.LEVEL_FATAL}
	ERROR_DELETE_USER_IN_DB               = definitionError.ErrorDescriptor{Code: "037", Message: "Error when trying to delete user in database", Level: definitionError.LEVEL_FATAL}
	ERROR_FILE_TYPE_BLANK                 = definitionError.ErrorDescriptor{Code: "038", Message: "file type can't be blank", Level: definitionError.LEVEL_FATAL}
)
