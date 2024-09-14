# マニュアルメモ

## Echoフレームワークのインストール

Echoフレームワークとミドルウェアをインストールします。

```bash
go get github.com/go-resty/resty/v2
go get github.com/labstack/echo/v4/middleware@v4.12.0
```

## Restyのインストール

Supabase APIと通信するために使用するHTTPクライアントライブラリRestyをインストールします。

```bash
go get github.com/go-resty/resty/v2
```

## 環境変数を読み込むためのパッケージ（オプション）

環境変数をファイルから読み込む場合、github.com/joho/godotenvを使用できます。

```bash
go get github.com/joho/godotenv
```

## Webアプリケーション起動

```bash
go run main.go
```

## Goのドキュメント生成ツール

Go言語では、標準ツールであるgodocを使ってソースコードからドキュメントを生成できます。godocは、ソースコード内のコメントからドキュメントを生成し、HTMLファイルとして出力するか、サーバーとして起動して閲覧できます。

```bash
sudo apt install golang-golang-x-tools
```

## テストコード導入

```bash
# パッケージインストール
go get github.com/stretchr/testify/assert
go mod tidy

# テスト実行
go test ./...
```

## UUID

```bash
go get github.com/google/uuid
```

## その他

```bash
go get github.com/go-playground/validator/v10
go get github.com/stretchr/testify/mock
```