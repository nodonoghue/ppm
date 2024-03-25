package models

var Variants map[int]string

type CommandFlags struct {
	NumVariants  *int
	Length       *int
	NumUpperCase *int
	NumLowerCase *int
	NumNumbers   *int
	NumSpecial   *int
	IsHelp       *bool
}

type BucketDrop struct {
	Name     string `json:"Name"`
	URI      string `json:"Uri"`
	Username string `json:"Username"`
	Password string `json:"Password"`
}
