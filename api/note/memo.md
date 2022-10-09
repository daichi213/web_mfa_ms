# Golang memo

## Script mode

Gore というライブラリを使用することで、golang でも対話形式にてコマンドを実行することができる。その際、必要になったり役になった情報をこの節にまとめる。

### 変数の型確認

[ここのページ](https://y0m0r.hateblo.jp/entry/20140413/1397397593)を参考にした。reflect と呼ばれる標準パッケージを使用することで変数の型を確認することができる。

# DB 関連

## SQL 関連

### 既存テーブルの作成クエリ発行

以下コマンドを MySQL サーバー上で実行することで、そのテーブルの作成クエリを表示することができる。[このページ](https://qiita.com/expajp/items/81a8773b49472925fe06)を参考にした。

```sql
 SHOW CREATE TABLE todos;
```

## エラー

### 外部キーの作成エラー

#### エラー内容

sql-migration を使用して api より db へマイグレートを実行する際に、外部キーを作成するマイグレーションファイルで以下のエラーが発生した。

```bash
$ sql-migrate up --env="development"
Migration failed: Error 1215: Cannot add foreign key constraint handling 20210901133655-add_relation_between_user_and_todo.sql
```

#### 解決方法

db サーバーにログインして、以下コマンドを実行する。

```sql
SHOW ENGINE INNODB STATUS \G
...
------------------------
LATEST FOREIGN KEY ERROR
------------------------
2021-09-01 13:40:24 0x7fbaf421f700 Error in foreign key constraint of table todo_dev/#sql-1_15:
 FOREIGN KEY (User_id) REFERENCES user (ID):
Cannot find an index in the referenced table where the
referenced columns appear as the first columns, or column types
in the table and the referenced table do not match for constraint.
Note that the internal storage type of ENUM and SET changed in
tables created with >= InnoDB-4.1.12, and such columns in old tables
cannot be referenced by such columns in new tables.
Please refer to http://dev.mysql.com/doc/refman/5.7/en/innodb-foreign-key-constraints.html for correct foreign key definition.
...
```

#### 原因

user テーブルのプライマリーキーを以下のように serial 型で作成していたが、作成しようとした外部キーは INTEGER 型だったためこのようなエラーが発生した。

```sql
-- +migrate Up
CREATE TABLE IF NOT EXISTS user(
    id serial PRIMARY KEY ,
    Name VARCHAR (50) NOT NULL,
    Email VARCHAR (100) NOT NULL,
    Password VARCHAR (50) NOT NULL
);
```

## debug

golang のデバッグには delve というパッケージを使用してコードのデバッグを行う。その使用方法について記載する。

### インストール

```bash
$ go get github.com/go-delve/delve/cmd/dlv@latest
```

### 使用方法

#### 通常のデバッグモード

```bash
# デバッグ開始
$ dlv debug .
# ブレークポイントの設置(package名.関数名)
(dlv) break main.main
# nextコマンドのエイリアスでコードを一行一行実効する
(dlv) n
# continueコマンドのエイリアスでブレークポイントまで実効する
(dlv) c
```

#### テストコードのデバッグモード

```bash
# デバッグ開始
$ dlv test .
# ブレークポイントの設置(package名.関数名)
(dlv) break main.TestMain
# ブレークポイントの設置(ファイル名:行数)
(dlv) break main_test.go:16
# nextコマンドのエイリアスでコードを一行一行実効する
(dlv) n
# continueコマンドのエイリアスでブレークポイントまで実効する
(dlv) c
```

## bcrypt

### 概要

gin 自体にパスワードのハッシュ化のための機能は付随していないため、別の bcrypt というライブラリを使用することでパスワードをハッシュ化してテーブルに保存することができる。

### GenerateFromPassword での SALT の使用

bcrypt 内でパスワードをハッシュ化するための関数が GenerateFromPassword になっている。この関数は内部で salt を付与してハッシュ化してくれるため、自分自身でカスタムの salt などを付与する必要はない。

### 使用方法

```go
func InternalAuthenticatorFunction(c *gin.Context) (interface{}, error) {
	var loginVals Login
	if err := c.ShouldBind(&loginVals); err != nil {
		fmt.Println("checkpoint shouldbind")
		return "", jwt.ErrMissingLoginValues
	}

	if err := GetUserByEmail(loginVals.Email); err != nil {
		// log.Fatalf("No existing password is sent")
		fmt.Println("checkpoint get user")
		return "", jwt.ErrMissingLoginValues
	}

    // CompareHashAndPasswordの第一引数にsaltつきのハッシュ値、第二引数に素のパスワードのバイト列を指定する
	if invalid := bcrypt.CompareHashAndPassword(UserFromDB.Password, []byte(loginVals.Password)); invalid != nil {
		log.Fatalf("Password is wrong...;dbside:%v,loginVals:%v", UserFromDB.Password, []byte(loginVals.Password))
		return "", jwt.ErrFailedAuthentication
	} else {
		return &Login{
			UserName: 	UserToDB.UserName,
			Email: 		UserToDB.Email,
			Password:	loginVals.Password,
		}, nil
	}
}
```
