package sports

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/aaronysj/RobotJarvis/pkg/utils"
)

type GameInfo struct {
	MatchType     string `json:"matchType"`
	Mid           string `json:"mid"`
	WebUrl        string `json:"webUrl"`
	MatchDesc     string `json:"matchDesc"`
	MatchPeriod   string `json:"matchPeriod"`
	LeftId        string `json:"leftId"`
	LeftName      string `json:"leftName"`
	LeftBadge     string `json:"leftBadge"`
	LeftGoal      string `json:"leftGoal"`
	LeftHasUrl    string `json:"leftHasUrl"`
	RightId       string `json:"rightId"`
	RightName     string `json:"rightName"`
	RightBadge    string `json:"rightBadge"`
	RightGoal     string `json:"rightGoal"`
	RightHasUrl   string `json:"rightHasUrl"`
	StartTime     string `json:"startTime"`
	LivePeriod    string `json:"livePeriod"`
	LiveType      string `json:"liveType"`
	LiveId        string `json:"liveId"`
	Quarter       string `json:"quarter"`
	QuarterTime   string `json:"quarterTime"`
	ProgramId     string `json:"programId"`
	IsPay         string `json:"isPay"`
	GroupName     string `json:"groupName"`
	CompetitionId string `json:"competitionId"`
	TvLiveId      string `json:"tvLiveId"`
	IfHasPlayback string `json:"ifHasPlayback"`
	Url           string `json:"url"`
	CategoryId    string `json:"categoryId"`
	ScheduleId    string `json:"scheduleId"`
	RoseNewsId    string `json:"roseNewsId"`
	LatestNews    string `json:"latestNews"`
}

type TencentApiResult struct {
	Code    int                   `json:"code"`
	Version string                `json:"version"`
	Data    map[string][]GameInfo `json:"data"`
}

type MarkDownMsgRequest struct {
	MsgType  string      `json:"msgtype"`
	Markdown MarkdownMsg `json:"markdown"`
}

type MarkdownMsg struct {
	Title string `json:"title"`
	Text  string `json:"text"`
}

var NUM_0 = "0"
var NUM_1 = "1"
var NUM_2 = "2"
var NUM_3 = "3"

func GetGameMarkdownInfo(game *GameInfo) string {
	var mardown = ""

	leftGoal, _ := strconv.Atoi(game.LeftGoal)
	rightGoal, _ := strconv.Atoi(game.RightGoal)
	leftName := fmt.Sprintf("[%s](https://sports.qq.com/kbsweb/teams.htm?tid=%s&cid=100000)", game.LeftName, game.LeftId)
	rightName := fmt.Sprintf("[%s](https://sports.qq.com/kbsweb/teams.htm?tid=%s&cid=100000)", game.RightName, game.RightId)
	if Equal(NUM_2, game.MatchPeriod) {
		if leftGoal < rightGoal {
			rightName = " 🏆" + rightName
		} else if leftGoal > rightGoal {
			// 客队 win
			leftName = leftName + "🏆 "
		}
	}

	mardown += fmt.Sprintf("%s%s%s %s ", letsGoWarroir(game), free(game), game.StartTime[11:16], parseMatchPeriod(game))
	mardown += fmt.Sprintf("%s %s vs %s %s ", leftName, game.LeftGoal, game.RightGoal, rightName)
	if gameOnOrGameOver(game) {
		mardown += fmt.Sprintf("[[%s](%s) [数据](https://nba.stats.qq.com/nbascore/?mid=%s) [回放](%s&replay=1)]", video(game), game.WebUrl, strings.Split(game.Mid, ":")[1], game.WebUrl)
	}
	return mardown + "\n\n"
}

func gameOnOrGameOver(game *GameInfo) bool {
	return Equal(NUM_1, game.MatchPeriod) || Equal(NUM_2, game.MatchPeriod)
}

func video(game *GameInfo) string {
	if Equal(NUM_1, game.LivePeriod) {
		return "直播"
	} else {
		return "集锦"
	}
}
func parseMatchPeriod(game *GameInfo) string {
	var matchPeriod = "未知"
	if Equal(NUM_0, game.MatchPeriod) {
		matchPeriod = "未开始"
	} else if Equal(NUM_1, game.MatchPeriod) {
		matchPeriod = game.Quarter + " " + game.QuarterTime
	} else if Equal(NUM_2, game.MatchPeriod) {
		matchPeriod = "已结束"
	} else if Equal(NUM_3, game.MatchPeriod) {
		matchPeriod = "比赛延期"
	}
	return matchPeriod
}

func letsGoWarroir(game *GameInfo) string {
	if Equal("勇士", game.LeftName) || Equal("勇士", game.RightName) {
		return "🏀"
	} else {
		return ""
	}
}

func free(game *GameInfo) string {
	if Equal(NUM_0, game.IsPay) {
		return "😎"
	} else {
		return ""
	}
}

func Equal(s1 string, s2 string) bool {
	return s1 == s2
}

var URL_FORMAT = "https://matchweb.sports.qq.com/kbs/list?from=NBA_PC&columnId=100000" +
	"&startTime=%s&endTime=%s&from=sporthp"

/**
* 今日NBA
 */
func GenerateMarkdown(date string) MarkDownMsgRequest {
	// 请求更新
	NBA_URL := fmt.Sprintf(URL_FORMAT, date, date)

	apiResult := new(TencentApiResult)

	err := utils.GetJson(NBA_URL, apiResult)
	if err != nil {
		// todo 异常处理
		panic(err)
	}

	games := apiResult.Data[date]
	title := fmt.Sprintf("NBA(%s)", date)
	markdown := fmt.Sprintf("# %s\n\n", title)
	for _, game := range games {
		markdown += GetGameMarkdownInfo(&game)
	}
	markdown += links()

	markdownMsg := MarkdownMsg{
		Title: title,
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
