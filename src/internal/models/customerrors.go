package models

type GeneralError struct {
	Message       string
	FunctionName  string
	Configuration CommandFlags
}

func (e GeneralError) Error() string {
	return e.Message
}
