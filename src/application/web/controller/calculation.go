package controller

import "Calculator/src/application/web"

type calculation struct {
	calculator calculator
}

func NewCalculation(calculator calculator) *calculation {
	return &calculation{calculator: calculator}
}

func (c calculation) Handle(response web.Response, request web.Request) {
	//TODO implement me
	panic("implement me")
}
