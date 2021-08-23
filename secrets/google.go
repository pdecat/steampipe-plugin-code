package secrets

import "regexp"

func init() {
	RegisterMatcher(&googleApiKey{})
}

type googleApiKey struct{}

func (*googleApiKey) Type() string {
	return "google_api_key"
}

func (*googleApiKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`AIza[0-9A-Za-z\\-_]{35}`),
	}
}

func (*googleApiKey) Verify(secret string) (*bool, error) {
	return nil, nil
}
