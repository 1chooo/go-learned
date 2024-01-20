# Hello Server

這是一個簡單的 Go 程式，是一個基於 HTTP 的 Web 伺服器，使用了標準庫的 `net/http`。這個程式的主要功能是回應不同的 HTTP 請求，根據請求的路徑執行相對應的處理函數。

在這個程式中，路徑 "/name" 會被傳遞到 `greet` 函數中處理。這個函數首先從 URL 的路徑中提取名字，如果路徑中沒有名字，則預設使用 "Gopher"。

因此，當你透過瀏覽器或其他 HTTP 客戶端訪問 http://localhost:8080/name 時，這個路徑會被傳遞給 `greet` 函數，它會將預設的問候詞 "Hello" 與提取到的名字組合成回應，最終顯示在瀏覽器上，形如 "Hello, name!"。