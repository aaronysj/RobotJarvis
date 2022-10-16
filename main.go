package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}
var URL_FORMAT = "https://matchweb.sports.qq.com/kbs/list?from=NBA_PC&columnId=100000" +
	"&startTime=%s&endTime=%s&from=sporthp"

var logger *log.Logger

func init() {
	logFile, err := os.OpenFile("./output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("output.log open failed")
	}
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}

func main() {
	logger.Println("Jarvis Started...")
	http.HandleFunc("/robots", robotsHandler)
	logger.Fatal(http.ListenAndServe(":8080", nil))
}

func robotsHandler(w http.ResponseWriter, r *http.Request) {
	msg := GenerateMarkdown()
	body, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

/**
* 今日NBA
 */
func GenerateMarkdown() MarkDownMsgRequest {
	// 请求更新
	today := time.Now().Format("2006-01-02")
	NBA_URL := fmt.Sprintf(URL_FORMAT, today, today)

	logger.Println(NBA_URL)
	apiResult := new(TencentApiResult)
	err := getJson(NBA_URL, apiResult)
	if err != nil {
		panic(err)
	}
	// fmt.Println(apiResult.Data)
	// 推送钉钉
	games := apiResult.Data[today]
	var markdown string
	for _, game := range games {
		markdown += GetGameMarkdownInfo(&game)
	}
	markdown += links()
	// fmt.Println(markdown)

	markdownMsg := MarkdownMsg{
		Title: "NBA",
		Text:  markdown,
	}
	return MarkDownMsgRequest{
		"markdown",
		markdownMsg,
	}
}

func links() string {
	return `
👉🏻[schedule](https://nba.stats.qq.com/schedule) [standings](https://nba.stats.qq.com/standings)
👉🏻[Maigc](http://24zhiboba.com)
👉🏻[Top10](https://sports.qq.com/nbavideo/topsk/)
✌🏻[@aaronysj](https://github.com/aaronysj)
`
}

func getJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}
