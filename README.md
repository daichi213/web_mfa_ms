# 環境構築

## Build 手順まとめ

```sh
$ docker-compose run --rm front npx create-next-app --ts next_app
$ docker-compose build
$ docker-compose exec front sh
$ tsc --init
```

## front について

### 初期設定

1. front ディレクトリを作成後に front 用の Dockerfile を配置

2. 以下コマンドを順番に実行する

```sh
$ docker-compose run --rm front npx create-next-app todo_app
$ docker-compose build
# front
$ docker-compose run front sh
$ tsc --init
$ npm install --save-dev jest jest-dom @types/jest ts-jest @testing-library/dom @testing-library/jest-dom @testing-library/react babel-jest identity-obj-proxy react-test-renderer
# api
```

3. tsconfig.json の設定を以下とする
   {
   "compilerOptions": {
   "jsx": "react-jsx",
   }
   }

### jest, React Testing Library の導入

jest と RTL の導入に関しては[Nextjs の公式ページ](https://nextjs.org/docs/testing)を参考にした。

1. nextjs のプロジェクトルートディレクトリに jest.config.js と jest.setup.js を配置する。

```js
// jest.config.js

module.exports = {
  setupFilesAfterEnv: ["<rootDir>/jest.setup.js"],
  collectCoverageFrom: [
    "**/*.{js,jsx,ts,tsx}",
    "!**/*.d.ts",
    "!**/node_modules/**",
  ],
  moduleNameMapper: {
    /* Handle CSS imports (with CSS modules)
      https://jestjs.io/docs/webpack#mocking-css-modules */
    "^.+\\.module\\.(css|sass|scss)$": "identity-obj-proxy",

    // Handle CSS imports (without CSS modules)
    "^.+\\.(css|sass|scss)$": "<rootDir>/__mocks__/styleMock.js",

    /* Handle image imports
      https://jestjs.io/docs/webpack#handling-static-assets */
    "^.+\\.(jpg|jpeg|png|gif|webp|svg)$": "<rootDir>/__mocks__/fileMock.js",
  },
  testPathIgnorePatterns: ["<rootDir>/node_modules/", "<rootDir>/.next/"],
  testEnvironment: "jsdom",
  transform: {
    /* Use babel-jest to transpile tests with the next/babel preset
      https://jestjs.io/docs/configuration#transform-objectstring-pathtotransformer--pathtotransformer-object */
    "^.+\\.(js|jsx|ts|tsx)$": ["babel-jest", { presets: ["next/babel"] }],
  },
  transformIgnorePatterns: [
    "/node_modules/",
    "^.+\\.module\\.(css|sass|scss)$",
  ],
};
```

以下のファイルにテストを走らせる際に毎回読み込ませたいライブラリを記載することで、各テストスィートの実行前に毎回準備してくれる。

```js
// jest.setup.js

import "@testing-library/jest-dom/extend-expect";
```

2. **mocks**, **tests**ディレクトリ node_modules と同階層のディレクトリに作成する

3. **mocks**ディレクトリに以下ファイルと設定を追加する

```js
// __mocks__/fileMock.js

module.exports = "test-file-stub";
```

```js
// __mocks__/styleMock.js

module.exports = {};
```

4. 以下コマンドを実行する

```bash
$ npm install --save-dev jest jest-dom @types/jest ts-jest @testing-library/dom @testing-library/jest-dom @testing-library/react babel-jest identity-obj-proxy react-test-renderer
```

5. package.json に test コマンドを追加する

```json
"scripts": {
  "dev": "next dev",
  "build": "next build",
  "start": "next start",
  "test": "jest --watchAll"
}
```

### Material UI の導入

基本的に[こちらのサイト](https://maku.blog/p/s6djqw3/)を参考にした。

1. コンテナ内で以下コマンドを実行する

```bash
$ npm install @material-ui/core @material-ui/icons
```

2. Nextjs で MaterialUI を使用する際は、SSR との兼ね合いからスタイルの処理順序を制御する必要があるとのこと。そのため、以下の設定ファイルを準備する。

```

```

### 本番環境で必要な npm モジュール

本番環境でも必要になるモジュールを以下コマンドでインストールする。node_modules が存在するディレクトリで npm install して導入する。

```bash
$ cd todo_app
$ npm install axios
```

### ライブラリを追加する場合

- Dcoerfile に追加するライブラリを記載して build する

#### 注意点

## api(gin)について

### 初期設定

1. api ディレクトリを作成後に front 用の Dockerfile を配置

2. golang のベースイメージを元にコンテナを起動し、作業ディレクトリにて以下コマンドを実行して go.mod ファイルを作成する。golang の web フレームワークである gin を使用する場合は以下 url を追加する

```sh
# コンテナの起動
$ docker-compose run api sh
# modファイルの作成
$ go mod init api
# modファイルへ必要なライブラリの追加
$ go get -u github.com/gin-gonic/gin
$ touch main.go
$ go run main.go
```

以下コマンドをコンテナーへ入った後にまとめて貼り付けてまとめて実行する

```sh
go get golang.org/x/tools/cmd/godoc
go get github.com/lib/pq
go get gorm.io/driver/postgres
go get github.com/rubenv/sql-migrate/...
go get github.com/gin-gonic/gin
go get gorm.io/gorm
go get github.com/go-delve/delve/cmd/dlv@latest
go get github.com/stretchr/testify
go get github.com/DATA-DOG/go-sqlmock
go get github.com/google/wire
go get github.com/gin-contrib/sessions
go get github.com/koron/go-dproxy
go install github.com/x-motemen/gore/cmd/gore@latest
```

3. 作成された go.mod, go.sum をイメージ内へ転写して `go mod download`を実行することでイメージにライブラリをインストールすることができる。

### 環境構築手順

```sh
$ docker-compose exec api sh
# Table作成
$ sql-migrate up -env="development"
# Tableが作成されているか確認する
$ sql-migrate status -env="development"
# saltを生成するためのシェル
$ sh /go/src/api/generatingSalt.sh
# saltが生成されてるか確認する
$ cat /go/src/api/salt.env
SALT="ivYD6S8xRZ0SfacCmhmcfROlBKz7VyjWmIUV"/go/src/api #
```

## 備考

### chache の削除

mysql のイメージを使用時に compose ファイルで environment を設定した時、一度コンテナを立てると chache に設定が保存されている。そのため、このような chache に設定が保存されている場合は chache を削除する必要がある。chache の削除は以下コマンドで実行することができる。

```bash
$ docker builder prune
```

## 既知の脆弱性への対策

### 認証系

#### ブルートフォース攻撃

#### レインボーテーブル攻撃
