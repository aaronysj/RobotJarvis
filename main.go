package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aaronysj/RobotJarvis/pkg/handler"
)

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
	http.HandleFunc("/robots", handler.RobotsHandler)
	logger.Fatal(http.ListenAndServe(":8080", nil))
}
