package secrets

import (
	"regexp"
)

func init() {
	RegisterMatcher(&linkedInSecretKey{})
}

type linkedInSecretKey struct{}

func (*linkedInSecretKey) Type() string {
	return "linkedin_secret_key"
}

func (*linkedInSecretKey) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?im)linkedin(.{0,20})?['\"][0-9a-z]{16}['\"]`),
	}
}

func (*linkedInSecretKey) Verify(secret string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}
