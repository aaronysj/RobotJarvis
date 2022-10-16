package handler

import (
	"encoding/json"
	"net/http"

	"github.com/aaronysj/RobotJarvis/pkg/sports"
)

func RobotsHandler(w http.ResponseWriter, r *http.Request) {
	msg := sports.GenerateMarkdown()
	body, _ := json.Marshal(msg)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write(body)
}
