package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

var myClient = &http.Client{Timeout: 10 * time.Second}

func GetJson(url string, target interface{}) error {
	r, err := myClient.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	return json.NewDecoder(r.Body).Decode(target)
}

var dingTalkApiPrefix = "https://oapi.dingtalk.com/robot/send?access_token=%s"

func SendToDingTalk(token string, body []byte) {
	dingTalkPushUrl := fmt.Sprintf(dingTalkApiPrefix, token)
	myClient.Post(dingTalkPushUrl, "application/json; charset=UTF-8", bytes.NewBuffer(body))
}
