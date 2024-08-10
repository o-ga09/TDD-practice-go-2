# Backend API in ogen

## テーマ

- マイクロサービスで作る
- ogenを使用する

## つくるもの

ノートアプリのバックエンド API

## MVP

- [ ] 機能を作成する
  - [ ] ノートを投稿できる
  - [x] ノートを閲覧できる
  - [ ] ノートを更新できる
  - [x] ノートを削除できる
  - [x] ユーザー登録できる
  - [x] ユーザー削除できる
  - [x] ユーザー情報を更新できる
  - [x] ユーザー情報を閲覧できる
- [ ] テストを作成する
  - [ ] E2Eテストを作成する
  - [ ] 単体テストを作成する
- [ ] 認証機能を作成する
  - [ ] specの追加
  - [ ] ユーザーを認可できる
  - [ ] ユーザーを認証できる

## 参考情報

- sqlcを使用したコード生成

```yaml
version: "2"
sql:
  - engine: "mysql"
    queries: "query.sql"
    schema: "schema.sql"
    gen:
      go:
        package: "db"
        out: "db"
```

```bash
# cqlc.ymlがあるディレクトリ
$ cd db
$ sqlc generate
```

- ogenを使用したコード生成

```go
// generate.go
package api

import _ "github.com/ogen-go/ogen"

//go:generate go run github.com/ogen-go/ogen/cmd/ogen --target . -package api --clean ../docs/openapi.yml
```

```bash
$ go generate ./...
```
