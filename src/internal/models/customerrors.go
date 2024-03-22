package models

type GeneralError struct {
	Message       string
	FunctionName  string
	Configuration CommandFlags
}

func (generationError GeneralError) Error() string {
	return generationError.Message
}
