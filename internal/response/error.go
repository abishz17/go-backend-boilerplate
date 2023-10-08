package response

type AppError struct {
	Message string
}

type NotFound struct {
	Model string
}

type DuplicateData struct {
	Message string
}

type NotFoundError struct {
	message string
}

func (a AppError) Error() string {
	return a.Message
}

func NewAppError(message string) error {
	return AppError{
		Message: message,
	}
}
