# autoALC
ALCを自動でしてくれるツールを作りたい人生だった…。

# 実行関係
## エラー
~~~
2023/03/07 18:59:10 err0: English
~~~
というような表示で終了した場合は、エラーです。
入力ミスの場合もありますし、実行上のエラーの場合もあります。
再度最初から行っても同じエラーの場合は、開発者に連絡ください。
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

# ディレクトリの説明
## archive/
アーカイブ化したプロジェクト。
## bml/
ブックマークレットをつかうプロジェクト。
## go/
go言語もメインプロジェクト。
