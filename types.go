package content

import (
	"net/url"
	"time"
)

// Collection is a list of Content items
type Collection interface {
	Content() []*Content
}

// Content is the typical set of fields defined for almost any generated or constructed content page
type Content interface {
	Title() string
	Body() string
	Summary() string
	Categories() []string
	CreatedOn() time.Time
	FeaturedImage() url.URL
	Keys() *Keys
}

// CuratedLink is content which is basically a link to some other content on the Internet
type CuratedLink interface {
	Content
	Target() url.URL
}
