/*
Package shorten
*/
package shorten

// New returns shorten instance.
func New() *Shorten {
	return &Shorten{
		providers: map[string]Provider{},
	}
}

type Shorten struct {
	providers map[string]Provider
	provider  Provider
}

func (s *Shorten) Shorten(u *UrlParams) (*Url, error) {
	var url Url
	var p Provider

	p = s.getProvider()

	if p == nil {
		return &url, ErrNoProviderAvailable
	}

	result, err := p.S(u.Destination)

	if err != nil {
		return &url, err
	}

	url.Link = result

	return &url, nil
}
func (s *Shorten) AddProvider(name string, provider Provider) {
	s.providers[name] = provider
}

func (s *Shorten) DeleteProvider(name string) {
	delete(s.providers, name)
}

func (s *Shorten) GetProviders() []Provider {
	var p []Provider
	for _, v := range s.providers {
		p = append(p, v)
	}

	return p
}

func (s *Shorten) getProvider() Provider {
	var p Provider

	if s.provider != nil {
		p = s.provider
	}

	// provider not set. Use the first provider.
	if p == nil && len(s.providers) > 0 {
		p = s.getFirstProvider()
	}

	return p
}

func (s *Shorten) getFirstProvider() Provider {
	var p Provider

	for _, v := range s.providers {
		p = v
		break
	}

	return p
}

func (s *Shorten) With(provider string) *Shorten {
	s.provider = s.providers[provider]

	return s
}
