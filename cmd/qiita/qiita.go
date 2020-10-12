package qiita

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/ikawaha/zenn-importer/qiita"
	"github.com/ikawaha/zenn-importer/zenn"
)

const (
	defaultOutputDir = "."
	defaultImageDir  = "img"
	defaultPrefix    = "qiita-"
)

type option struct {
	dir      string
	imageDir string
	user     string
	prefix   string
	publish  bool
	verbose  bool

	flagSet *flag.FlagSet
}

func newOption() *option {
	opt := &option{
		flagSet: flag.NewFlagSet("qiita-zenn", flag.ExitOnError),
	}
	opt.flagSet.StringVar(&opt.dir, "dir", defaultOutputDir, "output dir")
	opt.flagSet.StringVar(&opt.imageDir, "imgdir", defaultImageDir, "image dir")
	opt.flagSet.StringVar(&opt.user, "user", "", "user name of qiita")
	opt.flagSet.StringVar(&opt.prefix, "prefix", defaultPrefix, "prefix of articles")
	opt.flagSet.BoolVar(&opt.publish, "publish", false, "published option of zenn is set true")
	opt.flagSet.BoolVar(&opt.verbose, "verbose", false, "verbose")
	return opt
}

func (o *option) parse(args []string) error {
	if err := o.flagSet.Parse(args); err != nil {
		return err
	}
	if o.user == "" {
		return fmt.Errorf("invalid empty argument: -user %q", o.user)
	}
	return nil
}

func Cmd(args []string) error {
	opt := newOption()
	if err := opt.parse(args); err != nil {
		opt.flagSet.PrintDefaults()
		return err
	}
	return run(opt)
}

func run(opt *option) error {
	cl := qiita.NewClient()
	cl.SetVerbose(opt.verbose)
	for i := 1; ; i++ {
		as, err := cl.FetchArticlePage(opt.user, i)
		if err != nil {
			fmt.Errorf("fetch articles error: page=%d, %v", i, err)
		}
		if len(as) == 0 {
			break
		}
		for i := range as {
			if err := saveArticle(opt.dir, opt.prefix, opt.publish, &as[i]); err != nil {
				return err
			}
			as[i].ExtractImageURLFromBody()
			for _, v := range as[i].ImageURLs {
				if err := saveImage(opt.imageDir, v); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func downloadImage(w io.Writer, url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	_, err = io.Copy(w, resp.Body)
	return err
}

func mkdirp(dir string) error {
	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		return os.Mkdir(dir, 0777)
	}
	return err
}

func saveImage(dir, url string) error {
	name := path.Base(url)
	if err := mkdirp(dir); err != nil {
		return err
	}
	f, err := os.Create(filepath.Join(dir, name))
	if err != nil {
		return err
	}
	defer f.Close()
	return downloadImage(f, url)
}

func saveArticle(dir, prefix string, publish bool, a *qiita.Article) error {
	date := a.CreatedAt.Format("20060102-")
	slug := prefix + date + path.Base(a.URL) + ".md"
	z := zenn.NewZennArticleFromQiitaArticle(a)
	z.Published = publish
	f, err := os.Create(filepath.Join(dir, slug))
	if err != nil {
		return err
	}
	defer f.Close()
	if err := z.Write(f); err != nil {
		return err
	}
	return nil
}
