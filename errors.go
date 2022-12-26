package shorten

import "errors"

var (
	ErrNoProviderAvailable = errors.New("no provider. use shorten.AddProvider()")
)
