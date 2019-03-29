package content

import (
	"net/url"
	"time"
)

// Title is the content title which can be retrieved in different ways
type Title interface {
	Original() string
	Clean() string
}

// Summary is the content's description or summary which can be retrieved in different ways
type Summary interface {
	Original() string
	FirstSentenceOfBody(defaultValue *string) (string, error)
	OpenGraphDescription(ogKey string, defaultValue *string) (string, bool)
}

// Collection is a list of Content items
type Collection interface {
	Source() string
	Content() []Content
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
}

// CuratedLink is content which is basically a link to some other content on the Internet
type CuratedLink interface {
	Content
	Target() *url.URL
}
