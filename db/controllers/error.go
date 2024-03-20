package controllers

type ControllerError struct {
	message    string
	statusCode int
}

func NewControllerError(message string, statusCode int) *ControllerError {
	return &ControllerError{message, statusCode}
}

func (e ControllerError) Error() string {
	return e.message
}

func (e ControllerError) StatusCode() int {
	return e.statusCode
}
