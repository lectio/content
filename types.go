package content

import (
	"net/url"
	"time"

	"github.com/lectio/flexmap"
)

// Keys defines different ways the content can be indexed
type Keys interface {
	GloballyUniqueKey() string
}

// Title is the content title which can be retrieved in different ways
type Title interface {
	Original() string
	Clean() string
	OpenGraphTitle(clean bool) (string, bool)
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
	HasFrontMatter() bool
	FrontMatter() flexmap.TextKeyMap
}

// Collection is a list of Content items
type Collection interface {
	Source() string
	Content() []Content
	FilteredContent(keep func(c Content) bool) []Content
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
	Directives() flexmap.Map
}

// Link is an external URL
type Link interface {
	GloballyUniqueKey() string
	IsValid() (bool, bool)
	IsIgnored() (bool, string)
	FinalURL() (*url.URL, error)
}

// CuratedContent is content which is basically a link to some other content on the Internet
type CuratedContent interface {
	Content
	Link() Link
}
