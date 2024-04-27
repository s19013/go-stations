# TechTrain Go Railway について

Railway では Git で自分が取り組んだ内容を記録するときに、自動でテストが実行されます。この際、Station の内容に即した実装になっているかを最低限のラインとして確認します。
テストが通れば Station クリアとなります。
クリア後、TechTrain の画面に戻り、クリアになっているかを確認してみてください。

[ユーザーマニュアル](https://docs.google.com/presentation/d/1BJSPCWBfy5xtwvBanoRGB77Y0N6JWMAHemEOLP49IPE/edit?usp=sharing)

## バージョン情報

|言語、フレームワークなど|バージョン|
|:---:|:---:|
Go| 1.16.* or higher
SQLite| 3.35.* or higher

## 初期設定

### 必要なツール

|ツール名|目安となるバージョン|
|:---:|:---:|
|Node.js| 14.*  [ 12.* ,  16.* では動作しません]|
|Yarn|1.22.*|

バージョンが異なる場合、動作しない場合があります。  
Node.js, Yarnのインストールがまだの場合は[html-stations](https://github.com/TechBowl-japan/html-stations)を参考にインストールしてください。  
また、使用PCがWindowsの場合は、WSLを[この記事](https://docs.microsoft.com/ja-jp/windows/wsl/install-win10)を参考にインストールしてください。

### 「必要なツール」インストール済みの場合

次の手順で取り組み始めてください。

####  `go-stations`リポジトリのFork

画面右上にあるForkより [Go Railway](https://github.com/TechBowl-japan/go-stations)のリポジトリを自分のアカウントにForkしてください。

#### `go-stations`リポジトリのクローン

作成したリポジトリを作業するディレクトリにクローンしましょう。

* Macなら Terminal.app(iTerm2などでも良い)
* Windowsなら PowerShell(GitBashなどのインストールしたアプリでも良いです。アプリによってはコマンドが異なることがあります)

で作業するディレクトリを開き、次のコマンドでForkしたGo Railwayのリポジトリをローカルにクローンしてください。


```powershell
git clone https://github.com/{GitHubのユーザー名}/go-stations.git
```

SSHでクローンを行う場合には、次のようになります

```
git clone git@github.com:[GitHubのユーザー名]/go-stations.git
```

#### Goのインストール

Windows編: https://golang.org/doc/install
Mac編: https://golang.org/doc/install

#### パッケージのインストール

クローンしたばかりのリポジトリは歯抜けの状態なので、必要なファイルをダウンロードする必要があります。
10 分程度掛かることもあるため、気長に待ちましょう。上から順番に**１つずつ**コマンドを実行しましょう：

```powershell
cd go-stations
```

```powershell
go mod download // ←データベースのドライバーとテスト用のライブラリをダウンロードします。
yarn install    // ←こちらを実行した後に「TechTrainにログインします。GitHubでサインアップした方はお手数ですが、パスワードリセットよりパスワードを発行してください」と出てくるため、ログインを実行してください。出てこない場合は、コマンドの実行に失敗している可能性があるため、TechTrainの問い合わせかRailwayのSlackより問い合わせをお願いいたします。
```

上記のコマンドを実行すると、techtrainにログインするように表示が行われます。
GitHubでサインアップしており、パスワードがない方がいましたら、そのかたはパスワードを再発行することでパスワードを作成してください。

ログインが完了すれば、ひとまず事前準備はおしまいです。お疲れ様でした。
TechTrainの画面からチャレンジを始めることもお忘れなく！
Go Railway に取り組み始めてください。

## DB(SQLite)と接続をしたいという方へ

* Sequel Pro
* Sequel Ace
* Table Plus
* VSCodeの拡張

などで確認することができます。

## トラブルシューティング

### go testで404というエラーが返ってきます。

main.goなどでhandlerの登録を確認してみましょう。
テストの関係上router.NewRouterのメソッド内部で追加するようにしましょう。

### DBに接続して中身が見れないのですが？

次のような結果が返ってきていれば、正常です。

```
$ sqlite3 .sqlite3/todo.db
SQLite version 3.32.3 2020-06-18 14:16:19
Enter ".help" for usage hints.
sqlite> .tables
todos
```

もし、 `todos` が作成されていないようであれば、次のコマンドを実行しましょう。

```
$ sqlite3 .sqlite3/todo.db < db/schema.sql
```

これで、 `todos` が作成されていれば、問題なく接続できます。

### commitしたのにチェックが実行されていないようなのですが？

チェックのためには、次の二つの条件が必須となります。

1. 黒い画面（CLI,コマンドライン）からTechTrainへのログイン
2. pre-commit hook と呼ばれるcommit時に実行されるGitの仕組みが仕込まれていること

特に2については

* SourceTreeやGitHubAppでクローンした
* httpsでクローンした

際にうまくいかないことが多いということが報告されています。
もし上記のようなことが起こった場合には、Terminalなどの画面でSSHによるクローンを試していただき、その上で `yarn install` を実行していただくことで解決することが多いです。もし解決しなかった場合には、運営までお問い合わせいただくか、RailwayのSlackワークスペースにてご質問ください。

## 自分のリポジトリの状態を最新の TechBowl-japan/go-stations と合わせる

Forkしたリポジトリは、Fork元のリポジトリの状態を自動的に反映してくれません。
Stationの問題やエラーの修正などがなされておらず、自分で更新をする必要があります。
何かエラーが出た、または運営から親リポジトリを更新してくださいと伝えられた際には、こちらを試してみてください。

### 準備

```shell
# こちらは、自分でクローンした[GitHubユーザー名]/go-stationsの作業ディレクトリを前提としてコマンドを用意しています。
# 自分が何か変更した内容があれば、 stash した後に実行してください。
git remote add upstream git@github.com:TechBowl-japan/go-stations.git
git fetch upstream
```

これらのコマンドを実行後にうまくいっていれば、次のような表示が含まれています。

```shell
git branch -a ←このコマンドを実行

* main
  remotes/origin/HEAD -> origin/main
  remotes/origin/main
  remotes/upstream/main ←こちらのような upstream という文字が含まれた表示の行があれば成功です。
```

こちらで自分のリポジトリを TechBowl-japan/go-stations の最新の状態と合わせるための準備は終了です。

### 自分のリポジトリの状態を最新に更新

```shell
# 自分の変更の状態を stash した上で次のコマンドを実行してください。

# ↓main ブランチに移動するコマンド
git checkout main

# ↓ TechBowl-japan/go-stations の最新の状態をオンラインから取得
git fetch upstream

# ↓ 最新の状態を自分のリポジトリに入れてローカルの状態も最新へ
git merge upstream/main
git push
yarn install
```

### GitHubアカウントでサインアップしたので、パスワードがないという方へ

https://techbowl.co.jp/techtrain/resetpassword

上記のURLより自分の登録したメールアドレスより、パスワードリセットを行うことで、パスワードを発行してください。

メールアドレスがわからない場合は、ログイン後にユーザー情報の編集画面で確認してください。
ログインしていれば、次のURLから確認できます。

https://techbowl.co.jp/techtrain/mypage/profile


# 参考にしたサイト
公式ドキュメント

# go
* [Go 言語の構造体 (struct)](https://golang.keicode.com/basics/go-struct.php)
* [とほほのGo言語入門](https://www.tohoho-web.com/ex/golang.html)
* [お気楽 Go 言語プログラミング入門](http://www.nct9.ne.jp/m_hiroi/golang/)
* [a tour of go](https://go-tour-jp.appspot.com/list)
* [Golang 型確認](https://qiita.com/mykysyk@github/items/08a7203d6013ecd74e9b)
* [GoのMarshal/Unmarshalの基本的な使い方とプライベートフィールドを持つ構造体での利用方法](https://www.asobou.co.jp/blog/web/marshal-unmarshal)
* [【Go言語】jsonデータをstreamで扱うEncoder、Decoder型を試してみる](https://www.asobou.co.jp/blog/web/encoder-decoder)
* [[Go] json.Unmarshal と json.Decoder の使い分け ～ json.Decoder に親しむ](https://zenn.dev/kariya_mitsuru/articles/921f162262ce24)
* [Goにおけるjsonの扱い方を整理・考察してみた ~ データスキーマを添えて](https://zenn.dev/hsaki/articles/go-convert-json-struct)
* [Golangのjson.Unmarshalとjson.Decoder.Decodeの違い](https://otameshi61.hatenablog.com/entry/2022/08/11/100322)
* [Goの基礎](https://golang.keicode.com/basics/)
* [【Golang】RequestからQueryStringを取得するには](https://blog.ryskit.com/entry/2018/07/08/175749)
* [Go言語で初めてWebアプリを作る際に役立ちそうな参考文献リンク集(日/英)](https://qiita.com/AYukiEngineer/items/bef36bd4752535e8d1f2)
* [Goの初心者が見ると幸せになれる場所](https://qiita.com/tenntenn/items/0e33a4959250d1a55045)
* [Awesome Go : 素晴らしい Go のフレームワーク・ライブラリ・ソフトウェアの数々](https://qiita.com/hatai/items/f31914f37dc6c53b2bce)
* [Go/標準ライブラリー](https://ja.wikibooks.org/wiki/Go/%E6%A8%99%E6%BA%96%E3%83%A9%E3%82%A4%E3%83%96%E3%83%A9%E3%83%AA%E3%83%BC)

## sqllite
* [Windows10にsqlite3とgccとgo-sqliteをインストールする【Go言語のオンライン学習日記】](https://www.ellenismos.com/go-sqlite3-gcc-install)
* [CHECK制約の使い方](https://www.javadrive.jp/sqlite/table/index13.html)
* [Go言語でSQLite3を使う](https://zenn.dev/teasy/articles/go-sqlite3-sample)

## net/http
* [公式ドキュメント](https://pkg.go.dev/net/http)
* [Goのhttp.Handlerやhttp.HandlerFuncをちゃんと理解する](https://journal.lampetty.net/entry/understanding-http-handler-in-go)
* [【Go】net/httpでGetしたjsonのResponseを見る](https://zenn.dev/yoji/articles/39d00b851af9d3)
* [Go における HTTP リクエストの受け取り方](https://qiita.com/BitterBamboo/items/182659dddc5b4b195976)
* [【要約版】net/httpでつくるHTTPルーター自作入門](https://qiita.com/bmf_san/items/312fac5b3132d8bee4ca)

## database/sql
* [公式ドキュメント](https://pkg.go.dev/database/sql)
* [GolangでAPIを作る時に見るレシピ集](https://zenn.dev/chillout2san/articles/34aa952747880c)
* [SQLの実行](https://www.twihike.dev/docs/golang-database/queries)
* [Go言語でMySQL の基本的操作（SELECT、UPDATE、INSERT）を行う](https://qiita.com/ShinyaIshikawa/items/fede44cee7c71721247a)
* [golang MariaDB IN句へ配列データの渡し方](https://blog.wsd.sh/?p=953)

## reflect
* [公式ドキュメント](https://pkg.go.dev/reflect)