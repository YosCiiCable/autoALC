package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	// chromeのインスタンス作成(上はデバックON)
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	//ctx, cancel := chromedp.NewContext(context.Background(),)
	defer cancel()

	// chromedpのタスク
	var buf []byte
	err := chromedp.Run(ctx,
		startup(),
		donePrt(),
		inputRunner(),
		donePrt(),
		chromedp.FullScreenshot(&buf, 90),
	)
	if err != nil {
		log.Fatal(err)
	}

	// デバック用
	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
		log.Fatal("終了: スクリーンショット出力エラー")
	}
}

func donePrt() chromedp.Tasks	{
	fmt.Println("完了")
	return chromedp.Tasks{
		chromedp.WaitVisible(`/html/head/meta[1]`),
	}
}

func startup() chromedp.Tasks {
	// input (not using chromedp)
	var studentNum, password string
	fmt.Print("sを含めた学籍番号:")
	fmt.Scan(&studentNum)
	fmt.Print("ALCのパスワード:")
	fmt.Scan(&password)
	fmt.Println("実行中…")

	return chromedp.Tasks{
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/stlogin/index/nit-ariake/"),
// login
		chromedp.WaitVisible(`//*[@id="Password"]`),
		chromedp.SendKeys(`//*[@id="AccountId"]`, studentNum),
		chromedp.SendKeys(`//*[@id="Password"]`, password),
		chromedp.Click(`//*[@id="BtnLogin"]`),

		// page move
		chromedp.Click(`//*[@id="LbtSubCourseLink_1"]`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="DivAllSubCourseTable"]/table/tbody/tr/td[2]/table/tbody/tr[3]/td[1]/a`, chromedp.NodeVisible),

		chromedp.WaitVisible(`/html/body/div[1]/div[5]/div[4]/div/div`),
	}
}

func inputRunner() chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Click(`//*[@id="BtnLogin"]`),
	}	
}
