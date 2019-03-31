package content

import (
	"net/url"
	"time"
)

// ErrorsCollection tracks a collection of errors encountered
type ErrorsCollection interface {
	Errors() []error
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
	FirstSentenceOfBody() (string, error)
	OpenGraphDescription() (string, bool)
}

// Collection is a list of Content items
type Collection interface {
	Source() string
	Content() []Content
	Errors() ErrorsCollection
}

// Keys provides different ways of identifying content
type Keys interface {
	Content() Content
	UniqueID() uint32
	UniqueIDText(format string) string
	Slug() string
}

// Content is the typical set of fields defined for almost any generated or constructed content page
type Content interface {
	Title() Title
	Summary() Summary
	Body() string
	Categories() []string
	CreatedOn() time.Time
	FeaturedImage() *url.URL
	Keys() Keys
	OpenGraphContent(ogKey string, defaultValue *string) (string, bool)
	TwitterCardContent(twitterKey string, defaultValue *string) (string, bool)
	Errors() ErrorsCollection
}

// IgnoreCurationTargetRule is a rule
type IgnoreCurationTargetRule interface {
	IgnoreCurationTarget(url *url.URL) (bool, string)
}

// CleanCurationTargetRule is a rule
type CleanCurationTargetRule interface {
	CleanCurationTarget(url *url.URL) bool
	RemoveQueryParamFromCurationTargetURL(paramName string) (bool, string)
}

// FollowRedirectsInCurationTargetHTMLPayload defines whether we follow redirect rules in HTML <meta> refresh tags
type FollowRedirectsInCurationTargetHTMLPayload bool

// CuratedContent is content which is basically a link to some other content on the Internet
type CuratedContent interface {
	Content
	Target() *HarvestedResource
}
