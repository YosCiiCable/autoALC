# autoALC
ALCを自動でしてくれるツールを作りたい人生だった…。

# 実行
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

# ディレクトリの説明
## archive/
アーカイブ化したプロジェクト。
## bml/
ブックマークレットをつかうプロジェクト。
## go/
go言語もメインプロジェクト。
