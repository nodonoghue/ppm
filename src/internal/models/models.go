package models

type CommandFlags struct {
	NumVariants  *int
	Length       *int
	NumUpperCase *int
	NumLowerCase *int
	NumNumbers   *int
	NumSpecial   *int
	IsHelp       *bool
}
