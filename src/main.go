package main

import (
	"encoding/json"
	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/linebot"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type WeatherList struct {
	ReportDatetime string    `json:"reportDatetime"`
	TimeSeries     []OneTime `json:"timeSeries"`
}

type OneTime struct {
	TimeDefines []string `json:"timeDefines"`
	Areas       []Area   `json:"areas"`
}

type Area struct {
	AreaName    string   `json:"area>name"`
	Weathers    []string `json:"weathers"`
	Temperature []string `json:"temps"`
}

func httpGetStr(url string) []byte {
	// HTTPリクエストを発行しレスポンスを取得する
	response, err := http.Get(url)
	if err != nil {
		log.Fatal("Get Http Error:", err)
	}
	// レスポンスボディを読み込む
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal("IO Read Error:", err)
	}
	// 読み込み終わったらレスポンスボディを閉じる
	defer response.Body.Close()
	return body
}

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	bot, err := linebot.New(
		os.Getenv("LINE_BOT_CHANNEL_SECRET"),
		os.Getenv("LINE_BOT_CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	body := httpGetStr("https://www.jma.go.jp/bosai/forecast/data/forecast/340000.json")
	print(body)
	result := make([]*WeatherList, 0)
	err = json.Unmarshal(body, &result)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result[0].TimeSeries[0].Areas[0].Weathers[0])
	todayWeather := result[0].TimeSeries[0].Areas[0].Weathers[1]
	lowestTemp := result[0].TimeSeries[2].Areas[0].Temperature[0]
	highestTemp := result[0].TimeSeries[2].Areas[0].Temperature[1]
	modifiedTodayWeather := strings.Join(strings.Split(todayWeather, "　"), "")
	if strings.ContainsAny(todayWeather, "雨雪") {
		message := linebot.NewTextMessage(
			"明日の天気は\n" +
				modifiedTodayWeather + "\n" +
				"だよー\n\n" +
				"朝の最低気温　：" + lowestTemp + "℃\n" +
				"日中の最高気温：" + highestTemp + "℃")
		if _, err := bot.BroadcastMessage(message).Do(); err != nil {
			log.Fatal(err)
		}
	}

}
