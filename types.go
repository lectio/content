package content

import (
	"fmt"
	"net/url"
	"time"

	"github.com/lectio/harvester"
	"github.com/lectio/score"
)

// Instant is used for publication date/time
type Instant time.Time

// MarshalJSON creates the custom date/time format used by ContentDateTime
func (i Instant) MarshalJSON() ([]byte, error) {
	stamp := fmt.Sprintf("\"%s\"", time.Time(i).Format("Mon Jan 2 15:04:05 MST 2006"))
	return []byte(stamp), nil
}

// Collection is a list of Content items
type Collection interface {
	GetContent() []*Content
}

// Content is the typical set of fields defined for almost any generated or constructed content page
type Content interface {
	GetTitle() string
	GetBody() string
	GetSummary() string
	GetCategories() []string
	GetPublishedDate() Instant
	GetFeaturedImage() url.URL
	GetSlug() string
}

// CuratedLink is content which is basically a link to some other content on the Internet
type CuratedLink interface {
	Content
	GetResource() *harvester.HarvestedResource
	GetSource() string
}

// Scores are the social or other score types associated with some content
type Scores interface {
	GetTotalSharesCount() int
	GetFacebookGraph() *score.FacebookGraphResult
	GetLinkedInCount() *score.LinkedInCountServResult
}
