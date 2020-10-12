package qiita

import (
	"regexp"
	"time"
)

var imgURLPattern = regexp.MustCompile(`https://.+?\.png`)

type Tag struct {
	Name     string   `json:"name"`
	Versions []string `json:"versions"`
}

type Article struct {
	RenderedBody   string    `json:"rendered_body"`
	Body           string    `json:"body"`
	Coediting      bool      `json:"coediting"`
	CommentCount   int       `json:"comments_count"`
	CreatedAt      time.Time `json:"created_at"`
	Group          *Group    `json:"group"`
	ID             string    `json:"id"`
	LikesCount     int       `json:"likes_count"`
	Private        bool      `json:"private"`
	ReactionsCount int       `json:"reactions_count"`
	Tags           []Tag     `json:"tags"`
	Title          string    `json:"title"`
	UpdatedAt      time.Time `json:"updated_at"`
	URL            string    `json:"url"`
	User           User      `json:"user"`
	PageViewCount  *int      `json:"page_view_count"`
	ImageURLs      []string  `json:"-"`
}

func (a *Article) ExtractImageURLFromBody() {
	a.ImageURLs = parseImageURL(a.Body)
}

func parseImageURL(body string) []string {
	return imgURLPattern.FindAllString(body, -1)
}
