package content

import (
	"net/url"
	"regexp"
	"time"
)

// Title is the content title which can be retrieved in different ways
type Title string

var sourceNameAsSuffixRegEx = regexp.MustCompile(` \| .*$`) // Removes " | Healthcare IT News" from a title like "xyz title | Healthcare IT News"

// Original is the title's original text
func (t Title) Original() string {
	return string(t)
}

// Clean is the title's "cleaned up text" (which removes "| ..."" suffixes)
func (t Title) Clean() string {
	return sourceNameAsSuffixRegEx.ReplaceAllString(string(t), "")
}

// Collection is a list of Content items
type Collection interface {
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
	Body() string
	Summary() string
	Categories() []string
	CreatedOn() time.Time
	FeaturedImage() url.URL
	Keys() Keys
}

// CuratedLink is content which is basically a link to some other content on the Internet
type CuratedLink interface {
	Content
	Target() url.URL
}
