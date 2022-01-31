package kekasigohelper

import "errors"

//authValidation
var (
	ErrInvalidToken             = errors.New("Invalid authorization token")
	ErrInvalidTokenType         = errors.New("Authorization token type does not match")
	ErrNotMatchTokenCredentials = errors.New("Authorization token credentials do not match")
	ErrInvalidTokenCredentials  = errors.New("Invalid authorization token credentials")
	ErrInvalidTokenExpired      = errors.New("Authorization token has expired")
	ErrUsername                 = errors.New("Please Check again your Username")
	ErrPassword                 = errors.New("Please Check again your Password")
	LoginFailedMessage          = errors.New("please check username and password")
)

var (
	ErrInternalServerError = errors.New("Internal Server Error")
	ErrNotFound            = errors.New("Your requested Item is not found")
	ErrUnAuthorize         = errors.New("Unauthorize")
	ErrBadParamInput       = errors.New("Bad Request")
	ErrPublicKey           = errors.New("invalid Public Key")
	ErrInvalidDataType     = errors.New("invalid data type")
	ErrIsRequired          = errors.New("is required")
	ErrInvalidValue        = errors.New("invalid value")
	ErrForbidden           = errors.New("you don't have permission to access this resource")
)

var (
	GeneralSuccess    = "Success"
	ErrGeneralMessage = errors.New("something wrong")
	ErrMessageCaptcha = errors.New("Incorrect captcha. Please try again.")
)
