package hatena

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/motemen/go-wsse"
	"github.com/x-motemen/blogsync/atom"
)

type Client struct {
	*atom.Client
	config Config
}

func NewClient(conf Config) *Client {
	return &Client{
		Client: &atom.Client{
			Client: &http.Client{
				Transport: &wsse.Transport{
					Username: conf.Username,
					Password: conf.Password,
				},
			},
		},
		config: conf,
	}
}

func (b *Client) FetchRemoteEntries() ([]*Entry, error) {
	entries := []*Entry{}
	url := fmt.Sprintf(
		"https://blog.hatena.ne.jp/%s/%s/atom/entry",
		url.QueryEscape(b.config.Username),
		url.QueryEscape(b.config.RemoteRoot),
	)
	for {
		feed, err := b.Client.GetFeed(url)
		if err != nil {
			return nil, err
		}
		for _, ae := range feed.Entries {
			e, err := entryFromAtom(&ae)
			if err != nil {
				return nil, err
			}
			if b.config.Verbose {
				fmt.Println(e.Title)
			}
			entries = append(entries, e)
		}
		nextLink := feed.Links.Find("next")
		if nextLink == nil {
			break
		}
		url = nextLink.Href
	}
	return entries, nil
}

func entryFromAtom(e *atom.Entry) (*Entry, error) {
	alternateLink := e.Links.Find("alternate")
	if alternateLink == nil {
		return nil, fmt.Errorf("could not find link[rel=alternate]")
	}

	u, err := url.Parse(alternateLink.Href)
	if err != nil {
		return nil, err
	}

	editLink := e.Links.Find("edit")
	if editLink == nil {
		return nil, fmt.Errorf("could not find link[rel=edit]")
	}

	categories := make([]string, 0)
	for _, c := range e.Category {
		categories = append(categories, c.Term)
	}

	entry := &Entry{
		EntryHeader: &EntryHeader{
			URL:      &entryURL{URL: u},
			EditURL:  editLink.Href,
			Title:    e.Title,
			Category: categories,
			Date:     e.Updated,
		},
		LastModified: e.Edited,
		Content:      e.Content.Content,
		ContentType:  e.Content.Type,
	}

	if e.Control != nil && e.Control.Draft == "yes" {
		entry.IsDraft = true
	}

	return entry, nil
}
