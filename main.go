package main

import (
	"log"
	"net/http"
	"os"

	"github.com/aaronysj/RobotJarvis/pkg/handler"
)

// var logger *log.Logger

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	logFile, err := os.OpenFile("./output.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic("output.log open failed")
	}
	log.SetOutput(logFile)
	// logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)

}

func main() {
	log.Println("Jarvis Started...")
	http.HandleFunc("/robots", handler.RobotsHandler)
	http.HandleFunc("/schedule", handler.ScheduleHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
