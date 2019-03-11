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