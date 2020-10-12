package hatena

import (
	"net/url"
	"regexp"
	"time"
)

var imgURLPattern = regexp.MustCompile(`https://.+?\.png`)

func (e *Entry) ExtractImageURLFromBody() {
	e.ImageURLs = parseImageURL(e.Content)
}

func parseImageURL(body string) []string {
	return imgURLPattern.FindAllString(body, -1)
}

type Entry struct {
	*EntryHeader
	LastModified *time.Time
	Content      string
	ContentType  string
	ImageURLs    []string
}

type EntryHeader struct {
	Title      string     `yaml:"Title"`
	Category   []string   `yaml:"Category,omitempty"`
	Date       *time.Time `yaml:"Date"`
	URL        *entryURL  `yaml:"URL"`
	EditURL    string     `yaml:"EditURL"`
	IsDraft    bool       `yaml:"Draft,omitempty"`
	CustomPath string     `yaml:"CustomPath,omitempty"`
}

type entryURL struct {
	*url.URL
}
