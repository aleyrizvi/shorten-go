package shorten

type Url struct {
	// Link is the shortened link returned by the provider
	Link string
}

type UrlParams struct {
	// Destination is the long url that needs to be shorten
	Destination string
}

// String returns the pointer of string.
func String(v string) *string {
	return &v
}
