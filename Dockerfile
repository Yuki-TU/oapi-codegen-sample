
# ----------------------------------------------
# ビルド用環境
# ----------------------------------------------
FROM golang:1.23-bullseye AS deploy-builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags "-w -s" -o app

# ----------------------------------------------
# 本番環境
# ----------------------------------------------
FROM debian:bullseye-slim AS deploy

# X509: Certificate Signed by Unknown Authorityエラーを回避する
RUN apt-get update \
 && apt-get install -y --force-yes --no-install-recommends apt-transport-https curl ca-certificates \
 && apt-get clean \
 && apt-get autoremove \
 && rm -rf /var/lib/apt/lists/*

COPY --from=deploy-builder /app/app .

CMD ["./app"]

# ----------------------------------------------
# 開発環境
# ----------------------------------------------
FROM golang:1.23-alpine AS dev

WORKDIR /app

RUN apk update && apk add alpine-sdk jq mysql mysql-client nodejs npm binutils-gold

RUN go install github.com/air-verse/air@latest \
  && go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest \
  && go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest \
  && go install github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@latest \
  && go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest \
  && go install github.com/rubenv/sql-migrate/...@latest \
  && go install github.com/google/wire/cmd/wire@latest \
  && go install github.com/k1LoW/tbls@latest \
  && go install go.uber.org/mock/mockgen@latest \
  npm i -g @redocly/cli@latest
