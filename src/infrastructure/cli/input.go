package cli

import "os"

type input struct {
}

func NewInput() *input {
	return &input{}
}

func (i input) GetParameter(name string) string {
	if len(os.Args) > 2 {
		name = "-" + name
		if os.Args[1] == name {
			return os.Args[2]
		}
	}
	return ""
}
