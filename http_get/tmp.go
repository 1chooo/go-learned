package main

import (
	"fmt"
	"net/http"
	"time"
)

func test_get() {
	// 要測試的網站 URL
	url := "https://www.1chooo.com/"

	// 設定每分鐘執行的次數
	requestsPerMinute := 10

	// 計算每次間隔的時間
	interval := time.Minute / time.Duration(requestsPerMinute)
	
	// 無限迴圈，每隔一段時間執行HTTP GET請求
	for {
		fmt.Println("Ready to send HTTP GET request to:", url)

		// 執行HTTP GET請求
		resp, err := http.Get(url)
		if err != nil {
			fmt.Print("Error:", err)
		} else {
			// 在此處處理伺服器回應，例如檢查狀態碼等
			fmt.Println("Status code:", resp.StatusCode)
			// 關閉回應體
			resp.Body.Close()
		}

		// 等待一段時間
		time.Sleep(interval)
	}
}
