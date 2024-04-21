package response

type AppError struct {
	Message string
}

type DuplicateData struct {
	Message string
}

type NotFoundError struct {
	Message string
}

func (e AppError) Error() string {
	return e.Message
}

func (e NotFoundError) Error() string {
	return e.Message
}

func NewAppError(message string) error {
	return AppError{
		Message: message,
	}
}

func NewNotFoundError(message string) error {
	return NotFoundError{
		Message: message,
	}
}
