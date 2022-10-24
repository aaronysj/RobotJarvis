package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/aaronysj/RobotJarvis/pkg/sports"
	"github.com/aaronysj/RobotJarvis/pkg/utils"
)

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	today := time.Now().Format("2006-01-02")
	msg := sports.GenerateMarkdown(today)
	body, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}

func ScheduleHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 token
	tokens := utils.GetTokens()
	if tokens == nil {
		return
	}
	// 生成 markdown 数据
	today := time.Now().Format("2006-01-02")
	todayBody := getMarkdownBody(today)

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	tomorrowBody := getMarkdownBody(tomorrow)
	// 发送到 token
	for _, token := range tokens {
		utils.SendToDingTalk(token, todayBody)
		utils.SendToDingTalk(token, tomorrowBody)
	}
	w.WriteHeader(http.StatusOK)
}

func getMarkdownBody(date string) []byte {
	msg := sports.GenerateMarkdown(date)
	body, _ := json.Marshal(msg)
	return body
}
