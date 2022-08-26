package customerror

type CustomError struct {
	Message    any
	StatusCode int
}

func NewError(message any, code int) *CustomError {
	return &CustomError{
		message, code,
	}
}

var (
	NotFoundError = &CustomError{
		Message:    "Order not found",
		StatusCode: 404,
	}
	InternalServerError = &CustomError{
		Message:    "Internal server error occurred",
		StatusCode: 500,
	}
	InvalidBodyError = &CustomError{
		Message:    "Enter a correct order body",
		StatusCode: 400,
	}
	CustomerNotFoundError = &CustomError{
		Message:    "Customer is not found",
		StatusCode: 404,
	}
	AddressNotFoundError = &CustomError{
		Message:    "Customer does not have a address",
		StatusCode: 404,
	}
)
