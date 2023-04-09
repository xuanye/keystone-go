package exception

const (
	ACCESS_DENIED_CODE       = -10242403 // Access denied
	COMMON_BIZ_ERROR_CODE    = -10250000 // Common business error
	REQUIRE_PARAM_EMPTY_CODE = -10250001 // Required parameter is empty
	PARAM_FORMAT_WRONG_CODE  = -10250002 // Parameter format is wrong
	REQUIRED_AUTH_CODE       = -10250003 // Authorization is required
	BIZ_VALIDATE_FAIL_CODE   = -10250004 // Business validation failed
	DATA_NOT_FOUND_CODE      = -10250404 // Data not found
	OPERATE_FORBIDDEN_CODE   = -10250403 // Operation forbidden
	INTERNAL_ERROR_CODE      = -10250500 // Internal error
	SUCCESS_BUT              = -10250600 // Success, but with some extra information
	TOO_MANY_ACTION          = -10250601 // Too many actions in a certain time frame
	REMOTE_FAIL_CODE         = -10250602 // Remote service call failed
	REMOTE_ERROR_CODE        = -10250603 // Remote service returns error
)
