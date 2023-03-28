package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/chromedp/cdproto/input"
	"github.com/chromedp/cdproto/target"
	"github.com/chromedp/chromedp"

	"github.com/manifoldco/promptui"
)

func main() {
	// level0: chromeのインスタンス作成
	ctx, _ := chromedp.NewContext(context.Background()) /*
		// level0-debug1: ログあり でインスタンス作成
		ctx, _ := chromedp.NewContext(context.Background(), chromedp.WithDebugf(log.Printf)) /*
		// level0-debug2: no headless でインスタンス作成
		opts := append(chromedp.DefaultExecAllocatorOptions[:],
			chromedp.Flag("headless", false),
			chromedp.Flag("enable-automation", false),
		)
		allocCtx, _ := chromedp.NewExecAllocator(context.Background(), opts...)
		ctx, _ := chromedp.NewContext(allocCtx, chromedp.WithLogf(log.Printf))
		//*/
	//defer cancel()

	// level1: login
	login(ctx)

	// level2: 遷移チェック
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtMaRePassword"]/font`, chromedp.AtLeast(0)),
	); err == nil {
		log.Fatal("err2: Transition check fails")
	}
	fmt.Println("ログイン完了")

	// level3: Unitsに移動
	fmt.Println("ユニット一覧に移動開始………")
	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="LbtSubCourseLink_1"]`, chromedp.NodeVisible),
		chromedp.Click(
			`//*[@id="DivAllSubCourseTable"]/table/tbody/tr/td[2]/table/tbody/tr[3]/td[1]/a`,
			chromedp.NodeVisible,
		),
		chromedp.WaitVisible(`//*[@id="nan-contents"]/div[1]/div/label/font`),
	); err != nil {
		log.Fatal("err3: Failed to move units")
	}
	fmt.Println("ユニット一覧移動完了")

	// level4: Input開始
	if yesNo("インプット をやりますか？") {
		fmt.Println("インプット開始………")
		for i, k := 2, 0; i < 220; i += 2 {
			k++
			inputSelector(ctx, i, (i-(k/5))/2)
			debugPic(ctx)
			if k%5 == 0 {
				i++
			}
		}
	}

	// level5: Output開始
	if yesNo("ドリル をやりますか？") {
		fmt.Println("ドリル開始………")
	}
}

func login(ctx context.Context) {
	var studentNum, passwd string
	fmt.Print("sを含めた学籍番号: ")
	fmt.Scan(&studentNum)
	passwd = passwdInputer("ALCのパスワード")
	fmt.Println("ログイン開始………")

	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://nanext.alcnanext.jp/anetn/Student/stlogin/index/nit-ariake/"),

		chromedp.Click(`//*[@id="AccountId"]`, chromedp.NodeVisible),
		input.InsertText(studentNum),
		chromedp.Click(`//*[@id="Password"]`, chromedp.NodeVisible),
		input.InsertText(passwd),
		chromedp.Click(`//*[@id="BtnLogin"]`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err1: Failed login")
	}
}

func inputSelector(ctx context.Context, i int, unitNum int) {
	iXPath := "//*[@id=\"nan-contents\"]/div[7]/div/table/tbody/tr[" + strconv.Itoa(i)
	// if naming, iXPath + "]/td[4]"	is iStatusXPath
	// if naming, iXPath + "]/td[3]/a"	is iClickXPath
	iStatusText := "init"
	if err := chromedp.Run(ctx,
		chromedp.Text(iXPath+"]/td[4]", &iStatusText),
	); err != nil {
		log.Fatal("err@inputSelector-1: Failed to get the input-unit status")
	}

	if iStatusText == "未学習 / Not studied" || iStatusText == "参照のみ / Only opened" || iStatusText == "学習中 / In progress" {
		if err := chromedp.Run(ctx,
			chromedp.Click(iXPath+"]/td[3]/a", chromedp.NodeVisible),
		); err != nil {
			log.Fatal("err@inputSelector-2: Failed to click the input-unit")
		}
		inputer(ctx, unitNum)
	} else if iStatusText == "修了 / Completed" {
		fmt.Printf("修了済み Unit%d\n", unitNum)
	} else {
		fmt.Println("warning  : 続行可能な程度の例外発生、続行します。")
		fmt.Println("---------: よければ以下の1行を開発者にお知らせください。")
		fmt.Println("warn@inputSelector-3:Exception at last of inputer:" + iStatusText)
	}
}

func inputer(ctx context.Context, unitNum int) {
	var homeID, targetID target.ID = "", ""
	fmt.Printf("Unit%dを処理中…\n", unitNum)
	for i := 0; i < 1; {
		targets, err := chromedp.Targets(ctx)
		if err != nil {
			log.Fatal("err@inputer-1: Failed to make a new target")
		}
		for _, t := range targets {
			if t.Type == "page" && t.URL != "about:blank" {
				if homeID == "" {
					homeID = t.TargetID
				}
				if t.URL != "https://nanext.alcnanext.jp/anetn/Student/StUnitList#" {
					targetID = t.TargetID
					i++
				}
			}
		}
		time.Sleep(100 * time.Millisecond)
	}

	if homeID == "" {
		log.Fatal("err@inputer-2: The homeID is nil")
		return
	}
	if targetID == "" {
		log.Fatal("err@inputer-3: The targetID is nil")
		return
	}
	ctx, _ = chromedp.NewContext(ctx, chromedp.WithTargetID(targetID))

	if err := chromedp.Run(ctx,
		chromedp.Click(`//*[@id="nan-contents-cover-buttons"]/div/div[1]/button/span[1]`, chromedp.NodeVisible),
		chromedp.Click(`.ui-dialog-buttonset > button:nth-child(1)`, chromedp.NodeVisible),
		chromedp.Click(`//*[@id="nan-toolbox-footer"]/button`, chromedp.NodeVisible),
		chromedp.Click(`.ui-dialog-buttonset > button:nth-child(1)`, chromedp.NodeVisible),
	); err != nil {
		log.Fatal("err@inputer-4: Failed to click element in input")
	}

	ctx, _ = chromedp.NewContext(ctx, chromedp.WithTargetID(homeID))

	fmt.Printf("完了: Unit%d\n", unitNum)
}

func yesNo(labelMessage string) bool {
	prompt := promptui.Select{
		Label: labelMessage,
		Items: []string{"Yes", "No"},
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatal("err@yesNo: Failed to run prompt")
	}

	return result == "Yes"
}

func passwdInputer(labelMessage string) string {
	validate := func(input string) error {
		if len(input) < 6 {
			return errors.New("パスワードは6~20文字の範囲の筈です…")
		} else if len(input) > 20 {
			return errors.New("パスワードは6~20文字の範囲の筈です…")
		}
		return nil
	}

	prompt := promptui.Prompt{
		Label:    labelMessage,
		Validate: validate,
		Mask:     '*',
	}

	passwd, err := prompt.Run()
	if err != nil {
		log.Fatal("err@passwdInputer: Failed to run prompt")
	}

	return passwd
}

func debugURL(ctx context.Context) {
	var url string
	if err := chromedp.Run(ctx,
		chromedp.Location(&url),
	); err != nil {
		log.Fatal("err@debugURL: Failed to location url")
	}
	fmt.Printf("debugURL: %s\n", url)

}

func debugPic(ctx context.Context) {
	var buf []byte
	// debug level1: スクショ撮影
	if err := chromedp.Run(ctx,
		chromedp.FullScreenshot(&buf, 90),
	); err != nil {
		log.Fatal("err@debugPic-1: Failed to capture a screenshot")
	}
	// debug level2: スクショ出力
	if err := os.WriteFile(
		"fullScreenshot.png", buf, 0o644,
	); err != nil {
		log.Fatal("err@debugPic-2: Failed to output the screenshot")
	}
}
