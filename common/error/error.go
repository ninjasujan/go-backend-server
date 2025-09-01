package error

type ApiErrorType string

const (
	ValidationError     ApiErrorType = "validation_error"
	AuthenticationError ApiErrorType = "authentication_error"
	AuthorizationError  ApiErrorType = "authorization_error"
	NotFoundError       ApiErrorType = "not_found_error"
	InternalError       ApiErrorType = "internal_error"
	ConflictError       ApiErrorType = "conflict_error"
	BadRequestError     ApiErrorType = "bad_request_error"
)

type ApiError struct {
	Type    ApiErrorType `json:"type"`
	Status  int          `json:"status"`
	Message string       `json:"message"`
	Details string       `json:"details,omitempty"`
}

func NewApiError(errType ApiErrorType, status int, message string, details string) *ApiError {
	return &ApiError{
		Type:    errType,
		Status:  status,
		Message: message,
		Details: details,
	}
}

func NewValidationError(message string, details string) *ApiError {
	return NewApiError(ValidationError, 400, message, details)
}

func NewAuthenticationError(message string, details string) *ApiError {
	return NewApiError(AuthenticationError, 401, message, details)
}

func NewAuthorizationError(message string, details string) *ApiError {
	return NewApiError(AuthorizationError, 403, message, details)
}

func NewNotFoundError(message string, details string) *ApiError {
	return NewApiError(NotFoundError, 404, message, details)
}

func NewInternalError(message string, details string) *ApiError {
	return NewApiError(InternalError, 500, message, details)
}

func NewConflictError(message string, details string) *ApiError {
	return NewApiError(ConflictError, 409, message, details)
}

func NewBadRequestError(message string, details string) *ApiError {
	return NewApiError(BadRequestError, 400, message, details)
}
