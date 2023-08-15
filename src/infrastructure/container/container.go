package container

import (
	applicationCli "Calculator/src/application/cli"
	"Calculator/src/application/web/controller"
	"Calculator/src/domain"
	"Calculator/src/domain/operators"
	"Calculator/src/domain/values"
	"Calculator/src/infrastructure/cli"
	"Calculator/src/infrastructure/web"
)

type container struct {
	cli *applicationCli.Cli
}

func NewContainer() *container {
	return &container{}
}

func (c container) GetCli() *applicationCli.Cli {
	if c.cli == nil {
		calculator := domain.NewCalculator(
			[]domain.Operator{
				operators.Addition{},
				operators.Subtraction{},
			},
			values.Value(0),
		)

		c.cli = applicationCli.NewCli(
			cli.NewInput(),
			cli.NewOutput(),
			calculator,
			controller.NewCalculation(calculator),
			web.NewServer(),
		)
	}

	return c.cli
}
