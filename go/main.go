package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"github.com/chromedp/chromedp"
	"github.com/manifoldco/promptui"
)

func main() {
	// level0: chromeのインスタンス作成(上はデバックON)
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	//ctx, cancel := chromedp.NewContext(context.Background(),)
	defer cancel()

	// level1: login
	if err := chromedp.Run(ctx,
		login(),
	); err != nil {
		log.Fatal("err1: login")
	}
	// level1-2: エラーチェック
//	loginErrorMessage := "hoge"
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtMaRePassword"]/font`, chromedp.AtLeast(0)),
//		chromedp.Text(`"//*[@id="form1"]/div[2]/div[2]/div[5]/div/span[2]/span"`, &loginErrorMessage, chromedp.AtLeast(0)),
	); err == nil {
		log.Fatal("err1: Login Information Error")
//	} else if loginErrorMessage != "hoge" {
//		log.Fatal("err1: Login Information Error")
	}
	fmt.Println("ログイン完了")

	// level2: Unitsに移動
	fmt.Println("ユニット一覧に移動開始…")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_1"]`, chromedp.NodeVisible),
		chromedp.Click(
			`//*[@id="DivAllSubCourseTable"]/table/tbody/tr/td[2]/table/tbody/tr[3]/td[1]/a`,
			chromedp.NodeVisible,
		),
	); err != nil {
		log.Fatal("err2: move to Units")
	}
	fmt.Println("ユニット一覧移動完了")

	// level3: Input開始
//	fmt.Println("種別:インプット をやりますか？")
	if yesNo("種別:インプット をやりますか？") {
		//if err := chromedp.Run(ctx, inputer(), ); err != nil { log.Fatal("err3: run Input") }
		fmt.Println("Input開始")
	}

	// level4: Output開始
	fmt.Println("種別:ドリル をやりますか？")
}

func login() chromedp.Tasks {
	var studentNum, password string
	fmt.Print("sを含めた学籍番号:")
	fmt.Scan(&studentNum)
	fmt.Print("ALCのパスワード:")
	fmt.Scan(&password)
	fmt.Println("ログイン開始…")

	return chromedp.Tasks{
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/stlogin/index/nit-ariake/"),

		chromedp.WaitVisible(`//*[@id="Password"]`),
		chromedp.SendKeys(`//*[@id="AccountId"]`, studentNum),
		chromedp.SendKeys(`//*[@id="Password"]`, password),
		chromedp.Click(`//*[@id="BtnLogin"]`, chromedp.NodeVisible),
	}
}

//func inputer() chromedp.Tasks {
//	fmt.Println("入力開始")
//	return chromedp.Tasks{
//		chromedp.Click(`//`, chromedp.NodeVisible),
//	}
//}

func yesNo(labelMessage string) bool {
	prompt := promptui.Select{
		Label: labelMessage,
		Items: []string{"Yes", "No"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}
	return result == "Yes"
}

func debug() {
	var buf []byte
	ctx, cancel := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf))
	defer cancel()
	// debug level1: スクショ撮影
	if err := chromedp.Run(ctx, chromedp.FullScreenshot(&buf, 90), ); err != nil { log.Fatal("errD1: capture a screenshot") }
	// debug level2: スクショ出力
	if err := os.WriteFile("fullScreenshot.png", buf, 0o644); err != nil { log.Fatal("errD2: output the screenshot") }
}
