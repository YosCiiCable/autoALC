# autoALC
ALCを自動でしてくれるツールを作りたい人生だった…。

# エラー
~~~
2023/03/07 18:59:10 err0: English
~~~
というような表示で終了した場合は、エラーです。
入力ミスの場合もありますし、実行上のエラーの場合もあります。
再度最初から行っても同じエラーの場合は、開発者に連絡ください。

# Docker を使う
## 起動
go/ で以下のコードを分けて実行
~~~
docker-compose up -d
docker cp gocon:/usr/src/app/fullScreenshot.png ./
~~~
## 停止と削除
go/ で以下のコードを実行
~~~
docker compose down --rmi all --volumes --remove-orphans
~~~

# build する
## Windows (64bit amd)
~~~
GOOS=windows GOARCH=amd64 go build -o main.exe main.go
~~~
## GOARCHに利用可能なもの
- 386
- amd64
- arm

## Linux (64bit intel)
~~~
GOOS=linux GOARCH=arm64 go build -o main.exe main.go
~~~
## GOARCHに利用可能なもの
- 386
- amd64
- arm
- arm64
- mips
- mips64
- mips64le
- mipsle
- ppc64
- ppc64le
- riscv64
- s390x

# ディレクトリの説明
## archive/
アーカイブ化したプロジェクト。
## bml/
ブックマークレットをつかうプロジェクト。
## go/
go言語もメインプロジェクト。
