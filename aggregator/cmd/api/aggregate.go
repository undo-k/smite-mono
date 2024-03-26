package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sort"

	"github.com/undo-k/smite-mono/protos/protos"
)

type Affinity string
type Role string
type Mode string

const (
	Physical Affinity = "Physical"
	Magical  Affinity = "Magical"
	Hybrid   Affinity = "Hybrid"
)

const (
	Assassin Role = "Assassin"
	Guardian Role = "Guardian"
	Hunter   Role = "Hunter"
	Mage     Role = "Mage"
	Warrior  Role = "Warrior"
)

const (
	Arena    Mode = "Arena"
	Assault  Mode = "Assault"
	Conquest Mode = "Conquest"
	MOTD     Mode = "MOTD"
	Siege    Mode = "Siege"
	Slash    Mode = "Slash"
)

type God struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Affinity Affinity `json:"affinity"`
	Role     Role     `json:"role"`
}

type Item struct {
	Id       int      `json:"id"`
	Name     string   `json:"name"`
	Affinity Affinity `json:"affinity"`
}

type Player struct {
	Id      int    `json:"id"`
	Team    string `json:"team"`
	Won     bool   `json:"won"`
	God     God    `json:"god"`
	Gold    int    `json:"gold"`
	Kills   int    `json:"kills"`
	Deaths  int    `json:"deaths"`
	Assists int    `json:"assists"`
	Items   []Item `json:"items"`
}

type Match struct {
	Id          int      `json:"id"`
	Mode        Mode     `json:"mode"`
	Players     []Player `json:"players"`
	WinningTeam string   `json:"winning_team"`
}

type Totals struct {
	Id           int
	Name         string
	Role         Role
	TotalKills   int
	TotalDeaths  int
	TotalAssists int
	TotalGold    int
	TotalMatches int
	TotalWins    int
	TotalPicks   int
	TotalBans    int
	ItemTotals   map[Item]ItemTotals
}

type ItemTotals struct {
	Wins         int
	TotalMatches int
}

type ItemWinRates struct {
	Item    Item
	WinRate float32
}

func batchData(numOfRequests int) {

	matchIds := fetchMatchIds(numOfRequests)
	matches := fetchMatchDetails(matchIds)
	aggregate(matches)

}

func fetchMatchDetails(matchIds []string) []Match {
	reqUrl := "http://mock-api-service:8081/GetMatchDetailsBatch/5v346eg534/jables/1251235341/yyyyMMddHHmmss/"

	for i, id := range matchIds {
		reqUrl += id
		if i != len(matchIds)-1 {
			reqUrl += ","
		}
	}

	resp, err := http.Get(reqUrl)

	if err != nil {
		panic(err)
	}

	var responseObject struct {
		Matches []Match `json:"data"`
	}

	err = json.NewDecoder(resp.Body).Decode(&responseObject)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Successfully batched data for %v matches\n", len(responseObject.Matches))

	return responseObject.Matches
}

func fetchMatchIds(numOfRequests int) []string {
	var allMatchIds []string

	for i := 0; i < numOfRequests; i++ {
		resp, err := http.Get("http://mock-api-service:8081/GetMatchIdsByQueue/5v346eg534/jables/1251235341/yyyyMMddHHmmss/conquest/20240217/2")

		if err != nil {
			panic(err)
		}

		var matchIds struct {
			MatchIds []string `json:"match_ids"`
		}

		json.NewDecoder(resp.Body).Decode(&matchIds)

		allMatchIds = append(allMatchIds, matchIds.MatchIds...)
	}

	fmt.Printf("Read %v match IDs\n", len(allMatchIds))

	return allMatchIds
}

func aggregate(matches []Match) {
	totals := gatherTotals(matches)

	stats := gatherStats(totals)

	fmt.Printf("Gathered stats for %v gods\n", len(stats))

	for i := range stats {
		err := PutGodViaGRPC(&protos.God{
			Id:         stats[i].Id,
			Name:       stats[i].Name,
			Role:       stats[i].Role,
			AvgKills:   stats[i].AvgKills,
			AvgDeaths:  stats[i].AvgDeaths,
			AvgAssists: stats[i].AvgAssists,
			AvgGold:    stats[i].AvgGold,
			WinRate:    stats[i].WinRate,
			HotItems:   stats[i].HotItems,
		})

		if err != nil {
			panic(err)
		}
	}

}

func gatherStats(totals map[string]Totals) []protos.God {
	var godStats []protos.God

	for _, total := range totals {
		hotItems := calculateHotItems(total.ItemTotals)

		godStats = append(godStats, protos.God{
			Id:         int32(total.Id),
			Name:       total.Name,
			Role:       string(total.Role),
			AvgKills:   float32(total.TotalKills) / float32(total.TotalMatches),
			AvgDeaths:  float32(total.TotalDeaths) / float32(total.TotalMatches),
			AvgAssists: float32(total.TotalAssists) / float32(total.TotalMatches),
			AvgGold:    float32(total.TotalGold) / float32(total.TotalMatches),
			WinRate:    float32(total.TotalWins) / float32(total.TotalMatches),
			HotItems:   hotItems,
		})

		if total.Name == "Medusa" {
			fmt.Printf("total wins %v / total matches %v = win rate %v \n", total.TotalWins, total.TotalMatches, float32(total.TotalWins)/float32(total.TotalMatches))
		}

		// fmt.Printf("gatherStates: %v has win rate: %v \n", total.Name, float32(total.TotalWins)/float32(total.TotalMatches))
	}

	return godStats
}

func calculateHotItems(itemMap map[Item]ItemTotals) []*protos.Item {

	itemRates := make([]ItemWinRates, 0, len(itemMap))

	for item, totals := range itemMap {
		itemRates = append(itemRates, ItemWinRates{
			Item:    item,
			WinRate: float32(totals.Wins) / float32(totals.TotalMatches),
		})
	}

	sort.Slice(itemRates, func(i, j int) bool {
		return itemRates[i].WinRate > itemRates[j].WinRate
	})

	var hotItems []*protos.Item

	itemCap := 6
	if len(itemMap) < 6 {
		itemCap = len(itemMap)
	}

	for _, itemWR := range itemRates[:itemCap] {
		hotItems = append(hotItems, &protos.Item{
			Id:   0,
			Name: itemWR.Item.Name,
		})
	}

	return hotItems
}

func gatherTotals(matches []Match) map[string]Totals {

	totalsMap := make(map[string]Totals)
	for _, match := range matches {
		for _, player := range match.Players {

			godName := player.God.Name

			entry, ok := totalsMap[godName]
			if ok {
				entry.TotalKills += player.Kills
				entry.TotalDeaths += player.Deaths
				entry.TotalAssists += player.Assists
				entry.TotalGold += player.Gold
				entry.TotalMatches += 1
				entry.TotalWins += hasWon(player.Won)
				entry.ItemTotals = makeItemTotals(player, entry.ItemTotals)

				totalsMap[godName] = entry

			} else {
				itemsMap := make(map[Item]ItemTotals)

				totalsMap[godName] = Totals{
					Id:           player.God.Id,
					Name:         player.God.Name,
					Role:         player.God.Role,
					TotalKills:   player.Kills,
					TotalDeaths:  player.Deaths,
					TotalAssists: player.Assists,
					TotalGold:    player.Gold,
					TotalMatches: 1,
					TotalWins:    hasWon(player.Won),
					ItemTotals:   makeItemTotals(player, itemsMap),
				}
			}

		}
	}

	return totalsMap
}

func makeItemTotals(player Player, itemsMap map[Item]ItemTotals) map[Item]ItemTotals {
	for _, item := range player.Items {
		itemTotals := itemsMap[item]
		itemTotals.TotalMatches += 1
		itemTotals.Wins += hasWon(player.Won)

		itemsMap[item] = itemTotals
	}

	return itemsMap
}

func hasWon(w bool) int {
	if w {
		return 1
	}
	return 0
}
