package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	// 要測試的網站 URL
	url := "https://www.1chooo.com/"

	// 設定每分鐘執行的次數
	requestsPerMinute := 10

	// 計算每次間隔的時間
	interval := time.Minute / time.Duration(requestsPerMinute)

	// 用於設定 HTTP 客戶端的設定
	client := &http.Client{
		// 設定超時時間
		Timeout: time.Second * 30,
	}
	fmt.Println("Client:", client)

	// 無限迴圈，每隔一段時間執行 HTTP GET 請求
	for {
		fmt.Println("Sending HTTP GET request to:", url)

		// 建立 HTTP 請求
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			fmt.Println("Error creating request:", err)
			continue
		}

		// 設定 User-Agent 等相關選項
		req.Header.Set("User-Agent", "YourUserAgent")

		// 執行 HTTP GET 請求
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error:", err)
		} else {
			// 在此處處理伺服器回應，例如檢查狀態碼等
			fmt.Println("Status Code:", resp.StatusCode)

			// 關閉回應體
			resp.Body.Close()
		}

		// 等待一段時間
		time.Sleep(interval)
	}
}
