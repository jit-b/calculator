package cli

type input interface {
	GetParameter(string) string
}
