package constants

// service -
const (
	ServiceName        string = "BACKEND_SERVICES"          // change it here
	ServiceLogFileName string = "backend_services_logs.log" // change it here
)

// CUSTOM ERROR CODES FOR messageCode field
const (
	TestCode               string = "TEST_CODE"                  // dummy code
	NoTokenInRequestHeader string = "NO_TOKEN_IN_REQUEST_HEADER" // header not containing auth token
	TokenDecodeFail        string = "TOKEN_DECODE_FAIL"          // decoding of token to get email fails
	EmptyEmailFromToken    string = "EMPTY_EMAIL_FROM_TOKEN"     // token decoded successfully but no email from it
	FetchDBDataError       string = "FETCH_DB_DATA_ERROR"        // error while fetching from database table
	InsertDBDataError      string = "INSERT_DB_DATA_ERROR"       // error while inserting in database table
	UpdateDBDataError      string = "UPDATE_DB_DATA_ERROR"       // error while updating in database table
	DeleteDBDataError      string = "DELETE_DB_DATA_ERROR"       // error while deleting from database table
	UserNotFound           string = "USER_NOT_FOUND"             // user does not exist in our user table
	UserOrOrgNotFound      string = "USER_OR_ORG_NOT_FOUND"      // user or organization does not exist
	EmptyDBData            string = "EMPTY_DB_DATA"              // empty set returned from database
	RequestParseError      string = "REQUEST_PARSE_ERROR"        // error during parsing of request into Go struct
	ValueNegativeError     string = "VALUE_NEGATIVE_ERROR"       // when a certain value is expected to be positive
	DateEmptyError         string = "DATE_EMPTY_ERROR"           // when date value is expected to be present
	DateFormatError        string = "DATE_FORMAT_ERROR"          // wrong format of the date field
	RequestBodyEmptyError  string = "REQUEST_BODY_EMPTY_ERROR"   // no data provided in the request body
	UnauthorizedAccess     string = "UNAUTHORIZED_ACCESS"        // un authorized access request
	VendorAPIError         string = "VENDOR_API_ERROR"           // vendor api call failed error
	FetchRemoteDataError   string = "FETCH_REMOTE_DATA_ERR"      // fetch remote data error
	InvalidParams          string = "INVALID_PARAMS"
)

// Response Messages
const (
	InvalidUserMessage string = "Invalid User"            // Error message when user is not
	LoginSuccessfull   string = "User sigin successfully" // Login Success
	EnterEmail         string = "Please enter email"
	EnterPassword      string = "Please enter password"
	EnterPhoneNo       string = "Please enter phone no"
	EnterUserName      string = "Please enter enter name"
)

// Response Messages Constants
const (
	MessagesInvalidBlockId string = "Invalid Block Id" // Error message when user is not
)
