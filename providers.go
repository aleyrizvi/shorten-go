package shorten

type Provider interface {
	S(url string) (string, error)
}
