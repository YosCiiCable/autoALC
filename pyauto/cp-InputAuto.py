# インプット をクリップボードにコピーしておく。
# 開始したいユニットの数字に変える。

from time import sleep
import keyboard
import pyautogui
UNIT = 1

if keyboard.read_key() == "s":
    while True:
        pyautogui.press('home')
        pyautogui.hotkey('ctrl', 'f')
        sleep(0.3)
        pyautogui.typewrite(['n', 'p', 'u', 't'])
        for x in range(UNIT):
            pyautogui.press('enter')
        pyautogui.press('esc')
        sleep(0.5)
        pyautogui.press('enter')
        sleep(3)
        pyautogui.press('f11')
        sleep(3)
        pyautogui.click(700, 215)
        sleep(2)
        pyautogui.click(700, 500)
        sleep(24)
        pyautogui.click(1250, 700)
        sleep(1)
        pyautogui.click(600, 500)
        UNIT += 1
        sleep(8)
    if keyboard.read_key() == "q":
        exit()
