package utils

import (
	"bufio"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func ParseBody(r *http.Request, x interface{}) {
	if body, err := ioutil.ReadAll(r.Body); err == nil {
		if err := json.Unmarshal([]byte(body), x); err != nil {
			return
		}
	}
}

func GetTokens() []string {
	tokenFile, err := os.Open("./config/tokens.txt")
	if err != nil {
		log.Fatal(err)
		return nil
	}
	defer tokenFile.Close()

	fileScanner := bufio.NewScanner(tokenFile)
	fileScanner.Split(bufio.ScanLines)
	var tokens []string

	for fileScanner.Scan() {
		tokens = append(tokens, fileScanner.Text())
	}
	return tokens
}
