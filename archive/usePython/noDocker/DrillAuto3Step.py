# sleep()を調整すると効率UPするかも。
# 半角入力にしておくこと。
import pyautogui
import keyboard
from time import sleep
pyautogui.FAILSAFE = True
# 開始したいユニット-1 の数字に変える。
UNIT = 18

# wキー押したら始まる。
# 画面左上(0, 0)にカーソルを持っていくと強制停止。
if keyboard.read_key() == "w":
    while True:
        # ページの一番上
        pyautogui.press('home')
        # 検索開く
        pyautogui.hotkey('ctrl', 'f')
        sleep(.3)
        # 検索文字初期化
        pyautogui.press('backspace')
        # 検索文字入力
        pyautogui.typewrite(['r', 'i', 'l', 'l'])
        # 検索したやつ下っていく
        for x in range(UNIT):
            pyautogui.press('enter')
        # 検索閉じる
        pyautogui.press('esc')
        sleep(.5)
        # 選択
        pyautogui.press('enter')
        sleep(3)
        # 全画面
        pyautogui.press('f11')
        sleep(3)
        # STARTクリック
        pyautogui.click(600, 210)
        sleep(1)
        # 続きを開いた時用
        # NO
        pyautogui.click(750, 500)
        sleep(0.5)
        # --STEP1--
        # OK
        pyautogui.click(700, 500)
        sleep(.5)
        # 左右選択5セット
        for x in range(4):
            # 右側クリックx20
            for x in range(20):
                pyautogui.click(700, 560)
                sleep(.8)
            # 左側クリックx20
            for x in range(20):
                pyautogui.click(400, 560)
                sleep(.8)
        # NEXT STEP
        pyautogui.click(1250, 720)
        sleep(1)
        # --STEP2--
        # OK
        pyautogui.click(700, 500)
        sleep(.5)
        # 入力バー
        pyautogui.click(300, 230)
        sleep(.5)
        for i in range(40):
            pyautogui.typewrite('a\n')
            pyautogui.press('enter')
        # NEXT STEP
        pyautogui.click(1250, 720)
        sleep(0.5)
        # --STEP3--
        # OK
        pyautogui.click(700, 500)
        sleep(.5)
        # ANSER
        pyautogui.click(1270, 350)
        sleep(.5)
        # NEXT STEP
        pyautogui.click(1250, 720)
        sleep(0.5)
        # --ウィンドウ閉じる--
        # ウィンドウ閉じるショートカットキー(clickで座標指定しても可)
        pyautogui.hotkey('alt', 'f4')
        sleep(.5)
        # 警告を許可
        pyautogui.click(690, 425)
        sleep(1)
        # UNITを1増やす
        UNIT += 1
        # ページ読み込み待ち
        sleep(5)
