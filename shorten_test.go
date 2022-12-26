package shorten

import (
	"errors"
	"testing"
)

type mockProvider struct{}

func (m *mockProvider) S(url string) (string, error) {
	return url, nil
}

func TestNewInstance(t *testing.T) {
	s := New()
	_, err := s.Shorten(&UrlParams{Destination: "asad"})

	if !errors.Is(err, ErrNoProviderAvailable) {
		t.Error("[TestNewInstance]: Incorrect error")
	}
}

func TestShortenWithProvider(t *testing.T) {
	s := New()
	p := &mockProvider{}
	s.AddProvider("test", p)
	u, _ := s.Shorten(&UrlParams{Destination: "http://test"})

	if u.Link != "http://test" {
		t.Errorf("[TestNewInstance] - invalid shorten url. Wanted: %s, Got: %s", "http://test", u.Link)
	}

	u, _ = s.With("test").Shorten(&UrlParams{Destination: "http://test"})

	if u.Link != "http://test" {
		t.Errorf("[TestNewInstance] - invalid shorten url. Wanted: %s, Got: %s", "http://test", u.Link)
	}
}

func TestProviderFunctionalities(t *testing.T) {
	s := New()
	p := &mockProvider{}
	s.AddProvider("test", p)

	if len(s.GetProviders()) != 1 {
		t.Errorf("[TestProviderFunctionalities]: Incorrect numbers of providers. ")
	}

	if s.getFirstProvider() == nil {
		t.Errorf("[TestProviderFunctionalities]: Incorrect numbers of providers. ")
	}

	if s.getProvider() == nil {
		t.Errorf("[TestProviderFunctionalities]: Incorrect numbers of providers. ")
	}

	s.DeleteProvider("test")

	if len(s.GetProviders()) != 0 {
		t.Errorf("[TestProviderFunctionalities]: Incorrect numbers of providers. ")
	}
}
