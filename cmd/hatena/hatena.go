package hatena

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"path/filepath"

	"github.com/ikawaha/zenn-importer/hatena"
	"github.com/ikawaha/zenn-importer/zenn"
)

const (
	defaultOutputDir = "."
	defaultImageDir  = "img"
	defaultPrefix    = "hatena-"
)

type option struct {
	dir      string
	imageDir string
	blog     string
	user     string
	apikey   string
	prefix   string
	publish  bool
	verbose  bool

	flagSet *flag.FlagSet
}

func newOption() *option {
	opt := &option{
		flagSet: flag.NewFlagSet("hatena-zenn", flag.ExitOnError),
	}
	opt.flagSet.StringVar(&opt.dir, "dir", defaultOutputDir, "output dir")
	opt.flagSet.StringVar(&opt.imageDir, "imgdir", defaultImageDir, "image dir")
	opt.flagSet.StringVar(&opt.blog, "blog", "", "blog root, eg. ikawaha.hateblo.jp")
	opt.flagSet.StringVar(&opt.user, "user", "", "user name of hatena, eg. ikawaha")
	opt.flagSet.StringVar(&opt.apikey, "apikey", "", "Hatena blog API key (AtomPub API key), see. http://blog.hatena.ne.jp/my/config/detail")
	opt.flagSet.StringVar(&opt.prefix, "prefix", defaultPrefix, "prefix of articles")
	opt.flagSet.BoolVar(&opt.publish, "publish", false, "published option of zenn is set true")
	opt.flagSet.BoolVar(&opt.verbose, "verbose", false, "verbose")
	return opt
}

func (o *option) parse(args []string) error {
	if err := o.flagSet.Parse(args); err != nil {
		return err
	}
	switch {
	case o.blog == "":
		return fmt.Errorf("invalid empty argument: -blog %q", o.user)
	case o.user == "":
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
	config := hatena.Config{
		RemoteRoot: opt.blog,
		Username:   opt.user,
		Password:   opt.apikey,
		Verbose:    opt.verbose,
	}
	cl := hatena.NewClient(config)
	ents, err := cl.FetchRemoteEntries()
	if err != nil {
		fmt.Errorf("fetch articles error: %v", err)
	}
	if len(ents) == 0 {
		return nil
	}
	for i := range ents {
		if err := saveArticle(opt.dir, opt.prefix, opt.publish, ents[i]); err != nil {
			return err
		}
		//ents[i].ExtractImageURLFromBody() // TODO
		//for _, v := range ents[i].ImageURLs {
		//	if err := saveImage(opt.imageDir, v); err != nil {
		//		return err
		//	}
		//}
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

func saveArticle(dir, prefix string, publish bool, ent *hatena.Entry) error {
	date := ent.Date.Format("20060102-")
	slug := prefix + date + path.Base(ent.URL.String()) + ".md"
	z := zenn.NewZennArticleFromHatenaEntry(ent)
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
