# TodoApp

## 概要

Golangの標準ライブラリ等で作ってみたWebサーバー

Nginx*FCGI構成で、Goアプリケーション内ではhttprouterでルーティング定義。
HTMLを返却。

DBはGormで構造体を定義したものをそのまま自動マイグレーション
ディレクトリ構造はMVCアーキテクチャ風....??

## 利用ライブラリ

利用した主要なライブラリ達

- html/template
- net/http
- github.com/jinzhu/gorm
- github.com/julienschmidt/httprouter

## 使い方

1. 任意のDBにテーブルを作成
2. DBとDB作成時のパス等に合わせてmodel/init.goのパッケージとgorm.Openの引数を変える
3. Nginxをいい感じに立ち上げていい感じに設定してポートをいい感じにしてこのアプリケーションと接続する。別にUnixドメインソケットで接続でしてもいいし、TCPで接続してもいいし、なんならFCGI設定外してNginx使わなくてもいい。
4. ポートによるけどいい感じにアプリが起動して使えるようになる。