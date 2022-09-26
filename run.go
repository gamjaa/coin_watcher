package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
)

type ChartResponse struct {
	Result    string  `json:"result"`
	ErrorCode string  `json:"error_code"`
	IsLast    bool    `json:"is_last"`
	Charts    []Chart `json:"chart"`
}

type Chart struct {
	Timestamp    int64  `json:"timestamp"`
	Open         string `json:"open"`
	High         string `json:"high"`
	Low          string `json:"low"`
	Close        string `json:"close"`
	TargetVolume string `json:"target_volume"`
	QuoteVolume  string `json:"quote_volume"`
}

func main() {
	currency := "KRW"
	coin := "ITAMCUBE"
	interval := "6h"
	url := fmt.Sprintf("https://api.coinone.co.kr/public/v2/chart/%s/%s?interval=%s", currency, coin, interval)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	//필요시 헤더 추가 가능
	req.Header.Add("Content-type", "application/json")
	//req.Header.Add("X-COINONE-PAYLOAD", "application/json")
	//req.Header.Add("X-COINONE-SIGNATURE", "application/json")

	// Client객체에서 Request 실행
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("API ERROR")
		panic(err)
	}
	defer resp.Body.Close()

	// 결과 출력
	bytes, _ := io.ReadAll(resp.Body)
	str := string(bytes) //바이트를 문자열로
	fmt.Println(str)

	data := ChartResponse{}
	//data := make(map[string]interface{})
	parseError := json.Unmarshal(bytes, &data)
	if parseError != nil {
		fmt.Println("PARSE ERROR")
		panic(parseError)
	}

	fmt.Println("OK")
	fmt.Println(len(data.Charts))
}
