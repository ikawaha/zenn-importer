Zenn Importer
===

Download articles of Qiita/Hatena in Zenn format.

| | Media | URL |
|---|---|---|
|From|Qiita| https://qiita.com/ikawaha/items/ab9906581e34f26993a9|
|To|Zenn| https://zenn.dev/ikawaha/articles/qiita-20201009-ab9906581e34f26993a9|

| | Media | URL |
|---|---|---|
|From|Hatena|https://ikawaha.hateblo.jp/entry/2020/10/12/205602|
|To|Zenn| https://zenn.dev/ikawaha/articles/hatena-20201012-205602|


# Install

```shellsession
GO111MODULE=on go get -u github.com/ikawaha/zenn-importer
```

# Usage

```
$ zenn-importer
zenn-importer: a command to save blog posts in zenn format.
subcommands:
  - qiita
  - hatena
```

eg. 

```shellsession
$ zenn-importer qitta -user ikawaha -publish -verbose
https://qiita.com/api/v2/users/ikawaha/items?page=1
Goa v3 のテストをシュッとする
人生で何度目かのダブル配列TRIEを書いた
形態素解析器 kagome のユーザー辞書の使い方
goa でデザイン・ファーストをシュッとする
golang で string のポインタを取得する
Luceneで使われてるFSTを実装してみた（正規表現マッチ：VMアプローチへの招待）
メモ：golang で []byte と string の読み込みを透過的に扱う試行錯誤
Elasticsearch CheatSheet
Pure Go な形態素解析器で実行バイナリに辞書埋め込んだヤツを作ってみた(3) 完結編
json を pretty print するのに echo '{"apple": "red", "lemon": "yellow"}' | python -m json.tool は冗長じゃないですか？なので go でコマンド用意しました
【kagome】形態素解析のラティスをグラフで表示できるようにしてみた
golang のリファクタリングには gofmt が使える
golang で string を []byte にキャストするとメモリコピーが走ります
Elasticsearch で 文字の正規化を icu_normalizer でおこなう
goをソースからコンパイルしてgodocが見つからない
デフォルトのElasticsearchは常にクラスタを組みたがっているので気をつける，というかクラスタ設定とめとく
pythonクライアントで始める「はじめてのElasticsearch」
Travis CI で goxc を使うとなぜかgo1.3のときだけビルドがこける
travis setup releases で Validation Failed になるときの解消法
全裸改善：形態素解析器 kagome に go-bindata を使ったらビルド時間とバイナリのサイズが劇的に改善，『全裸で形態素解析スクリプト』もビルドでこけなくなるはず！

https://qiita.com/api/v2/users/ikawaha/items?page=2
go-bindata でコンパイル時にリソースを埋め込んじゃおう！
Pure Go で辞書同梱な形態素解析器 kagome を公開してみました
Go で利用できるプロファイリングツール pprof の読み方
Go でフォルダに含まれる特定の拡張子のファイルだけ列挙するには path/filepath を使う
Pure Go な形態素解析器で実行バイナリに辞書埋め込んだヤツを作ってみた (2) 未知語処理編
Pure Go な形態素解析器で実行バイナリに辞書埋め込んだヤツを作ってみた (1)
Go で euc-jp や sjis の csv ファイルを読み込むには変換用のリーダーを1つかませるだけでよかった
Go で DoubleArray を実装して Common Prefix Search する
Go でスライスの拡張をレシーバでおこなう（スライスのレシーバを作ろうとしてはまったこと ）
Go でファイルを1行ずつ読み込む（csv ファイルも）
Go でグラフを plot するパッケージを試した
Go で機械学習(の手習い)

https://qiita.com/api/v2/users/ikawaha/items?page=3

$ ls
img/					qiita-20140808-f43a78f171836f1ec86b.md
qiita-20140521-b57ce0d934f080cb8f72.md	qiita-20140821-ea1ca5ee90296307c4bc.md
qiita-20140522-61dac34ebfe2f8fa5c61.md	qiita-20140829-79fdd69c524310db065e.md
qiita-20140527-28186d965780fab5533d.md	qiita-20140903-3c3994559dfeffb9f8c9.md
qiita-20140605-1235a6c93c969591f517.md	qiita-20140904-73c0a1d975680276b2c7.md
qiita-20140605-f9e8b6ad2cae1eb593f7.md	qiita-20141011-6fd4071187640f6bf66d.md
qiita-20140612-540c2af34b1c381c37c1.md	qiita-20141012-8a01c5739401e26e8794.md
qiita-20140620-235371a7b1066ba5ea6c.md	qiita-20141020-5d21f16218bb13eeb286.md
qiita-20140626-8a94b728105e03a76468.md	qiita-20141021-228ee3f481e9636b3065.md
qiita-20140627-fdcf3ccf13f23a860390.md	qiita-20141218-2d58f58e4ab12918e8c9.md
qiita-20140704-e3b35f09fb49e9217924.md	qiita-20141219-be95304a803020e1b2d1.md
qiita-20140707-ff27ac03e22b7f36811b.md	qiita-20150305-adb24788f818dc9835f9.md
qiita-20140710-c02d84cfd00f8f442500.md	qiita-20161219-6638ee8b6978aef50d65.md
qiita-20140712-8c2b3bd56ff427e9d2d0.md	qiita-20171218-9ebe3e1104fb80706d99.md
qiita-20140717-c4859421d67a27e9ffc6.md	qiita-20181215-edb4e18960ae6e4babc3.md
qiita-20140717-f0a046428f907bf9ba98.md	qiita-20191202-e0c2b3ed0fedb12f4847.md
qiita-20140730-c654f746cfe76b888a27.md	qiita-20201009-ab9906581e34f26993a9.md
```

## Options

**Qiita**

```shellsession
$ zenn-importer qiita -help
Usage of qiita-zenn:
  -dir string
    	output dir (default ".")
  -imgdir string
    	image dir (default "img")
  -prefix string
    	prefix of articles (default "qiita-")
  -publish
    	published option of zenn is set true
  -user string
    	user name of qiita
  -verbose
    	verbose
```

**Hatena**

```
$ go run main.go hatena
  -apikey string
    	Hatena blog API key (AtomPub API key), see. http://blog.hatena.ne.jp/my/config/detail
  -blog string
    	blog root, eg. ikawaha.hateblo.jp
  -dir string
    	output dir (default ".")
  -imgdir string
    	image dir (default "img")
  -prefix string
    	prefix of articles (default "hatena-")
  -publish
    	published option of zenn is set true
  -user string
    	user name of hatena, eg. ikawaha
  -verbose
    	verbose
```

