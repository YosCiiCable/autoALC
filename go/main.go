
package main

import (
    "github.com/sclevine/agouti"
    "log"
    "time"
)

func main() {
    // ブラウザはChromeを指定して起動
    driver := agouti.ChromeDriver(agouti.Browser("chrome"))
    if err := driver.Start(); err != nil {
        log.Fatalf("Failed to start driver:%v", err)
    }
    defer driver.Stop()

    page, err := driver.NewPage()
    if err != nil {
        log.Fatalf("Failed to open page:%v", err)
    }
    // ログインページに遷移
    if err := page.Navigate("https://qiita.com/login"); err != nil {
        log.Fatalf("Failed to navigate:%v", err)
    }
    // ID, Passの要素を取得し、値を設定
    identity := page.FindByID("identity")
    password := page.FindByID("password")
    identity.Fill("Your Id Here.")
    password.Fill("Your Passowrd Here.")
    // formをサブミット
    if err := page.FindByClass("loginSessionsForm_submit").Submit(); err != nil {
        log.Fatalf("Failed to login:%v", err)
    }
    // 処理完了後、3秒間ブラウザを表示しておく
    time.Sleep(3 * time.Second)
}
