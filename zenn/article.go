package zenn

import (
	"io"
	"path"
	"strings"
	"text/template"

	"github.com/ikawaha/zenn-importer/hatena"
	"github.com/ikawaha/zenn-importer/qiita"
)

var ZennTmpl = `---
title: "{{ replace .Title "\"" "\\\"" }}"
emoji: "{{ .Emoji }}"
type: "{{ .Type }}"
topics: [{{ join .Topics "," }}]
published: {{ .Published }}
published_at: {{ .PublishedAt }}
---
{{ .Body }}
`

type ZennArticle struct {
	Title       string
	Emoji       string
	Type        string
	Topics      []string
	Published   bool
	PublishedAt string
	Body        string
	ImageURLs   []string
	Slug        string
}

func NewZennArticleFromQiitaArticle(a *qiita.Article) *ZennArticle {
	slug := path.Base(a.URL)
	topics := make([]string, 0, len(a.Tags))
	for _, v := range a.Tags {
		topics = append(topics, v.Name)
	}
	return &ZennArticle{
		Title:       a.Title,
		Emoji:       "😀",
		Type:        "tech",
		Topics:      topics,
		Published:   false,
		PublishedAt: a.CreatedAt.Format("2006-01-02 15:04"),
		Body:        a.Body,
		ImageURLs:   a.ImageURLs,
		Slug:        slug,
	}
}

func NewZennArticleFromHatenaEntry(ent *hatena.Entry) *ZennArticle {
	slug := path.Base(ent.URL.String())
	topics := make([]string, 0, len(ent.Category))
	for _, v := range ent.Category {
		topics = append(topics, v)
	}
	return &ZennArticle{
		Title:       ent.Title,
		Emoji:       "😀",
		Type:        "tech",
		Topics:      topics,
		Published:   false,
		PublishedAt: ent.Date.Format("2006-01-02 15:04"),
		Body:        ent.Content,
		ImageURLs:   nil, // TODO
		Slug:        slug,
	}
}

var tmpl = template.Must(template.New("zenn").Funcs(template.FuncMap{
	"join":    strings.Join,
	"replace": strings.ReplaceAll,
}).Parse(ZennTmpl))

func (a ZennArticle) Write(w io.Writer) error {
	if rs := []rune(a.Title); len(rs) > 60 {
		a.Title = string([]rune(a.Title)[:60])
	}
	if a.Body == "" {
		a.Body = "<empty>"
	}
	return tmpl.Execute(w, a)
}
