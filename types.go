package content

import (
	"net/url"
	"time"
)

// Title is the content title which can be retrieved in different ways
type Title interface {
	Original() string
	Clean() string
	OpenGraphTitle() (string, bool)
	ForSlug() string
}

// Summary is the content's description or summary which can be retrieved in different ways
type Summary interface {
	Original() string
	FirstSentenceOfBody() (string, error)
	OpenGraphDescription() (string, bool)
}

// Collection is a list of Content items
type Collection interface {
	Source() string
	Content() []Content
	Errors() []error
}

// Content is the typical set of fields defined for almost any generated or constructed content page
type Content interface {
	Title() Title
	Summary() Summary
	Body() string
	Categories() []string
	CreatedOn() time.Time
	FeaturedImage() *url.URL
	OpenGraphContent(ogKey string, defaultValue *string) (string, bool)
	TwitterCardContent(twitterKey string, defaultValue *string) (string, bool)
	Errors() []error
}

// IgnoreResourceRule is a rule
type IgnoreResourceRule interface {
	IgnoreResource(url *url.URL) (bool, string)
}

// CleanResourceParamsRule is a rule
type CleanResourceParamsRule interface {
	CleanResourceParams(url *url.URL) bool
	RemoveQueryParamFromResourceURL(paramName string) (bool, string)
}

// FollowRedirectsInCurationTargetHTMLPayload defines whether we follow redirect rules in HTML <meta> refresh tags
type FollowRedirectsInCurationTargetHTMLPayload bool

// CuratedContent is content which is basically a link to some other content on the Internet
type CuratedContent interface {
	Content
	TargetResource() *HarvestedResource
}
