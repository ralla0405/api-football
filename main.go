package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type StandingTable struct {
	Response []struct {
		League struct {
			ID        uint   `json:"id"`
			Name      string `json:"name"`
			Standings [][]struct {
				Rank uint `json:"rank"`
				Team struct {
					ID   uint   `json:"id"`
					Name string `json:"name"`
					Logo string `json:"logo"`
				}
			}
		}
	}
}

func getStandings(leagueID, season string) (*StandingTable, error) {
	url := "https://v3.football.api-sports.io/standings?league=" + leagueID + "&season=" + season
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	req.Header.Add("x-apisports-key", "705b214fd2f1c3a7fe8fb97d81118fb1")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	var table StandingTable
	err = json.Unmarshal(body, &table)
	if err != nil {
		return nil, err
	}

	return &table, nil
}

func main() {
	premierLeagueID := "39" // 프리미어 리그의 리그 ID 예시
	kLeagueID := "292"      // K리그의 리그 ID 예시
	season := "2024"
	plSeason := "2023"

	fmt.Println("K리그 순위:")
	kLeagueTable, err := getStandings(kLeagueID, season)
	if err != nil {
		fmt.Println("Error fetching K-League rankings:", err)
		return
	}

	for _, s := range kLeagueTable.Response[0].League.Standings[0] {
		fmt.Printf("%d - %s\n", s.Rank, s.Team.Name)
	}

	fmt.Println("\n프리미어 리그 순위:")
	premierLeagueTable, err := getStandings(premierLeagueID, plSeason)
	if err != nil {
		fmt.Println("Error fetching Premier League rankings:", err)
		return
	}

	for _, s := range premierLeagueTable.Response[0].League.Standings[0] {
		fmt.Printf("%d - %s\n", s.Rank, s.Team.Name)
	}
}
