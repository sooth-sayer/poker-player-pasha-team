package rank_api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lean-poker/poker-player-go/leanpoker"
)

type Rank struct {
	Rank        int64            `json:"rank"`
	Value       int64            `json:"value"`
	SecondValue int64            `json:"second_value"`
	Kickers     []int64          `json:"kickers"`
	CardsUsed   []leanpoker.Card `json:"cards_used"`
	Cards       `json:"cards"`
}

type Cards []leanpoker.Card

func GetRank(cards Cards) Rank {
	apiUrl := "http://rainman.leanpoker.org/rank"

	data, _ := json.Marshal(cards)
	str := "cards=" + string(data)
	fmt.Println(str)

	client := &http.Client{}
	r, _ := http.NewRequest("GET", apiUrl, bytes.NewBufferString(str))
	r.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, _ := client.Do(r)
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	var result Rank
	decoder.Decode(&result)

	fmt.Println(resp.Status)
	fmt.Println(result)

	return result
}
