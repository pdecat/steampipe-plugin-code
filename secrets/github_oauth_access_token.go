package secrets

import "regexp"

func init() {
	RegisterMatcher(&githubOAuthAccessToken{})
}

type githubOAuthAccessToken struct{}

func (*githubOAuthAccessToken) Type() string {
	return "github_oauth_access_token"
}

func (*githubOAuthAccessToken) DenyList() []*regexp.Regexp {
	return []*regexp.Regexp{
		regexp.MustCompile(`(?m)gho_[0-9a-zA-Z]{36}`),
	}
}

func (*githubOAuthAccessToken) Verify(secret string) (VerifiedResult, error) {
	return UNVERIFIED, nil
}
