# ベースイメージとしてGoの公式イメージを使用
FROM golang:1.23-alpine

# ホットリロードツール air をインストール
RUN go install github.com/air-verse/air@latest

# 作業ディレクトリを設定
WORKDIR /app

# go.mod と go.sum をコンテナにコピー
COPY go.mod go.sum ./

# 依存関係をダウンロード
RUN go mod download

# アプリケーションのソースコードをすべてコンテナにコピー
COPY . .

# ホットリロード用の air を使用して起動
CMD ["air"]
