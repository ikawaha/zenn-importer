package qiita

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	verbose bool
}

func NewClient() *Client {
	return &Client{}
}

func (c *Client) SetVerbose(verbose bool) {
	c.verbose = verbose
}

func (c Client) FetchArticles(user string) ([]Article, error) {
	var ret []Article
	for i := 1; ; i++ {
		as, err := c.FetchArticlePage(user, i)
		if err != nil {
			return ret, fmt.Errorf("fetch articles failed: page=%d, %v", i, err)
		}
		if len(as) == 0 {
			break
		}
		ret = append(ret, as...)
	}
	return ret, nil
}

func (c Client) FetchArticlePage(user string, page int) ([]Article, error) {
	q := fmt.Sprintf("https://qiita.com/api/v2/users/%s/items?page=%d", url.QueryEscape(user), page)
	if c.verbose {
		fmt.Println(q)
	}
	resp, err := http.Get(q)
	if err != nil {
		return nil, err
	}
	defer func() {
		_, _ = io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}()
	var ret []Article
	err = json.NewDecoder(resp.Body).Decode(&ret)
	if c.verbose {
		for _, v := range ret {
			fmt.Println(v.Title)
		}
		fmt.Println()
	}
	return ret, err
}
