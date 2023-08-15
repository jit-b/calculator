package cli

import "fmt"

type output struct {
}

func NewOutput() *output {
	return &output{}
}

func (o output) PrintLine(content any) {
	fmt.Println(content)
}
