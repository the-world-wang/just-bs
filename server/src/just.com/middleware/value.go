package middleware
import "errors"

const (
	MLEARNING_CONTENT = "MLEARNING_CONTENT"
	MLEARNING__RESPONSE = "MLEARNING__RESPONSE"
	MLEARNING__EMAIL = "MLEARNING__EMAIL"
	MIDDLEWARE_TOKEN = "TOKEN_TEST"
)
var (
	NO_TOKEN_ERR = errors.New("NO_TOKEN_ERR")
	NO_AUTHORITATION = errors.New("NO_AUTHORITATION")
)