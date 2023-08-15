package cli

type calculator interface {
	Calculate(expression string) (float32, error)
}
