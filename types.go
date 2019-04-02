package content

import (
	"net/url"
	"time"
)

// Keys defines different ways the content can be indexed
type Keys interface {
	GloballyUniqueKey() string
}

// Title is the content title which can be retrieved in different ways
type Title interface {
	Original() string
	Clean() string
	OpenGraphTitle() (string, bool)
}

// Summary is the content's description or summary which can be retrieved in different ways
type Summary interface {
	Original() string
	OpenGraphDescription() (string, bool)
}

// Body is the content's body which can be retrieved in different ways
type Body interface {
	Original() string
	FirstSentence() (string, error)
	WithoutFrontMatter() string
	HaveFrontMatter() bool
	FrontMatter() (interface{}, error)
	FrontMatterValue(key interface{}) (interface{}, bool, error)
}

// Collection is a list of Content items
type Collection interface {
	Source() string
	Content() []Content
	Errors() []error
}

// Content is the typical set of fields defined for almost any generated or constructed content page
type Content interface {
	Keys() Keys
	Title() Title
	Summary() Summary
	Body() Body
	Categories() []string
	CreatedOn() time.Time
	FeaturedImage() *url.URL
	OpenGraphContent(ogKey string, defaultValue *string) (string, bool)
	TwitterCardContent(twitterKey string, defaultValue *string) (string, bool)
	Errors() []error
	Directives() (interface{}, error)
	Directive(key interface{}) (interface{}, bool, error)
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
