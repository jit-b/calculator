package web

type Server interface {
	Run(port string, controller Controller)
}
