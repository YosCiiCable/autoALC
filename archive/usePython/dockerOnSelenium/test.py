from selenium import webdriver
import os

# Chrome のオプションを設定する
options = webdriver.ChromeOptions()
options.add_argument('--lang=ja-JP')
#options.add_argument('--headless')

# Selenium Server に接続する
driver = webdriver.Remote(
    command_executor='http://localhost:4444/wd/hub',
    desired_capabilities=options.to_capabilities(),
    options=options,
)

# Selenium 経由でブラウザを操作する
driver.get('https://qiita.com')
FILENAME = os.path.join(os.path.dirname(os.path.abspath(__file__)), "img1.png") #ファイル名
w = driver.execute_script("return document.body.scrollWidth;")
h = driver.execute_script("return document.body.scrollHeight;")
driver.set_window_size(w,h)
driver.save_screenshot(FILENAME)
print(driver.current_url)

# ブラウザを終了する
driver.quit()
