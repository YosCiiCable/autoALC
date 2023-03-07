# autoALC
ALCを自動でしてくれるツールを作りたい人生だった…。

# 実行
~~~
docker-compose up -d
docker cp gocon:/usr/src/app/fullScreenshot.png ./
~~~

# ディレクトリの説明
## /
開発時に、dockerの中で作業する用。
## archive/
アーカイブ化したプロジェクト。
### go-robotogo/**
RobotGoをつかうプロジェクト。
### usePython/
Pythonをつかうプロジェクト。
## bml/
ブックマークレットをつかうプロジェクト。
## go/
goをつかうプロジェクト。
Selenium WebDriver を利用しているので、以下の**いずれか**の操作が必要。
### Chromeを使う
- Chromeのインストール
- Chromeのバージョンに合った Selenium WebDriver のインストール
　バージョンは、Chromeの右上︙ → 設定 → 左の一覧にある一番下の「Chrome について」
　WebDriver は [こちら](https://chromedriver.chromium.org/downloads)
### Chrome以外を使う
- ブラウザに合うように [SeleniumHQ/docker-selenium](https://github.com/SeleniumHQ/docker-selenium) を参考に、Dockerfile や compose.yml の書き換え
- ブラウザに合った Selenium WebDriver のインストール
