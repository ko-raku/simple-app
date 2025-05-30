# 簡易テキスト解析ウェブアプリケーション
## 概要
このアプリケーションは、Go言語で作成された簡易的なウェブアプリケーションです。テキストデータの保存、表示、編集、そして日本語テキストの形態素解析による検索機能を提供します。
## 主な機能
- テキストデータの作成と保存
- 保存されたテキストの表示と編集
- 日本語テキストの形態素解析による検索
- 検索クエリの出現回数カウント（形態素解析による方法と直接テキスト検索による方法）

## 技術スタック
- Go 1.24.2
- 標準ライブラリ（net/http, text/template など）
- カスタムライブラリ（mylib）- 形態素解析機能を提供

## インストールと実行方法
### 前提条件
- Go 1.24.2 以上がインストールされていること

### インストール手順
1. リポジトリをクローン
``` 
   git clone https://github.com/ko-raku/simple-app.git
   cd simpleApp
```
1. 依存関係のインストール
``` 
   go mod tidy
```
1. アプリケーションの実行
``` 
   go run main.go
```
1. ブラウザで以下のURLにアクセス
``` 
   http://localhost:8080/view/test
```
## 使用方法
### テキストの表示
- `/view/[ページ名]` にアクセスすることで、保存されたテキストを表示できます。

### テキストの編集
- `/edit/[ページ名]` にアクセスするか、表示ページの「Edit」リンクをクリックすることで、テキストを編集できます。

### テキストの検索
1. テキスト表示ページにある検索フォームに検索したい語句を入力
2. 「検索」ボタンをクリックすると、形態素解析と直接検索の2つの方法で検索クエリの出現回数が表示されます

## ファイル構成
- - メインアプリケーションコード `main.go`
- - テキスト表示のためのテンプレート `view.html`
- - テキスト編集のためのテンプレート `edit.html`
- - 検索結果表示のためのテンプレート `search.html`
- `*.txt` - テキストデータの保存ファイル

## 形態素解析機能
このアプリケーションでは `mylib` パッケージを使用して、日本語テキストの形態素解析を行っています。主に以下の関数を利用しています：
- `TokenizeText` - テキストを形態素解析してトークンに分割
- `CountPhraseFrequency` - 形態素列からフレーズの出現回数をカウント
- `CountPhraseInOriginalText` - 元のテキストから直接フレーズの出現回数をカウント

## URL構造
- `/view/[ページ名]` - テキストの表示
- `/edit/[ページ名]` - テキストの編集
- `/save/[ページ名]` - 編集したテキストの保存
- `/search/[ページ名]?query=[検索クエリ]` - テキストの検索

## 注意点
- このアプリケーションは教育・学習目的で作成されており、本番環境での使用は想定していません
- エラー処理やセキュリティ対策は最小限の実装となっています
- テキストデータはサーバーのローカルファイルシステムに保存されます

## ライセンス
このプロジェクトは [MIT ライセンス](LICENSE) のもとで公開されています。
