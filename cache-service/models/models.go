package models

type God struct {
	Id         int32   `json:"id"`
	Name       string  `json:"name"`
	Role       string  `json:"role"`
	AvgKills   float32 `json:"avg_kills"`
	AvgDeaths  float32 `json:"avg_deaths"`
	AvgAssists float32 `json:"avg_assists"`
	AvgGold    float32 `json:"avg_gold"`
	WinRate    float32 `json:"win_rate"`
	PickRate   float32 `json:"pick_rate"`
	BanRate    float32 `json:"ban_rate"`
	HotItems   []Item  `json:"hot_items"`
	TopItems   []Item  `json:"top_items"`
}

type Item struct {
	Id   int32  `json:"id"`
	Name string `json:"name"`
}

// message God{
//     int32 id = 1;
//     string name = 2;
//     string role = 3;
//     float avg_kills = 4;
//     float avg_deaths = 5;
//     float avg_assists = 6;
//     float avg_gold = 7;
//     float win_rate = 8;
// }
