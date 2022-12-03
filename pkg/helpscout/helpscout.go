package helpscout

import "net/url"

type HelpScout struct {
	apiKey          string
	apiURL          *url.URL
	rateLimitRemain int
	rateLimitReset  int
}

const (
	userAgent = "helpscout-toolkit/0.1"
)

func NewHelpScout(apiKey string) *HelpScout {
	hs := &HelpScout{
		apiKey: apiKey,
	}

	hs.apiURL, _ = url.Parse("https://docsapi.helpscout.net/v1/")

	return hs
}
