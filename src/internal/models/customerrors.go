package models

type GenerationError struct {
	Message       string
	FunctionName  string
	Configuration CommandFlags
}

func (generationError *GenerationError) Error() string {
	return generationError.Message
}
