package cli

import (
	"Calculator/src/application/web"
)

type Cli struct {
	input      input
	output     output
	calculator calculator
	controller web.Controller
	server     web.Server
}

func NewCli(
	input input,
	output output,
	calculator calculator,
	controller web.Controller,
	server web.Server,
) *Cli {
	return &Cli{input: input, output: output, calculator: calculator, controller: controller, server: server}
}

func (c Cli) Run() {
	expression := c.input.GetParameter("c")

	if expression != "" {
		c.calculate(expression)
		return
	}

	port := c.input.GetParameter("s")
	if port != "" {
		c.server.Run(port, c.controller)
	}
}

func (c Cli) calculate(value string) {
	result, err := c.calculator.Calculate(value)
	if err != nil {
		c.output.PrintLine(err.Error())
	} else {
		c.output.PrintLine(result)
	}
}
